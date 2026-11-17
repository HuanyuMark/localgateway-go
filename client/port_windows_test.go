//go:build windows

package localgateway

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/sys/windows/registry"
)

// sandboxSubkey 返回一个本次测试独有的沙箱子键，绝不触碰真实配置值。
func sandboxSubkey(t *testing.T, suffix string) string {
	t.Helper()
	random := make([]byte, 4)
	if _, err := rand.Read(random); err != nil {
		t.Fatalf("generate sandbox suffix: %v", err)
	}
	return fmt.Sprintf(
		`Software\LocalGateway\tests\go-sdk-%s%d-%s`,
		suffix,
		os.Getpid(),
		hex.EncodeToString(random),
	)
}

func TestReadPortFollowsConfiguredRegistryScope(t *testing.T) {
	subkey := sandboxSubkey(t, "")
	key, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		subkey,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	if err != nil {
		t.Fatalf("create sandbox key HKCU\\%s: %v", subkey, err)
	}
	t.Cleanup(func() {
		_ = key.Close()
		if err := registry.DeleteKey(registry.CURRENT_USER, subkey); err != nil && !errors.Is(err, registry.ErrNotExist) {
			t.Errorf("delete sandbox key HKCU\\%s: %v", subkey, err)
		}
	})

	const wantPort = 54321
	if err := key.SetDWordValue("port", wantPort); err != nil {
		t.Fatalf("write sandbox port: %v", err)
	}

	// 使用带根前缀和空白的形式，同时验证环境变量归一化。
	t.Setenv(registryConfigEnv, `HKCU:\`+subkey+"  ")

	got, err := ReadPort()
	if err != nil {
		t.Fatalf("ReadPort() error = %v", err)
	}
	if got != wantPort {
		t.Fatalf("ReadPort() = %d, want %d", got, wantPort)
	}
}

func TestReadPortNotConfiguredIncludesSubkey(t *testing.T) {
	subkey := sandboxSubkey(t, "missing-")
	t.Setenv(registryConfigEnv, subkey)

	_, err := ReadPort()
	if !errors.Is(err, ErrPortNotConfigured) {
		t.Fatalf("ReadPort() error = %v, want errors.Is(ErrPortNotConfigured)", err)
	}
	if !strings.Contains(err.Error(), subkey) {
		t.Fatalf("ReadPort() error = %q, want actual subkey %q", err, subkey)
	}
}
