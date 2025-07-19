package main

import (
	"context"

	multiplexer "infrastructure/apps/workloads/private/event/v2alpha/event-multiplexer"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

func main() {
	bounds := []sdkv2betalib.Binding{}

	infrastructure := sdkv2betalib.NewInfrastructure(context.Background(), sdkv2betalib.WithInfrastructureBounds(bounds))

	// name := sdkv2betalib.ShortenString(cnf.App.EnvironmentName+"-"+cnf.App.Name, 63)

	// Create Config Store
	// Create DNS Records
	// Create

	infrastructure.Run(multiplexer.EventInfrastructure)
}
