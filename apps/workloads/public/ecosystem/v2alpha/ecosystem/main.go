package main

import (
	"context"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/vanguard"

	"github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/ecosystem"
	"github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/iam"
	internal "github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/internal"
	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"
	natsnodev1 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/nats"
	nebulav1 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/nebula"
	nebulav1ca "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/nebula/ca"
	opentelemetryv1 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/opentelemetry"
	protovalidatev0 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/protovalidate"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/zap"
	ecosystemv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapbconnect "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/gen/platform/ecosystem/v2alpha/ecosystemv2alphapbconnect"
	iamv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/gen/platform/iam/v2alpha"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/gen/platform/iam/v2alpha/iamv2alphapbconnect"
)

func main() {
	bounds := []sdkv2betalib.Binding{
		&protovalidatev0.Binding{},
		&opentelemetryv1.Binding{},
		&zaploggerv1.Binding{},
		&nebulav1ca.Binding{},
		&nebulav1.Binding{},
		&natsnodev1.Binding{SpecEventListeners: []natsnodev1.SpecEventListener{
			&ecosystem.CreateEcosystemListener{},
			//&accountauthority.CreateAccountAuthorityListener{},
			//&certificate.SignCertificateListener{},
			&iam.CreateAccountListener{},
		}},

		// Add PushPin Server
		// Listen on outbound.channels and PushPin to Clients
		// Create a new Connector Listener and listen of outbound channels
	}

	//for _, s := range sdkv2betalib.GlobalSystems.GetSystems() {
	//	for _, c := range s.Connectors {
	//		services = append(services, vanguard.NewService(c.ProcedureName, sdkv2betalib.NewDynamicConnectorHandler[any, any]((*sdkv2betalib.Connector[any, any])(c), interceptors)))
	//	}
	//}

	telemetry, _ := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	interceptors := connect.WithInterceptors(sdkv2betalib.NewSpecInterceptor(), telemetry)

	var publicServices []*vanguard.Service
	publicServices = append(publicServices, vanguard.NewService(iamv2alphapbconnect.NewAccountServiceHandler(&iamv2alphapb.AccountServiceHandler{}, interceptors)))

	var meshServices []*vanguard.Service
	meshServices = append(meshServices, vanguard.NewService(ecosystemv2alphapbconnect.NewEcosystemServiceHandler(&ecosystemv2alphapb.EcosystemServiceHandler{}, interceptors)))
	meshServices = append(meshServices, vanguard.NewService(iamv2alphapbconnect.NewAccountServiceHandler(&iamv2alphapb.AccountServiceHandler{}, interceptors)))

	c := &internal.Configuration{}
	//_, err := c.ResolveConfiguration()
	//if err != nil {
	//	fmt.Println("error resolving configuration: ", err)
	//	return
	//}
	//settings := c.GetConfiguration()

	multiplexedServer := sdkv2betalib.NewServer(
		context.Background(),
		sdkv2betalib.WithBounds(bounds),
		sdkv2betalib.WithPublicServices(publicServices),
		sdkv2betalib.WithMeshServices(meshServices),
		sdkv2betalib.WithConfigurationProvider(c),
	)

	//if err := sdkv2betalib.GlobalSystems.RegisterSystems(provider); err != nil {
	//	return
	//}

	//meshEndpoint := settings.Platform.Mesh.GetEndpoint()
	//ln, err3 := nebulav1.Bound.GetMeshListener(meshEndpoint)
	//if err3 != nil {
	//	fmt.Println("get socket error: ", err3)
	//}

	// multiplexedServer.ListenAndServeWithProvidedSocket(ln)
	multiplexedServer.ListenAndServe()
}
