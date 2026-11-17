package localgateway

import (
	"errors"
	"strings"
	"testing"
)

func TestNormalizeRegistrySubkey(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want string
	}{
		{name: "bare subkey", raw: `Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "hkcu colon", raw: `HKCU:\Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "hkcu colon lower", raw: `hkcu:\Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "hkcu backslash", raw: `HKCU\Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "hkey current user", raw: `HKEY_CURRENT_USER\Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "hkey current user lower", raw: `hkey_current_user\Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "surrounding whitespace", raw: "  Software\\LocalGateway\\tests\\a \t", want: `Software\LocalGateway\tests\a`},
		{name: "leading backslash", raw: `\Software\LocalGateway\tests\a`, want: `Software\LocalGateway\tests\a`},
		{name: "empty", raw: "", want: defaultRegistrySubkey},
		{name: "whitespace only", raw: " \t\r\n", want: defaultRegistrySubkey},
		{name: "root only colon", raw: `HKCU:\`, want: defaultRegistrySubkey},
		{name: "root only backslash", raw: `HKCU\`, want: defaultRegistrySubkey},
		{name: "root only hkey", raw: `HKEY_CURRENT_USER\`, want: defaultRegistrySubkey},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := normalizeRegistrySubkey(testCase.raw); got != testCase.want {
				t.Fatalf("normalizeRegistrySubkey(%q) = %q, want %q", testCase.raw, got, testCase.want)
			}
		})
	}
}

func TestRegistrySubkeyReadsEnvironment(t *testing.T) {
	t.Setenv(registryConfigEnv, `HKCU:\Software\LocalGateway\tests\env`)
	if got := registrySubkey(); got != `Software\LocalGateway\tests\env` {
		t.Fatalf("registrySubkey() = %q, want environment override", got)
	}

	t.Setenv(registryConfigEnv, "   ")
	if got := registrySubkey(); got != defaultRegistrySubkey {
		t.Fatalf("registrySubkey() with blank env = %q, want %q", got, defaultRegistrySubkey)
	}
}

func TestPortNotConfiguredErrorKeepsSentinel(t *testing.T) {
	err := portNotConfiguredError(`Software\LocalGateway\tests\x`)
	if !errors.Is(err, ErrPortNotConfigured) {
		t.Fatalf("error = %v, want errors.Is(ErrPortNotConfigured)", err)
	}
	if !strings.Contains(err.Error(), `HKCU\Software\LocalGateway\tests\x\port`) {
		t.Fatalf("error = %q, want actual subkey in message", err)
	}
}
