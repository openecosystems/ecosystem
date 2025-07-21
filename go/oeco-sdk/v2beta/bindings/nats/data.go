package natsnodev1

import (
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	specproto "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
)

// NatsSpecWrapper wrapper for Data and Error
type NatsSpecWrapper struct {
	SpecData  *specproto.SpecData
	SpecError sdkv2betalib.SpecError
}
