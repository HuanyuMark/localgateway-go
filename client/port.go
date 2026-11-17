package localgateway

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf16"
)

// 注册表作用域相关常量。
const (
	// defaultRegistrySubkey 是 HKEY_CURRENT_USER 下的默认配置子键。
	defaultRegistrySubkey = `Software\LocalGateway`
	// registryConfigEnv 可把配置读写重定向到另一个子键；与网关和安装脚本同名。
	registryConfigEnv = "LOCALGATEWAY_CONFIG_REGISTRY"
)

// 可接受的注册表根前缀；大小写不敏感，按从长到短匹配。
var registryRootPrefixes = []string{`HKCU:\`, `HKCU\`, `HKEY_CURRENT_USER\`}

// registrySubkey 返回当前进程应使用的配置子键（相对 HKEY_CURRENT_USER）。
//
// 环境变量为空、纯空白或只剩根前缀时回退到 defaultRegistrySubkey。
func registrySubkey() string {
	return normalizeRegistrySubkey(os.Getenv(registryConfigEnv))
}

// normalizeRegistrySubkey 把环境变量值归一化为裸子键路径。
//
// 接受 `HKCU:\Software\...`、`HKCU\Software\...`、
// `HKEY_CURRENT_USER\Software\...` 与裸子键 `Software\...`；忽略首尾空白，
// 前缀大小写不敏感，空值回退到 defaultRegistrySubkey。
func normalizeRegistrySubkey(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return defaultRegistrySubkey
	}
	for _, prefix := range registryRootPrefixes {
		if len(trimmed) >= len(prefix) && strings.EqualFold(trimmed[:len(prefix)], prefix) {
			trimmed = trimmed[len(prefix):]
			break
		}
	}
	trimmed = strings.TrimLeft(trimmed, `\`)
	if trimmed == "" {
		return defaultRegistrySubkey
	}
	return trimmed
}

// 端口读取错误。
var (
	// ErrPortNotConfigured 表示注册表中缺少端口值；具体子键见错误信息。
	ErrPortNotConfigured = errors.New(`localgateway: gateway port is not configured in the registry`)
	// ErrPortInvalid 表示端口值不是 1..65535 的十进制整数。
	ErrPortInvalid = errors.New("localgateway: invalid gateway port")
	// ErrPortUnsupported 表示当前平台无法从注册表读取端口。
	ErrPortUnsupported = errors.New("localgateway: reading the gateway port from the registry is only supported on Windows")
)

// Windows 注册表值类型（与 windows/registry 的常量一致）。
const (
	registryTypeString       uint32 = 1 // REG_SZ
	registryTypeExpandString uint32 = 2 // REG_EXPAND_SZ
	registryTypeDWORD        uint32 = 4 // REG_DWORD
)

// parsePortValue 解析注册表原始值：优先 REG_DWORD，兼容十进制字符串。
func parsePortValue(kind uint32, data []byte) (uint16, error) {
	switch kind {
	case registryTypeDWORD:
		if len(data) < 4 {
			return 0, fmt.Errorf("%w: REG_DWORD is %d bytes, want 4", ErrPortInvalid, len(data))
		}
		value := binary.LittleEndian.Uint32(data[:4])
		if value == 0 || value > 65535 {
			return 0, fmt.Errorf("%w: %d is outside 1..65535", ErrPortInvalid, value)
		}
		return uint16(value), nil

	case registryTypeString, registryTypeExpandString:
		text := strings.TrimSpace(strings.Trim(decodeUTF16LE(data), "\x00"))
		if text == "" {
			return 0, fmt.Errorf("%w: empty string value", ErrPortInvalid)
		}
		value, err := strconv.ParseUint(text, 10, 16)
		if err != nil || value == 0 {
			return 0, fmt.Errorf("%w: %q is not a decimal port", ErrPortInvalid, text)
		}
		return uint16(value), nil

	default:
		return 0, fmt.Errorf("%w: unsupported registry value type %d", ErrPortInvalid, kind)
	}
}

func decodeUTF16LE(data []byte) string {
	if len(data)%2 != 0 {
		data = data[:len(data)-1]
	}
	units := make([]uint16, 0, len(data)/2)
	for i := 0; i+1 < len(data); i += 2 {
		units = append(units, binary.LittleEndian.Uint16(data[i:i+2]))
	}
	return string(utf16.Decode(units))
}
