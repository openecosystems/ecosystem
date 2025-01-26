package main

import (
	"context"
	natsnodev2 "libs/partner/go/nats/v2"

	//"libs/partner/go/nebula/v1"
	zaploggerv1 "libs/partner/go/zap/v1"
	configurationv2alphalib "libs/private/go/configuration/v2alpha"
	connectorv2alphalib "libs/public/go/connector/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		//&nebulav1.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
			&CreateConfigurationListener{},
			&GetConfigurationListener{},
		}},
		&configurationv2alphalib.Binding{},
	}

	connector := connectorv2alphalib.NewConnector(context.Background(), bounds)
	connector.ListenAndProcess()
}
