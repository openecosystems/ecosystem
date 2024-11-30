package main

import (
	"context"
	"libs/partner/go/nats/v2"
	"libs/partner/go/zap/v1"
	"libs/private/go/configuration/v2alpha"
	"libs/public/go/connector/v2alpha"
	"libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
			&CreateConfigurationListener{},
			&GetConfigurationListener{},
		}},
		&configurationv2alphalib.Binding{},
	}

	connector := connectorv2alphalib.NewConnector(context.Background(), bounds)
	connector.ListenAndProcess()
}
