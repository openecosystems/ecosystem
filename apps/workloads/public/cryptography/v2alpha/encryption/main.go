package main

import (
	"context"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

func main() {
	//bounds := []sdkv2alphalib.Binding{
	//	&opentelemetryv1.Binding{},
	//	&zaploggerv1.Binding{},
	//	&tinkv2.Binding{},
	//}
	//
	//provider, err := sdkv2alphalib.NewCredentialProvider()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//if err2 := provider.WatchConfigurations(); err2 != nil {
	//	fmt.Println("watch settings error: ", err2)
	//	os.Exit(1)
	//}

	// telemetry, _ := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	// interceptors := connect.WithInterceptors(telemetry)
	// path, handler := cryptographyv2alphapbconnect.NewEncryptionServiceHandler(&cryptographyv2alphasrv.EncryptionServiceHandler{}, interceptors)
	// server := sdkv2alphalib.NewServer(context.Background(), bounds, path, &handler)
	server := sdkv2alphalib.NewServer(context.Background(), nil)

	server.ListenAndServe()
}
