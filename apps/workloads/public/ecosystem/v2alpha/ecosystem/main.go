package main

import (
	"context"
	"fmt"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/vanguard"

	accountauthority "apps/workloads/public/ecosystem/v2alpha/ecosystem/account-authority"
	configuration "apps/workloads/public/ecosystem/v2alpha/ecosystem/configuration"
	natsnodev2 "libs/partner/go/nats/v2"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	opentelemetryv2 "libs/partner/go/opentelemetry/v2"
	advertisementv1pbconnect "libs/partner/go/protobuf/gen/kevel/advertisement/v1/advertisementv1pbconnect"
	protovalidatev0 "libs/partner/go/protovalidate/v0"
	advertisementv1pbsrv "libs/partner/go/server/v2alpha/gen/kevel/advertisement/v1"
	zaploggerv1 "libs/partner/go/zap/v1"
	configurationv2alphalib "libs/private/go/configuration/v2alpha"
	configurationv2alphapbconnect "libs/public/go/protobuf/gen/platform/configuration/v2alpha/configurationv2alphapbconnect"
	iamv2alphapbconnect "libs/public/go/protobuf/gen/platform/iam/v2alpha/iamv2alphapbconnect"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	serverv2alphalib "libs/public/go/server/v2alpha"
	configurationv2alphapbsrv "libs/public/go/server/v2alpha/gen/platform/configuration/v2alpha"
	iamv2alphapbsrv "libs/public/go/server/v2alpha/gen/platform/iam/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&protovalidatev0.Binding{},
		&opentelemetryv2.Binding{},
		&zaploggerv1.Binding{},
		&nebulav1ca.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
			&configuration.CreateConfigurationListener{},
			&configuration.GetConfigurationListener{},
			&accountauthority.CreateAccountAuthorityListener{},
		}},
		&configurationv2alphalib.Binding{},

		//&nebulav1.Binding{},
		// Add PushPin Server
		// Listen on outbound.channels and PushPin to Clients
		// Create a new Connector Listener and listen of outbound channels
	}

	provider, err := sdkv2alphalib.NewDotConfigSettingsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err2 := provider.WatchSettings(); err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

	if err = sdkv2alphalib.GlobalSystems.RegisterSystems(provider); err != nil {
		return
	}

	telemetry, _ := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	interceptors := connect.WithInterceptors(sdkv2alphalib.NewSpecInterceptor(), telemetry)

	var services []*vanguard.Service

	// TODO: Work on Dynamic Connector Handlers
	//for _, s := range sdkv2alphalib.GlobalSystems.GetSystems() {
	//	for _, c := range s.Connectors {
	//		services = append(services, vanguard.NewService(c.ProcedureName, sdkv2alphalib.NewDynamicConnectorHandler[any, any]((*sdkv2alphalib.Connector[any, any])(c), interceptors)))
	//	}
	//}

	services = append(services, vanguard.NewService(configurationv2alphapbconnect.NewConfigurationServiceHandler(&configurationv2alphapbsrv.ConfigurationServiceHandler{}, interceptors)))
	services = append(services, vanguard.NewService(iamv2alphapbconnect.NewAccountAuthorityServiceHandler(&iamv2alphapbsrv.AccountAuthorityServiceHandler{}, interceptors)))
	services = append(services, vanguard.NewService(advertisementv1pbconnect.NewDecisionServiceHandler(&advertisementv1pbsrv.DecisionServiceHandler{}, interceptors)))

	multiplexedServer := serverv2alphalib.NewMultiplexedServer(context.Background(), bounds, services)

	multiplexedServer.ListenAndServe()
}
