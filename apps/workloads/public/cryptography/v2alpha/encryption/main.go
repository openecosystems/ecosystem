package main

import (
	"context"
	"fmt"
	"libs/partner/go/opentelemetry/v2"
	"libs/partner/go/tink/v2"
	"libs/partner/go/zap/v1"
	"libs/public/go/protobuf/gen/platform/cryptography/v2alpha/cryptographyv2alphapbconnect"
	"libs/public/go/sdk/v2alpha"
	"libs/public/go/server/v2alpha"
	"libs/public/go/server/v2alpha/platform/cryptography/v2alpha"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&opentelemetryv2.Binding{},
		&zaploggerv1.Binding{},
		&tinkv2.Binding{},
	}

	provider, err := sdkv2alphalib.NewDotConfigSettingsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err2 := provider.WatchSettings(); err2 != nil {
		fmt.Println("watch settings error: ", err2)
		os.Exit(1)
	}

	telemetry, _ := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	interceptors := connect.WithInterceptors(telemetry)
	path, handler := cryptographyv2alphapbconnect.NewEncryptionServiceHandler(&cryptographyv2alphasrv.EncryptionServiceHandler{}, interceptors)
	server := serverv2alphalib.NewRawServer(context.Background(), bounds, path, &handler)

	server.ListenAndServe()
}
