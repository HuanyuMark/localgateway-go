# localgateway-go

Go client SDK for the LocalGateway management API.

- `github.com/HuanyuMark/localgateway-go/client` (package `localgateway`) contains
  port discovery, the session handshake, register/PATCH/DELETE and retries;
- `github.com/HuanyuMark/localgateway-go/generated/gatewayapi` is the OpenAPI
  generated transport and must not be edited by hand.

## Usage

```powershell
go get github.com/HuanyuMark/localgateway-go@v0.1.0
```

```go
import localgateway "github.com/HuanyuMark/localgateway-go/client"

client, err := localgateway.New(localgateway.Options{})
registration, err := client.Register(ctx, "my-service", localgateway.RegisterOptions{
    Port:    8080,
    Version: "1.0.0",
    Probe:   localgateway.Probe{Path: "/health"},
})
```

Port discovery reads `HKCU\Software\LocalGateway\port` on Windows; inject
`Options.PortReader` (or pass `Options.BaseURL`) elsewhere.