package main

import (
	"context"
	"libs/partner/go/nats/v2"
	nebulav1 "libs/partner/go/nebula/v1"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	"libs/partner/go/zap/v1"
	"libs/private/go/configuration/v2alpha"
	"libs/public/go/connector/v2alpha"
	"libs/public/go/sdk/v2alpha"
)

func main() {
	bounds := []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&nebulav1.Binding{},
		&nebulav1ca.Binding{},
		&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
			&CreateCertificateAuthorityListener{},
		}},
		&configurationv2alphalib.Binding{},
	}

	connector := connectorv2alphalib.NewConnector(context.Background(), bounds)
	connector.ListenAndProcess()
}
