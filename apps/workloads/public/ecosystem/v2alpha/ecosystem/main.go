package main

import (
	"context"
	"fmt"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/vanguard"

	certificate "apps/workloads/public/ecosystem/v2alpha/ecosystem/certificate"
	configuration "apps/workloads/public/ecosystem/v2alpha/ecosystem/configuration"
	ecosystem "apps/workloads/public/ecosystem/v2alpha/ecosystem/ecosystem"
	iam "apps/workloads/public/ecosystem/v2alpha/ecosystem/iam"
	natsnodev2 "libs/partner/go/nats/v2"
	nebulav1 "libs/partner/go/nebula/v1"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	opentelemetryv2 "libs/partner/go/opentelemetry/v2"
	advertisementv1pbconnect "libs/partner/go/protobuf/gen/kevel/advertisement/v1/advertisementv1pbconnect"
	protovalidatev0 "libs/partner/go/protovalidate/v0"
	advertisementv1pbsrv "libs/partner/go/server/v2alpha/gen/kevel/advertisement/v1"
	zaploggerv1 "libs/partner/go/zap/v1"
	configurationv2alphalib "libs/private/go/configuration/v2alpha"
	configurationv2alphapbconnect "libs/public/go/protobuf/gen/platform/configuration/v2alpha/configurationv2alphapbconnect"
	ecosystemv2alphapbconnect "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha/ecosystemv2alphapbconnect"
	iamv2alphapbconnect "libs/public/go/protobuf/gen/platform/iam/v2alpha/iamv2alphapbconnect"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	serverv2alphalib "libs/public/go/server/v2alpha"
	configurationv2alphapbsrv "libs/public/go/server/v2alpha/gen/platform/configuration/v2alpha"
	ecosystemv2alphapbsrv "libs/public/go/server/v2alpha/gen/platform/ecosystem/v2alpha"
	iamv2alphapbsrv "libs/public/go/server/v2alpha/gen/platform/iam/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&protovalidatev0.Binding{},
		&opentelemetryv2.Binding{},
		&zaploggerv1.Binding{},
		&nebulav1ca.Binding{},
		&nebulav1.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
			&ecosystem.CreateEcosystemListener{},
			&configuration.CreateConfigurationListener{},
			&configuration.GetConfigurationListener{},
			//&accountauthority.CreateAccountAuthorityListener{},
			&certificate.SignCertificateListener{},
			&iam.CreateAccountListener{},
		}},
		&configurationv2alphalib.Binding{},

		// Add PushPin Server
		// Listen on outbound.channels and PushPin to Clients
		// Create a new Connector Listener and listen of outbound channels
	}

	provider, err := sdkv2alphalib.NewDotConfigSettingsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err = provider.WatchSettings(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = sdkv2alphalib.GlobalSystems.RegisterSystems(provider); err != nil {
		return
	}

	//for _, s := range sdkv2alphalib.GlobalSystems.GetSystems() {
	//	for _, c := range s.Connectors {
	//		services = append(services, vanguard.NewService(c.ProcedureName, sdkv2alphalib.NewDynamicConnectorHandler[any, any]((*sdkv2alphalib.Connector[any, any])(c), interceptors)))
	//	}
	//}

	telemetry, _ := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	interceptors := connect.WithInterceptors(sdkv2alphalib.NewSpecInterceptor(), telemetry)

	var publicServices []*vanguard.Service
	publicServices = append(publicServices, vanguard.NewService(iamv2alphapbconnect.NewAccountServiceHandler(&iamv2alphapbsrv.AccountServiceHandler{}, interceptors)))

	var meshServices []*vanguard.Service
	meshServices = append(meshServices, vanguard.NewService(ecosystemv2alphapbconnect.NewEcosystemServiceHandler(&ecosystemv2alphapbsrv.EcosystemServiceHandler{}, interceptors)))
	meshServices = append(meshServices, vanguard.NewService(configurationv2alphapbconnect.NewConfigurationServiceHandler(&configurationv2alphapbsrv.ConfigurationServiceHandler{}, interceptors)))
	meshServices = append(meshServices, vanguard.NewService(iamv2alphapbconnect.NewAccountServiceHandler(&iamv2alphapbsrv.AccountServiceHandler{}, interceptors)))
	meshServices = append(meshServices, vanguard.NewService(advertisementv1pbconnect.NewDecisionServiceHandler(&advertisementv1pbsrv.DecisionServiceHandler{}, interceptors)))

	multiplexedServer := serverv2alphalib.NewMultiplexedServer(context.Background(), bounds, meshServices, publicServices)

	multiplexedServer.ListenAndServe()
}
