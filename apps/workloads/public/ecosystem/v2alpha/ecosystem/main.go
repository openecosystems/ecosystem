package main

import (
	"context"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/vanguard"

	"github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/configuration"
	"github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/ecosystem"
	"github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/iam"
	internal "github.com/openecosystems/ecosystem/apps/workloads/public/ecosystem/v2alpha/ecosystem/internal"
	configurationv2alphalib "github.com/openecosystems/ecosystem/libs/partner/go/configuration/v2alpha"
	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	nebulav1 "github.com/openecosystems/ecosystem/libs/partner/go/nebula"
	nebulav1ca "github.com/openecosystems/ecosystem/libs/partner/go/nebula/ca"
	opentelemetryv1 "github.com/openecosystems/ecosystem/libs/partner/go/opentelemetry"
	protovalidatev0 "github.com/openecosystems/ecosystem/libs/partner/go/protovalidate"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	configurationv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/configuration/v2alpha"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/configuration/v2alpha/configurationv2alphapbconnect"
	ecosystemv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapbconnect "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/ecosystem/v2alpha/ecosystemv2alphapbconnect"
	iamv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/iam/v2alpha"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/iam/v2alpha/iamv2alphapbconnect"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&protovalidatev0.Binding{},
		&opentelemetryv1.Binding{},
		&zaploggerv1.Binding{},
		&nebulav1ca.Binding{},
		&nebulav1.Binding{},
		&natsnodev1.Binding{SpecEventListeners: []natsnodev1.SpecEventListener{
			&ecosystem.CreateEcosystemListener{},
			&configuration.CreateConfigurationListener{},
			&configuration.GetConfigurationListener{},
			//&accountauthority.CreateAccountAuthorityListener{},
			//&certificate.SignCertificateListener{},
			&iam.CreateAccountListener{},
		}},
		&configurationv2alphalib.Binding{},

		// Add PushPin Server
		// Listen on outbound.channels and PushPin to Clients
		// Create a new Connector Listener and listen of outbound channels
	}

	//for _, s := range sdkv2alphalib.GlobalSystems.GetSystems() {
	//	for _, c := range s.Connectors {
	//		services = append(services, vanguard.NewService(c.ProcedureName, sdkv2alphalib.NewDynamicConnectorHandler[any, any]((*sdkv2alphalib.Connector[any, any])(c), interceptors)))
	//	}
	//}

	telemetry, _ := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	interceptors := connect.WithInterceptors(sdkv2alphalib.NewSpecInterceptor(), telemetry)

	var publicServices []*vanguard.Service
	publicServices = append(publicServices, vanguard.NewService(iamv2alphapbconnect.NewAccountServiceHandler(&iamv2alphapb.AccountServiceHandler{}, interceptors)))

	var meshServices []*vanguard.Service
	meshServices = append(meshServices, vanguard.NewService(ecosystemv2alphapbconnect.NewEcosystemServiceHandler(&ecosystemv2alphapb.EcosystemServiceHandler{}, interceptors)))
	meshServices = append(meshServices, vanguard.NewService(configurationv2alphapbconnect.NewConfigurationServiceHandler(&configurationv2alphapb.ConfigurationServiceHandler{}, interceptors)))
	meshServices = append(meshServices, vanguard.NewService(iamv2alphapbconnect.NewAccountServiceHandler(&iamv2alphapb.AccountServiceHandler{}, interceptors)))

	c := &internal.Configuration{}
	//_, err := c.ResolveConfiguration()
	//if err != nil {
	//	fmt.Println("error resolving configuration: ", err)
	//	return
	//}
	//settings := c.GetConfiguration()

	multiplexedServer := sdkv2alphalib.NewServer(
		context.Background(),
		sdkv2alphalib.WithBounds(bounds),
		sdkv2alphalib.WithPublicServices(publicServices),
		sdkv2alphalib.WithMeshServices(meshServices),
		sdkv2alphalib.WithConfigurationProvider(c),
	)

	//if err := sdkv2alphalib.GlobalSystems.RegisterSystems(provider); err != nil {
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
