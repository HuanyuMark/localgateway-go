//go:build !windows

package localgateway

// ReadPort 在非 Windows 平台不可用；请通过 Options.PortReader 注入端口。
func ReadPort() (uint16, error) {
	return 0, ErrPortUnsupported
}
