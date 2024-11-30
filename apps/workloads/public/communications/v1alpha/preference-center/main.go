package main

import (
	"libs/partner/go/nats/v2"
	"libs/partner/go/zap/v1"
	"libs/private/go/configuration/v2alpha"
	"libs/public/go/connector/v2alpha"
	"libs/public/go/sdk/v2alpha"
)

func main() {

	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{}},
		&configurationv2alphalib.Binding{},
	}

	connector := connectorv2alphalib.NewConnectorA(bounds)
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
