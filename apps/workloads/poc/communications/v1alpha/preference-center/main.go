package main

import (
	"context"
	natsnodev2 "libs/partner/go/nats/v2"
	nebulav1 "libs/partner/go/nebula/v1"
	zaploggerv1 "libs/partner/go/zap/v1"
	configurationv2alphalib "libs/private/go/configuration/v2alpha"
	connectorv2alphalib "libs/public/go/connector/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&nebulav1.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{}},
		&configurationv2alphalib.Binding{},
	}

	connector := connectorv2alphalib.NewConnector(context.Background(), bounds)
	connector.ListenAndProcess()

	//_ = []sdkv2alphalib.Binding{
	//  &zaploggerv1.Binding{},
	//  &sendgridcontactsv3.Binding{},
	//  &sendgridlistv3.Binding{},
	//  &natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
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
