package main

import (
	"github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"libs/private/go/infrastructure/v2alpha"
)

func main() {

	bounds := []sdkv2alphalib.Binding{}

	infrastructure := infrastructurev2alphalib.NewInfrastructure(bounds)

	//config := infrastructure.Config
	//name := infrastructurev2alphalib.ShortenString(config.App.EnvironmentName+"-"+config.App.Name, 63)

	infrastructure.Run(func(ctx *pulumi.Context) error {

		_, err := fastly.NewServiceVcl(ctx, "myservice", &fastly.ServiceVclArgs{
			Name: pulumi.String("myawesometestservice"),
		})
		if err != nil {
			return err
		}
		return nil

	})
}
