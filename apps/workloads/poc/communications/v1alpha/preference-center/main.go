package main

import (
	"context"

	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	nebulav1 "github.com/openecosystems/ecosystem/libs/partner/go/nebula"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	configurationv2alphalib "github.com/openecosystems/ecosystem/libs/private/go/configuration/v2alpha"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&nebulav1.Binding{},
		&natsnodev1.Binding{SpecEventListeners: []natsnodev1.SpecEventListener{}},
		&configurationv2alphalib.Binding{},
	}

	connector := sdkv2alphalib.NewConnector(context.Background(), bounds)
	connector.ListenAndProcess()

	//_ = []sdkv2alphalib.Binding{
	//  &zaploggerv1.Binding{},
	//  &sendgridcontactsv3.Binding{},
	//  &sendgridlistv3.Binding{},
	//  &natsnodev1.Binding{SpecEventListeners: []natsnodev1.SpecEventListener{
	//    //&listener.PreferenceCenterListener{},
	//  }},
	//  //&configurationv2alphalib.Binding{},
	//}
	//serviceHandler := &communicationv1alphapb.PreferenceCenterService{
	//	//QueryHandler:    &query.Handler{},
	//	//MutationHandler: &mutation.Handler{},
	//}
	//
	//telemetry, err := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	//if err != nil {
	//	fmt.Println("error initializing otelconnect interceptor" + err.Error())
	//}
	//
	//interceptors := connect.WithInterceptors(sdkv2alphalib.NewSpecInterceptor(), telemetry)
	//path, handler := communicationv1alphapbconnect.NewPreferenceCenterServiceHandler(serviceHandler, interceptors)
	//server := communicationv1alphapb.RegisterPreferenceCenterServiceSpecServer(bounds, path, &handler)
	//
	////subscribe.RegisterSubscriptionGroups()
	//server.ListenAndServe()
}
