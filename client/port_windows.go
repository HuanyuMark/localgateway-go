//go:build windows

package localgateway

import (
	"errors"
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// ReadPort 读取配置子键下的 port 值。
//
// 子键默认是 HKCU\Software\LocalGateway；可用环境变量
// LOCALGATEWAY_CONFIG_REGISTRY 重定向（规则与网关一致）。
// 优先接受 REG_DWORD，也兼容 REG_SZ / REG_EXPAND_SZ 的十进制写法。
func ReadPort() (uint16, error) {
	subkey := registrySubkey()
	key, err := registry.OpenKey(registry.CURRENT_USER, subkey, registry.QUERY_VALUE)
	if err != nil {
		if errors.Is(err, registry.ErrNotExist) {
			return 0, portNotConfiguredError(subkey)
		}
		return 0, fmt.Errorf(`localgateway: open HKCU\%s: %w`, subkey, err)
	}
	defer key.Close()

	// 端口值的实际长度很短；用固定缓冲区即可覆盖字符串写法。
	buffer := make([]byte, 64*1024)
	size, kind, err := key.GetValue("port", buffer)
	if err != nil {
		if errors.Is(err, registry.ErrNotExist) {
			return 0, portNotConfiguredError(subkey)
		}
		return 0, fmt.Errorf(`localgateway: read HKCU\%s\port: %w`, subkey, err)
	}
	return parsePortValue(kind, buffer[:size])
}

// portNotConfiguredError 保留 ErrPortNotConfigured 的 errors.Is 语义，
// 同时在错误信息里带上实际使用的子键。
func portNotConfiguredError(subkey string) error {
	return fmt.Errorf(`%w: HKCU\%s\port`, ErrPortNotConfigured, subkey)
}
