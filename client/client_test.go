package localgateway

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"unicode/utf16"

	"github.com/HuanyuMark/localgateway-go/generated/gatewayapi"
)

func newTestClient(t *testing.T, opts Options) *Client {
	t.Helper()
	client, err := New(opts)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	return client
}

func writeJSON(t *testing.T, w http.ResponseWriter, status int, value any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		t.Errorf("encode response: %v", err)
	}
}

func sessionJSON(token string) map[string]any {
	return map[string]any{
		"token":     token,
		"expiresAt": "2026-11-14T12:00:00Z",
		"gateway": map[string]any{
			"baseUrl": "http://127.0.0.1:47821",
			"version": "0.1.0",
			"pid":     100,
		},
		"caller": map[string]any{
			"pid":       200,
			"userSid":   "S-1-5-21-1",
			"imagePath": `C:\provider.exe`,
		},
		"policy": map[string]any{
			"probeIntervalMs":    10000,
			"probeTimeoutMs":     2000,
			"offlineAfterMisses": 3,
			"expireAfterMs":      60000,
		},
	}
}

func registrationJSON(name string) map[string]any {
	return map[string]any{
		"name":     name,
		"status":   "healthy",
		"revision": 1,
	}
}

func TestRegisterPerformsHandshakeAndSendsDocumentedBody(t *testing.T) {
	var (
		mu            sync.Mutex
		sessionCalls  int
		registerCalls int
		authHeader    string
		body          map[string]any
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			mu.Lock()
			sessionCalls++
			mu.Unlock()
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-1"))
		case r.Method == http.MethodPut && r.URL.Path == "/gw/services/demo":
			raw, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("read body: %v", err)
			}
			mu.Lock()
			registerCalls++
			authHeader = r.Header.Get("Authorization")
			if err := json.Unmarshal(raw, &body); err != nil {
				t.Errorf("decode body: %v", err)
			}
			mu.Unlock()
			writeJSON(t, w, http.StatusCreated, registrationJSON("demo"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	timeout := int32(1500)
	client := newTestClient(t, Options{BaseURL: server.URL})
	response, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port:        18080,
		Version:     "1.2.3",
		Probe:       Probe{Path: "/health", TimeoutMs: &timeout},
		APIInfo:     map[string]any{"openapi": "3.1.0", "paths": map[string]any{}},
		ProjectRoot: `D:\demo`,
		NotifyURL:   "http://127.0.0.1:18080/events",
	})
	if err != nil {
		t.Fatalf("Register() error = %v", err)
	}
	if response.Name != "demo" || response.Status != gatewayapi.HEALTHY {
		t.Fatalf("unexpected response: %+v", response)
	}

	mu.Lock()
	defer mu.Unlock()
	if sessionCalls != 1 {
		t.Errorf("session calls = %d, want 1", sessionCalls)
	}
	if registerCalls != 1 {
		t.Errorf("register calls = %d, want 1", registerCalls)
	}
	if authHeader != "Bearer tok-1" {
		t.Errorf("Authorization = %q, want Bearer tok-1", authHeader)
	}
	if body["port"] != float64(18080) {
		t.Errorf("port = %v", body["port"])
	}
	if body["version"] != "1.2.3" {
		t.Errorf("version = %v", body["version"])
	}
	if body["protocol"] != "http" {
		t.Errorf("protocol = %v", body["protocol"])
	}
	probe, ok := body["probe"].(map[string]any)
	if !ok || probe["path"] != "/health" || probe["timeoutMs"] != float64(1500) {
		t.Errorf("probe = %v", body["probe"])
	}
	if body["projectRoot"] != `D:\demo` {
		t.Errorf("projectRoot = %v", body["projectRoot"])
	}
	if body["notifyUrl"] != "http://127.0.0.1:18080/events" {
		t.Errorf("notifyUrl = %v", body["notifyUrl"])
	}
	apiInfo, ok := body["apiInfo"].(map[string]any)
	if !ok || apiInfo["openapi"] != "3.1.0" {
		t.Errorf("apiInfo = %v", body["apiInfo"])
	}
}

func TestRegisterRetriesExactlyOnceAfterUnauthorized(t *testing.T) {
	var (
		mu           sync.Mutex
		sessionCalls int
		registerHits int
		seenTokens   []string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			mu.Lock()
			sessionCalls++
			token := "tok-" + string(rune('0'+sessionCalls))
			mu.Unlock()
			writeJSON(t, w, http.StatusOK, sessionJSON(token))
		case r.Method == http.MethodPut && r.URL.Path == "/gw/services/demo":
			mu.Lock()
			registerHits++
			token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			seenTokens = append(seenTokens, token)
			mu.Unlock()
			if token == "tok-1" {
				writeJSON(t, w, http.StatusUnauthorized, map[string]any{
					"code":    "unauthorized",
					"message": "token expired",
				})
				return
			}
			writeJSON(t, w, http.StatusOK, registrationJSON("demo"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})
	if _, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port:    18080,
		Version: "1.0.0",
		Probe:   Probe{Path: "/health"},
	}); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if sessionCalls != 2 {
		t.Errorf("session calls = %d, want 2", sessionCalls)
	}
	if registerHits != 2 {
		t.Errorf("register hits = %d, want 2", registerHits)
	}
	if len(seenTokens) != 2 || seenTokens[0] != "tok-1" || seenTokens[1] != "tok-2" {
		t.Errorf("tokens seen = %v, want [tok-1 tok-2]", seenTokens)
	}
}

func TestPersistentUnauthorizedDoesNotLoop(t *testing.T) {
	var (
		mu           sync.Mutex
		sessionCalls int
		registerHits int
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			sessionCalls++
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-"+string(rune('0'+sessionCalls))))
		case r.Method == http.MethodPut:
			registerHits++
			writeJSON(t, w, http.StatusUnauthorized, map[string]any{
				"code":    "unauthorized",
				"message": "still unauthorized",
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})
	_, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port:    18080,
		Version: "1.0.0",
		Probe:   Probe{Path: "/health"},
	})
	if err == nil {
		t.Fatal("Register() error = nil, want unauthorized error")
	}

	mu.Lock()
	defer mu.Unlock()
	if sessionCalls != 2 {
		t.Errorf("session calls = %d, want 2 (initial + exactly one retry)", sessionCalls)
	}
	if registerHits != 2 {
		t.Errorf("register hits = %d, want 2 (initial + exactly one retry)", registerHits)
	}
}

func TestPatchSendsOnlyProvidedFields(t *testing.T) {
	var (
		mu         sync.Mutex
		patchCalls int
		body       map[string]any
		authHeader string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-1"))
		case r.Method == http.MethodPatch && r.URL.Path == "/gw/services/demo":
			raw, _ := io.ReadAll(r.Body)
			mu.Lock()
			patchCalls++
			authHeader = r.Header.Get("Authorization")
			_ = json.Unmarshal(raw, &body)
			mu.Unlock()
			writeJSON(t, w, http.StatusOK, registrationJSON("demo"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})
	root := `D:\demo`
	_, err := client.Patch(context.Background(), "demo", PatchOptions{
		APIInfo:     map[string]any{"openapi": "3.0.3"},
		ProjectRoot: &root,
		Probe:       &Probe{Path: "/ready"},
	})
	if err != nil {
		t.Fatalf("Patch() error = %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if patchCalls != 1 {
		t.Errorf("patch calls = %d, want 1", patchCalls)
	}
	if authHeader != "Bearer tok-1" {
		t.Errorf("Authorization = %q", authHeader)
	}
	if body["projectRoot"] != `D:\demo` {
		t.Errorf("projectRoot = %v", body["projectRoot"])
	}
	probe, ok := body["probe"].(map[string]any)
	if !ok || probe["path"] != "/ready" {
		t.Errorf("probe = %v", body["probe"])
	}
	apiInfo, ok := body["apiInfo"].(map[string]any)
	if !ok || apiInfo["openapi"] != "3.0.3" {
		t.Errorf("apiInfo = %v", body["apiInfo"])
	}
}

func TestValidationErrorsDoNotSendRequests(t *testing.T) {
	var (
		mu       sync.Mutex
		requests int
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests++
		mu.Unlock()
		http.NotFound(w, r)
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})

	cases := []struct {
		name string
		call func() error
		want error
	}{
		{
			name: "missing probe",
			call: func() error {
				_, err := client.Register(context.Background(), "demo", RegisterOptions{
					Port: 18080, Version: "1.0.0",
				})
				return err
			},
			want: ErrInvalidProbe,
		},
		{
			name: "relative probe path",
			call: func() error {
				_, err := client.Register(context.Background(), "demo", RegisterOptions{
					Port: 18080, Version: "1.0.0", Probe: Probe{Path: "health"},
				})
				return err
			},
			want: ErrInvalidProbe,
		},
		{
			name: "zero port",
			call: func() error {
				_, err := client.Register(context.Background(), "demo", RegisterOptions{
					Port: 0, Version: "1.0.0", Probe: Probe{Path: "/health"},
				})
				return err
			},
			want: ErrInvalidPort,
		},
		{
			name: "empty version",
			call: func() error {
				_, err := client.Register(context.Background(), "demo", RegisterOptions{
					Port: 18080, Version: " ", Probe: Probe{Path: "/health"},
				})
				return err
			},
			want: ErrInvalidVersion,
		},
		{
			name: "empty name",
			call: func() error {
				_, err := client.Register(context.Background(), "", RegisterOptions{
					Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
				})
				return err
			},
			want: ErrInvalidName,
		},
		{
			name: "empty patch",
			call: func() error {
				_, err := client.Patch(context.Background(), "demo", PatchOptions{})
				return err
			},
			want: ErrEmptyPatch,
		},
		{
			name: "invalid patch probe",
			call: func() error {
				_, err := client.Patch(context.Background(), "demo", PatchOptions{
					Probe: &Probe{Path: "ready"},
				})
				return err
			},
			want: ErrInvalidProbe,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if err := testCase.call(); !errors.Is(err, testCase.want) {
				t.Fatalf("error = %v, want errors.Is(%v)", err, testCase.want)
			}
		})
	}

	mu.Lock()
	defer mu.Unlock()
	if requests != 0 {
		t.Errorf("HTTP requests = %d, want 0", requests)
	}
}

func TestInjectedPortReaderSuppliesBaseURL(t *testing.T) {
	var (
		mu           sync.Mutex
		readerCalls  int
		sessionCalls int
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			mu.Lock()
			sessionCalls++
			mu.Unlock()
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-1"))
		case r.Method == http.MethodPut && r.URL.Path == "/gw/services/demo":
			writeJSON(t, w, http.StatusCreated, registrationJSON("demo"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	port := uint16(server.Listener.Addr().(*net.TCPAddr).Port) // #nosec G115 -- tests use an ephemeral port.
	client := newTestClient(t, Options{
		PortReader: func() (uint16, error) {
			mu.Lock()
			readerCalls++
			mu.Unlock()
			return port, nil
		},
	})
	if _, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
	}); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if readerCalls != 1 {
		t.Errorf("port reader calls = %d, want 1", readerCalls)
	}
	if sessionCalls != 1 {
		t.Errorf("session calls = %d, want 1", sessionCalls)
	}
}

func TestTokenIsCachedAcrossCalls(t *testing.T) {
	var (
		mu           sync.Mutex
		sessionCalls int
		registerHits int
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			sessionCalls++
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-1"))
		case r.Method == http.MethodPut:
			registerHits++
			writeJSON(t, w, http.StatusOK, registrationJSON("demo"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})
	for range 2 {
		if _, err := client.Register(context.Background(), "demo", RegisterOptions{
			Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
		}); err != nil {
			t.Fatalf("Register() error = %v", err)
		}
	}

	mu.Lock()
	defer mu.Unlock()
	if sessionCalls != 1 {
		t.Errorf("session calls = %d, want 1", sessionCalls)
	}
	if registerHits != 2 {
		t.Errorf("register hits = %d, want 2", registerHits)
	}
}

func TestReadPortErrorPropagates(t *testing.T) {
	_, err := New(Options{
		PortReader: func() (uint16, error) { return 0, ErrPortNotConfigured },
	})
	if !errors.Is(err, ErrPortNotConfigured) {
		t.Fatalf("New() error = %v, want ErrPortNotConfigured", err)
	}
}

func TestParsePortValueAcceptsDWORDAndString(t *testing.T) {
	dword := make([]byte, 4)
	binary.LittleEndian.PutUint32(dword, 47821)

	cases := []struct {
		name string
		kind uint32
		data []byte
		want uint16
	}{
		{name: "dword", kind: registryTypeDWORD, data: dword, want: 47821},
		{name: "string", kind: registryTypeString, data: utf16LE("47821"), want: 47821},
		{name: "expand string", kind: registryTypeExpandString, data: utf16LE(" 41234\x00"), want: 41234},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := parsePortValue(testCase.kind, testCase.data)
			if err != nil {
				t.Fatalf("parsePortValue() error = %v", err)
			}
			if got != testCase.want {
				t.Fatalf("parsePortValue() = %d, want %d", got, testCase.want)
			}
		})
	}
}

func TestParsePortValueRejectsInvalidValues(t *testing.T) {
	zero := make([]byte, 4)
	tooLarge := make([]byte, 4)
	binary.LittleEndian.PutUint32(tooLarge, 70000)

	cases := []struct {
		name string
		kind uint32
		data []byte
	}{
		{name: "zero dword", kind: registryTypeDWORD, data: zero},
		{name: "large dword", kind: registryTypeDWORD, data: tooLarge},
		{name: "short dword", kind: registryTypeDWORD, data: []byte{1}},
		{name: "empty string", kind: registryTypeString, data: utf16LE("")},
		{name: "non numeric string", kind: registryTypeString, data: utf16LE("abc")},
		{name: "string zero", kind: registryTypeString, data: utf16LE("0")},
		{name: "string too large", kind: registryTypeString, data: utf16LE("70000")},
		{name: "unsupported type", kind: 7, data: []byte("47821")},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if _, err := parsePortValue(testCase.kind, testCase.data); !errors.Is(err, ErrPortInvalid) {
				t.Fatalf("parsePortValue() error = %v, want ErrPortInvalid", err)
			}
		})
	}
}

func utf16LE(value string) []byte {
	units := utf16.Encode([]rune(value))
	data := make([]byte, len(units)*2)
	for i, unit := range units {
		binary.LittleEndian.PutUint16(data[i*2:], unit)
	}
	return data
}

// deadServerURL 返回一个已经关闭的 httptest 地址，用于构造连接失败。
func deadServerURL(t *testing.T) string {
	t.Helper()
	server := httptest.NewServer(http.NotFoundHandler())
	url := server.URL
	server.Close()
	return url
}

// gatewayMock 是只处理 session 与单个写操作的网关桩。
type gatewayMock struct {
	mu           sync.Mutex
	sessionCalls int
	writeCalls   int
	server       *httptest.Server
}

func newGatewayMock(t *testing.T, method, path string) *gatewayMock {
	t.Helper()
	mock := &gatewayMock{}
	mock.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mock.mu.Lock()
		defer mock.mu.Unlock()
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			mock.sessionCalls++
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-refreshed"))
		case r.Method == method && r.URL.Path == path:
			mock.writeCalls++
			if method == http.MethodDelete {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			writeJSON(t, w, http.StatusOK, registrationJSON("demo"))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(mock.server.Close)
	return mock
}

func (m *gatewayMock) counts() (sessions, writes int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.sessionCalls, m.writeCalls
}

func TestSetBaseURLRetargetsRequestsAndDropsToken(t *testing.T) {
	first := newGatewayMock(t, http.MethodPut, "/gw/services/demo")
	second := newGatewayMock(t, http.MethodPut, "/gw/services/demo")

	client := newTestClient(t, Options{BaseURL: first.server.URL})
	if _, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
	}); err != nil {
		t.Fatalf("first Register() error = %v", err)
	}
	if client.BaseURL() != first.server.URL {
		t.Fatalf("BaseURL() = %q, want %q", client.BaseURL(), first.server.URL)
	}

	if err := client.SetBaseURL(second.server.URL + "/"); err != nil {
		t.Fatalf("SetBaseURL() error = %v", err)
	}
	if client.BaseURL() != second.server.URL {
		t.Fatalf("BaseURL() after SetBaseURL = %q, want %q", client.BaseURL(), second.server.URL)
	}

	if _, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
	}); err != nil {
		t.Fatalf("second Register() error = %v", err)
	}

	if _, writes := first.counts(); writes != 1 {
		t.Errorf("first gateway writes = %d, want 1", writes)
	}
	sessions, writes := second.counts()
	if sessions != 1 {
		t.Errorf("second gateway sessions = %d, want 1 (token must be dropped)", sessions)
	}
	if writes != 1 {
		t.Errorf("second gateway writes = %d, want 1", writes)
	}
}

func TestSetBaseURLRejectsInvalidURL(t *testing.T) {
	client := newTestClient(t, Options{BaseURL: "http://127.0.0.1:1"})
	if err := client.SetBaseURL("not a url"); !errors.Is(err, ErrInvalidBaseURL) {
		t.Fatalf("SetBaseURL() error = %v, want ErrInvalidBaseURL", err)
	}
}

func TestWriteOperationsRefreshBaseURLAfterConnectionFailure(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
		call   func(ctx context.Context, client *Client) error
	}{
		{
			name:   "register",
			method: http.MethodPut,
			path:   "/gw/services/demo",
			call: func(ctx context.Context, client *Client) error {
				_, err := client.Register(ctx, "demo", RegisterOptions{
					Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
				})
				return err
			},
		},
		{
			name:   "patch",
			method: http.MethodPatch,
			path:   "/gw/services/demo",
			call: func(ctx context.Context, client *Client) error {
				root := `D:\demo`
				_, err := client.Patch(ctx, "demo", PatchOptions{ProjectRoot: &root})
				return err
			},
		},
		{
			name:   "delete",
			method: http.MethodDelete,
			path:   "/gw/services/demo",
			call: func(ctx context.Context, client *Client) error {
				return client.Delete(ctx, "demo")
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			replacement := newGatewayMock(t, testCase.method, testCase.path)
			readerCalls := 0
			client := newTestClient(t, Options{
				BaseURL: deadServerURL(t),
				PortReader: func() (uint16, error) {
					readerCalls++
					return uint16(replacement.server.Listener.Addr().(*net.TCPAddr).Port), nil // #nosec G115 -- test port.
				},
			})

			if err := testCase.call(context.Background(), client); err != nil {
				t.Fatalf("%s error = %v", testCase.name, err)
			}
			if readerCalls != 1 {
				t.Errorf("PortReader calls = %d, want 1", readerCalls)
			}
			if client.BaseURL() != replacement.server.URL {
				t.Errorf("BaseURL() = %q, want %q", client.BaseURL(), replacement.server.URL)
			}
			sessions, writes := replacement.counts()
			if sessions != 1 || writes != 1 {
				t.Errorf("replacement sessions/writes = %d/%d, want 1/1", sessions, writes)
			}
		})
	}
}

func TestConnectionFailureWithoutPortReaderReturnsError(t *testing.T) {
	client := newTestClient(t, Options{BaseURL: deadServerURL(t)})
	if _, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"},
	}); err == nil {
		t.Fatal("Register() error = nil, want connection error")
	}
	if err := client.RefreshBaseURL(); !errors.Is(err, ErrNoPortReader) {
		t.Fatalf("RefreshBaseURL() error = %v, want ErrNoPortReader", err)
	}
}

func TestCORSIsSentOnRegisterAndManagedWithDedicatedEndpoints(t *testing.T) {
	var (
		mu             sync.Mutex
		registerBody   map[string]any
		corsBodies     []map[string]any
		corsMethods    []string
		corsAuthHeader string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gw/session":
			writeJSON(t, w, http.StatusOK, sessionJSON("tok-1"))
		case r.Method == http.MethodPut && r.URL.Path == "/gw/services/demo":
			raw, _ := io.ReadAll(r.Body)
			mu.Lock()
			_ = json.Unmarshal(raw, &registerBody)
			mu.Unlock()
			writeJSON(t, w, http.StatusCreated, registrationJSON("demo"))
		case r.URL.Path == "/gw/services/demo/cors":
			raw, _ := io.ReadAll(r.Body)
			var body map[string]any
			if len(raw) > 0 {
				_ = json.Unmarshal(raw, &body)
			}
			mu.Lock()
			corsMethods = append(corsMethods, r.Method)
			corsBodies = append(corsBodies, body)
			corsAuthHeader = r.Header.Get("Authorization")
			mu.Unlock()
			if r.Method == http.MethodDelete {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			writeJSON(t, w, http.StatusOK, map[string]any{
				"name":     "demo",
				"revision": 5,
				"cors":     body,
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})
	_, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port:    18080,
		Version: "1.2.3",
		Probe:   Probe{Path: "/health"},
		CORS:    &Cors{AllowAll: true},
	})
	if err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	response, err := client.SetCORS(context.Background(), "demo", Cors{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowCredentials: true,
		MaxAgeSeconds:    PtrInt64(600),
	})
	if err != nil {
		t.Fatalf("SetCORS() error = %v", err)
	}
	if got := response.Cors.Get().AllowOrigins; len(got) != 1 || got[0] != "http://localhost:5173" {
		t.Errorf("response allowOrigins = %v", got)
	}

	if err := client.ClearCORS(context.Background(), "demo"); err != nil {
		t.Fatalf("ClearCORS() error = %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	registered, ok := registerBody["cors"].(map[string]any)
	if !ok || registered["allowAll"] != true {
		t.Errorf("register cors = %v", registerBody["cors"])
	}
	if len(corsMethods) != 2 || corsMethods[0] != http.MethodPut || corsMethods[1] != http.MethodDelete {
		t.Errorf("cors methods = %v", corsMethods)
	}
	if corsAuthHeader != "Bearer tok-1" {
		t.Errorf("Authorization = %q", corsAuthHeader)
	}
	put := corsBodies[0]
	if put["allowCredentials"] != true || put["maxAgeSeconds"] != float64(600) {
		t.Errorf("put cors body = %v", put)
	}
	if origins, ok := put["allowOrigins"].([]any); !ok || len(origins) != 1 {
		t.Errorf("put allowOrigins = %v", put["allowOrigins"])
	}
}

func TestInvalidCORSDoesNotSendRequests(t *testing.T) {
	var (
		mu       sync.Mutex
		requests int
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests++
		mu.Unlock()
		http.NotFound(w, r)
	}))
	defer server.Close()

	client := newTestClient(t, Options{BaseURL: server.URL})
	if _, err := client.Register(context.Background(), "demo", RegisterOptions{
		Port: 18080, Version: "1.0.0", Probe: Probe{Path: "/health"}, CORS: &Cors{},
	}); !errors.Is(err, ErrInvalidCORS) {
		t.Fatalf("Register() error = %v, want ErrInvalidCORS", err)
	}
	if _, err := client.SetCORS(context.Background(), "demo", Cors{MaxAgeSeconds: PtrInt64(604_801)}); !errors.Is(err, ErrInvalidCORS) {
		t.Fatalf("SetCORS() error = %v, want ErrInvalidCORS", err)
	}
	if _, err := client.Patch(context.Background(), "demo", PatchOptions{CORS: &Cors{}}); !errors.Is(err, ErrInvalidCORS) {
		t.Fatalf("Patch() error = %v, want ErrInvalidCORS", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if requests != 0 {
		t.Errorf("requests = %d, want 0", requests)
	}
}
