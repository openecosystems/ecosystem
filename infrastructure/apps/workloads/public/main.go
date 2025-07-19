package main

import (
	"context"

	ecosystem "infrastructure/apps/workloads/public/ecosystem/v2alpha/ecosystem"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"
)

func main() {
	bounds := []sdkv2betalib.Binding{}

	infrastructure := sdkv2betalib.NewInfrastructure(context.Background(), sdkv2betalib.WithInfrastructureBounds(bounds))

	// cnf := *infrastructure.ConfigurationProvider
	// name := sdkv2betalib.ShortenString(cnf.App.EnvironmentName+"-"+cnf.App.Name, 63)

	infrastructure.Run(ecosystem.EcosystemInfrastructure)
}
