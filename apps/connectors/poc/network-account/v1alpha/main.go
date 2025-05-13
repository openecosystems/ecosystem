package main

import (
	"context"

	internal "github.com/openecosystems/ecosystem/apps/connectors/poc/network-account/v1alpha/internal"
	configurationv2alphalib "github.com/openecosystems/ecosystem/libs/partner/go/configuration/v2alpha"
	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	nebulav1 "github.com/openecosystems/ecosystem/libs/partner/go/nebula"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&nebulav1.Binding{},
		&natsnodev1.Binding{SpecEventListeners: []natsnodev1.SpecEventListener{
			//&listeners.CreateEcosystemListener{},
		}},
		&configurationv2alphalib.Binding{},
	}

	c := &internal.Configuration{}

	connector := sdkv2alphalib.NewConnector(
		context.Background(),
		sdkv2alphalib.WithConnectorBounds(bounds),
		sdkv2alphalib.WithConnectorConfigurationProvider(c),
	)
	connector.ListenAndProcess()
}
