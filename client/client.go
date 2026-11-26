// Package localgateway 是 LocalGateway 管理面的薄 Go 运行时。
//
// 它负责四件事：
//   - 从 HKCU\Software\LocalGateway\port 找到网关端口（可用
//     LOCALGATEWAY_CONFIG_REGISTRY 重定向子键，且可注入读取函数）；
//   - 用 POST /gw/session 握手并缓存 bearer token；
//   - 封装注册、PATCH、注销，并在收到 401 时静默重新握手并重试一次；
//   - 在网关重启/换端口时刷新基地址，并对连接失败重试一次。
//
// 请求与响应模型来自 generated/gatewayapi 包（由 OpenAPI 文档生成）。
package localgateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/HuanyuMark/localgateway-go/generated/gatewayapi"
)

// 校验错误。调用方可用 errors.Is 判定，错误信息中会附带具体原因。
var (
	// ErrInvalidName 表示服务名为空。
	ErrInvalidName = errors.New("localgateway: service name is required")
	// ErrInvalidPort 表示端口不在 1..65535 范围内。
	ErrInvalidPort = errors.New("localgateway: port must be between 1 and 65535")
	// ErrInvalidVersion 表示版本号为空。
	ErrInvalidVersion = errors.New("localgateway: version is required")
	// ErrInvalidProbe 表示 probe.path 缺失或不以 "/" 开头。
	ErrInvalidProbe = errors.New("localgateway: probe.path must start with '/'")
	// ErrEmptyPatch 表示 PATCH 没有携带任何可更新字段。
	ErrEmptyPatch = errors.New("localgateway: patch must change at least one of APIInfo, ProjectRoot, Probe, CORS")
	// ErrInvalidCORS 表示跨域配置缺少来源或预检缓存超出上限。
	ErrInvalidCORS = errors.New("localgateway: CORS requires AllowOrigins unless AllowAll is true, and MaxAgeSeconds must be at most 604800")
	// ErrInvalidBaseURL 表示管理面基地址不是合法的 http(s) URL。
	ErrInvalidBaseURL = errors.New("localgateway: base URL must be an absolute http or https URL")
	// ErrNoPortReader 表示没有配置 PortReader，无法重新发现网关端口。
	ErrNoPortReader = errors.New("localgateway: no PortReader configured to refresh the gateway base URL")
)

// PortReader 返回网关管理面端口。默认实现 ReadPort 只在 Windows 上可用。
type PortReader func() (uint16, error)

// PtrInt32 返回指向 v 的指针，便于构造 Probe.TimeoutMs。
func PtrInt32(v int32) *int32 { return &v }

// PtrString 返回指向 v 的指针，便于构造 PatchOptions.ProjectRoot。
func PtrString(v string) *string { return &v }

// PtrInt64 返回指向 v 的指针，便于构造 Cors.MaxAgeSeconds。
func PtrInt64(v int64) *int64 { return &v }

// maxCorsMaxAgeSeconds 与网关的预检缓存上限一致（7 天）。
const maxCorsMaxAgeSeconds = 604_800

// Cors 是服务级跨域配置。
//
// AllowAll 为 true 时一键放行所有来源、方法与请求头；否则至少需要一个
// AllowOrigins。配置后由网关直接应答预检并统一改写 Access-Control-* 响应头，
// provider 不需要自行实现 CORS。
type Cors struct {
	AllowAll bool
	// AllowOrigins 例如 []string{"http://localhost:5173"}；AllowAll 时忽略。
	AllowOrigins []string
	// AllowMethods 为空时回显预检请求声明的方法。
	AllowMethods []string
	// AllowHeaders 为空时回显预检请求声明的请求头。
	AllowHeaders []string
	// ExposeHeaders 是暴露给浏览器的响应头。
	ExposeHeaders []string
	// AllowCredentials 允许携带 Cookie / Authorization。
	AllowCredentials bool
	// MaxAgeSeconds 是预检缓存秒数；nil 时不发送 Access-Control-Max-Age。
	MaxAgeSeconds *int64
}

func (c Cors) validate() error {
	if !c.AllowAll && len(c.AllowOrigins) == 0 {
		return ErrInvalidCORS
	}
	if c.MaxAgeSeconds != nil && (*c.MaxAgeSeconds < 0 || *c.MaxAgeSeconds > maxCorsMaxAgeSeconds) {
		return ErrInvalidCORS
	}
	return nil
}

func (c Cors) toModel() gatewayapi.CorsConfig {
	model := gatewayapi.CorsConfig{
		AllowAll:         &c.AllowAll,
		AllowOrigins:     append([]string(nil), c.AllowOrigins...),
		AllowMethods:     append([]string(nil), c.AllowMethods...),
		AllowHeaders:     append([]string(nil), c.AllowHeaders...),
		ExposeHeaders:    append([]string(nil), c.ExposeHeaders...),
		AllowCredentials: &c.AllowCredentials,
	}
	if c.MaxAgeSeconds != nil {
		model.MaxAgeSeconds.Set(c.MaxAgeSeconds)
	}
	return model
}

// Options 控制客户端的网关地址与底层依赖。
type Options struct {
	// BaseURL 是网关管理面地址，例如 http://127.0.0.1:47821。
	// 为空时使用 PortReader（缺省 ReadPort）读取端口并拼成
	// http://127.0.0.1:<port>。
	BaseURL string
	// PortReader 可替换端口读取实现。
	//
	// BaseURL 为空时用它解析初始地址；配置后注册、PATCH、DELETE 遇到连接失败
	// 时会重新调用它刷新基地址。只有 BaseURL 而没有 PortReader 时不做刷新。
	PortReader PortReader
	// HTTPClient 可替换底层 HTTP 客户端，测试时注入。
	HTTPClient *http.Client
}

// Probe 是探活配置；Path 必填且以 "/" 开头。
type Probe struct {
	Path string
	// TimeoutMs 为空时使用网关默认值（2000ms）。
	TimeoutMs *int32
}

// RegisterOptions 是 PUT /gw/services/{name} 的请求体。
type RegisterOptions struct {
	Port    int
	Version string
	Probe   Probe
	// APIInfo 是原始 OpenAPI 3.0/3.1 文档，可为 nil，之后用 PATCH 补充。
	APIInfo any
	// ProjectRoot 为空时不发送。
	ProjectRoot string
	// NotifyURL 为空时不发送。
	NotifyURL string
	// CORS 为 nil 时不发送；设置后由网关接管该服务的跨域响应头。
	CORS *Cors
}

// PatchOptions 是 PATCH /gw/services/{name} 的请求体。
// 零值字段表示不修改；APIInfo 首版不支持用 nil 清空。
type PatchOptions struct {
	APIInfo any
	// ProjectRoot 为 nil 时不修改。
	ProjectRoot *string
	// Probe 为 nil 时不修改。
	Probe *Probe
	// CORS 为 nil 时不修改；非 nil 时替换整个跨域配置。
	CORS *Cors
}

// Client 是并发安全的管理面客户端。token 在首次写操作时懒握手获取。
//
// 基地址可以在运行时更新（网关重启或换端口）：SetBaseURL 直接指定，
// RefreshBaseURL 重新读取 PortReader。注册、PATCH、DELETE 遇到连接失败时
// 会重新读取 PortReader 并重试一次。
type Client struct {
	urlMu      sync.Mutex
	baseURL    string
	api        *gatewayapi.APIClient
	httpClient *http.Client
	portReader PortReader

	mu    sync.Mutex
	token string
}

// New 创建客户端。BaseURL 为空时会调用端口读取函数。
func New(opts Options) (*Client, error) {
	reader := opts.PortReader

	baseURL := strings.TrimSpace(opts.BaseURL)
	if baseURL == "" {
		if reader == nil {
			reader = ReadPort
		}
		port, err := reader()
		if err != nil {
			return nil, fmt.Errorf("localgateway: resolve gateway port: %w", err)
		}
		baseURL = fmt.Sprintf("http://127.0.0.1:%d", port)
	} else {
		normalized, err := normalizeBaseURL(baseURL)
		if err != nil {
			return nil, err
		}
		baseURL = normalized
	}

	httpClient := opts.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	return &Client{
		baseURL:    baseURL,
		api:        newAPIClient(baseURL, httpClient),
		httpClient: httpClient,
		portReader: reader,
	}, nil
}

// BaseURL 返回客户端实际使用的网关地址。
func (c *Client) BaseURL() string {
	c.urlMu.Lock()
	defer c.urlMu.Unlock()
	return c.baseURL
}

// SetBaseURL 更新管理面基地址并丢弃缓存的 token。
//
// 网关换了端口或重启后调用。地址必须是绝对 http(s) URL；末尾斜杠会被去掉。
func (c *Client) SetBaseURL(baseURL string) error {
	normalized, err := normalizeBaseURL(baseURL)
	if err != nil {
		return err
	}
	c.setBaseURL(normalized)
	return nil
}

// RefreshBaseURL 重新调用 PortReader 发现网关端口并更新基地址。
//
// 没有配置 PortReader 时返回 ErrNoPortReader。成功后缓存的 token 会被丢弃。
func (c *Client) RefreshBaseURL() error {
	changed, err := c.refreshBaseURL()
	if err != nil {
		return err
	}
	if !changed {
		return ErrNoPortReader
	}
	return nil
}

// setBaseURL 替换底层生成客户端；token 属于旧网关，必须丢弃。
func (c *Client) setBaseURL(baseURL string) {
	c.urlMu.Lock()
	c.baseURL = baseURL
	c.api = newAPIClient(baseURL, c.httpClient)
	c.urlMu.Unlock()

	c.mu.Lock()
	c.token = ""
	c.mu.Unlock()
}

// refreshBaseURL 按注册表端口刷新基地址；返回是否配置了 PortReader。
func (c *Client) refreshBaseURL() (bool, error) {
	c.urlMu.Lock()
	reader := c.portReader
	c.urlMu.Unlock()
	if reader == nil {
		return false, nil
	}
	port, err := reader()
	if err != nil {
		return false, fmt.Errorf("localgateway: resolve gateway port: %w", err)
	}
	c.setBaseURL(fmt.Sprintf("http://127.0.0.1:%d", port))
	return true, nil
}

func (c *Client) currentAPI() *gatewayapi.APIClient {
	c.urlMu.Lock()
	defer c.urlMu.Unlock()
	return c.api
}

func newAPIClient(baseURL string, httpClient *http.Client) *gatewayapi.APIClient {
	cfg := gatewayapi.NewConfiguration()
	cfg.Servers = gatewayapi.ServerConfigurations{{URL: baseURL}}
	cfg.HTTPClient = httpClient
	return gatewayapi.NewAPIClient(cfg)
}

func normalizeBaseURL(baseURL string) (string, error) {
	trimmed := strings.TrimRight(strings.TrimSpace(baseURL), "/")
	parsed, err := url.Parse(trimmed)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
		return "", fmt.Errorf("%w: got %q", ErrInvalidBaseURL, baseURL)
	}
	return trimmed, nil
}

// Register 注册或接管同名服务，返回网关的注册结果。
func (c *Client) Register(ctx context.Context, name string, opts RegisterOptions) (*gatewayapi.RegistrationResponse, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	if opts.Port < 1 || opts.Port > 65535 {
		return nil, ErrInvalidPort
	}
	if strings.TrimSpace(opts.Version) == "" {
		return nil, ErrInvalidVersion
	}
	probe, err := newProbe(opts.Probe)
	if err != nil {
		return nil, err
	}

	protocol := gatewayapi.HTTP
	request := gatewayapi.RegisterRequest{
		Port:     int32(opts.Port), // #nosec G115 -- range checked above.
		Protocol: &protocol,
		Version:  opts.Version,
		Probe:    *gatewayapi.NewNullableProbeConfig(probe),
	}
	if opts.APIInfo != nil {
		request.ApiInfo = opts.APIInfo
	}
	if opts.ProjectRoot != "" {
		request.ProjectRoot = *gatewayapi.NewNullableString(&opts.ProjectRoot)
	}
	if opts.NotifyURL != "" {
		request.NotifyUrl = *gatewayapi.NewNullableString(&opts.NotifyURL)
	}
	if opts.CORS != nil {
		if err := opts.CORS.validate(); err != nil {
			return nil, err
		}
		model := opts.CORS.toModel()
		request.Cors = *gatewayapi.NewNullableCorsConfig(&model)
	}

	response, _, err := callWithToken(c, ctx, func(api *gatewayapi.APIClient, callCtx context.Context) (*gatewayapi.RegistrationResponse, *http.Response, error) {
		return api.ServicesAPI.RegisterService(callCtx, name).RegisterRequest(request).Execute()
	})
	return response, err
}

// Patch 更新服务的 apiInfo / projectRoot / probe。
func (c *Client) Patch(ctx context.Context, name string, opts PatchOptions) (*gatewayapi.RegistrationResponse, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	var request gatewayapi.PatchRequest
	changed := false
	if opts.APIInfo != nil {
		request.ApiInfo = opts.APIInfo
		changed = true
	}
	if opts.ProjectRoot != nil {
		request.ProjectRoot = *gatewayapi.NewNullableString(opts.ProjectRoot)
		changed = true
	}
	if opts.Probe != nil {
		probe, err := newProbe(*opts.Probe)
		if err != nil {
			return nil, err
		}
		request.Probe = *gatewayapi.NewNullableProbeConfig(probe)
		changed = true
	}
	if opts.CORS != nil {
		if err := opts.CORS.validate(); err != nil {
			return nil, err
		}
		model := opts.CORS.toModel()
		request.Cors = *gatewayapi.NewNullableCorsConfig(&model)
		changed = true
	}
	if !changed {
		return nil, ErrEmptyPatch
	}

	response, _, err := callWithToken(c, ctx, func(api *gatewayapi.APIClient, callCtx context.Context) (*gatewayapi.RegistrationResponse, *http.Response, error) {
		return api.ServicesAPI.PatchService(callCtx, name).PatchRequest(request).Execute()
	})
	return response, err
}

// SetCORS 设置服务的跨域配置（PUT /gw/services/{name}/cors）。
//
// cors.AllowAll 为 true 时一键放行所有域。
func (c *Client) SetCORS(ctx context.Context, name string, cors Cors) (*gatewayapi.CorsResponse, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := cors.validate(); err != nil {
		return nil, err
	}
	model := cors.toModel()
	response, _, err := callWithToken(c, ctx, func(api *gatewayapi.APIClient, callCtx context.Context) (*gatewayapi.CorsResponse, *http.Response, error) {
		return api.ServicesAPI.PutServiceCors(callCtx, name).CorsConfig(model).Execute()
	})
	return response, err
}

// ClearCORS 清空服务的跨域配置（DELETE /gw/services/{name}/cors），
// 之后网关恢复透传 provider 的 CORS 响应头。
func (c *Client) ClearCORS(ctx context.Context, name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	_, _, err := callWithToken(c, ctx, func(api *gatewayapi.APIClient, callCtx context.Context) (struct{}, *http.Response, error) {
		response, err := api.ServicesAPI.DeleteServiceCors(callCtx, name).Execute()
		return struct{}{}, response, err
	})
	return err
}

// Delete 注销服务。
func (c *Client) Delete(ctx context.Context, name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	_, _, err := callWithToken(c, ctx, func(api *gatewayapi.APIClient, callCtx context.Context) (struct{}, *http.Response, error) {
		response, err := api.ServicesAPI.UnregisterService(callCtx, name).Execute()
		return struct{}{}, response, err
	})
	return err
}

// callWithToken 在首次写操作前握手，并在 401 时重新握手后恰好重试一次。
//
// 如果首次调用因连接失败而失败（响应为 nil，通常是网关重启或换端口），
// 会重新读取 PortReader 刷新基地址后重试一次。
func callWithToken[T any](
	client *Client,
	ctx context.Context,
	call func(*gatewayapi.APIClient, context.Context) (T, *http.Response, error),
) (T, *http.Response, error) {
	value, response, err := invoke(client, ctx, call)
	if err == nil || response != nil || ctx.Err() != nil {
		return value, response, err
	}

	changed, refreshErr := client.refreshBaseURL()
	if refreshErr != nil || !changed {
		return value, response, err
	}
	return invoke(client, ctx, call)
}

// invoke 执行单次带 token 的调用；401 时重新握手并重试一次。
func invoke[T any](
	client *Client,
	ctx context.Context,
	call func(*gatewayapi.APIClient, context.Context) (T, *http.Response, error),
) (T, *http.Response, error) {
	var zero T

	token, err := client.accessToken(ctx)
	if err != nil {
		return zero, nil, err
	}
	value, response, err := call(client.currentAPI(), withToken(ctx, token))
	if response == nil || response.StatusCode != http.StatusUnauthorized {
		return value, response, err
	}
	if response.Body != nil {
		_ = response.Body.Close()
	}

	// 401：token 可能过期或网关重启。重新握手后只重试一次。
	token, herr := client.refreshToken(ctx)
	if herr != nil {
		return zero, response, herr
	}
	return call(client.currentAPI(), withToken(ctx, token))
}

func (c *Client) accessToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token == "" {
		if err := c.handshakeLocked(ctx); err != nil {
			return "", err
		}
	}
	return c.token, nil
}

func (c *Client) refreshToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = ""
	if err := c.handshakeLocked(ctx); err != nil {
		return "", err
	}
	return c.token, nil
}

// handshakeLocked 调用 POST /gw/session；调用方必须持有 c.mu。
func (c *Client) handshakeLocked(ctx context.Context) error {
	session, _, err := c.currentAPI().GatewayAPI.CreateSession(ctx).Execute()
	if err != nil {
		return fmt.Errorf("localgateway: session handshake: %w", err)
	}
	if session == nil || session.Token == "" {
		return errors.New("localgateway: session handshake returned an empty token")
	}
	c.token = session.Token
	return nil
}

func withToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, gatewayapi.ContextAccessToken, token)
}

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidName
	}
	return nil
}

func newProbe(probe Probe) (*gatewayapi.ProbeConfig, error) {
	if !strings.HasPrefix(probe.Path, "/") {
		return nil, fmt.Errorf("%w: got %q", ErrInvalidProbe, probe.Path)
	}
	config := gatewayapi.NewProbeConfig(probe.Path)
	if probe.TimeoutMs != nil {
		config.TimeoutMs = probe.TimeoutMs
	}
	return config, nil
}
