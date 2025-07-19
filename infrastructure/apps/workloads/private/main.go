package main

import (
	"context"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"

	"github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	bounds := []sdkv2betalib.Binding{}

	infrastructure := sdkv2betalib.NewInfrastructure(context.Background(), sdkv2betalib.WithInfrastructureBounds(bounds))

	// name := sdkv2betalib.ShortenString(cnf.App.EnvironmentName+"-"+cnf.App.Name, 63)

	// Create Config Store
	// Create DNS Records
	// Create

	infrastructure.Run(edge)
}
