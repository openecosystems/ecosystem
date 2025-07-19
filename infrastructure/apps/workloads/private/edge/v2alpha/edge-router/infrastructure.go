package main

import (
	"context"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"

	"github.com/pulumi/pulumi-fastly/sdk/v8/go/fastly"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func EdgeInfrastructure(ctx *pulumi.Context) error {
	//exampleConfigstore, err := fastly.NewConfigstore(ctx, "config", &fastly.ConfigstoreArgs{
	//	Name: pulumi.String("config"),
	//})
	//if err != nil {
	//	return err
	//}

	// name := "api.system." + cnf.App.EnvironmentName + ".oeco.cloud"
	name := "api.system.dev-1.oeco.cloud"

	exampleConfigstore, err := fastly.NewConfigstore(ctx, "example", &fastly.ConfigstoreArgs{
		Name: pulumi.String("my_config_store"),
	})
	if err != nil {
		return err
	}

	pkg, err := fastly.GetPackageHash(ctx, &fastly.GetPackageHashArgs{
		Filename: pulumi.StringRef("communication-edge-router.tar.gz"),
	}, nil)
	if err != nil {
		return err
	}

	_, err = fastly.NewServiceCompute(ctx, name, &fastly.ServiceComputeArgs{
		Name: pulumi.String(name + "1"),
		// Activate: pulumi.Bool(true),
		// Comment:  pulumi.String("Communication System Edge Router"),
		Domains: fastly.ServiceComputeDomainArray{
			&fastly.ServiceComputeDomainArgs{
				Name: pulumi.String(name),
			},
		},
		Package: &fastly.ServiceComputePackageArgs{
			Filename:       pulumi.String("communication-edge-router.tar.gz"),
			SourceCodeHash: pulumi.String(pkg.Hash),
		},
		ResourceLinks: fastly.ServiceComputeResourceLinkArray{
			&fastly.ServiceComputeResourceLinkArgs{
				Name:       pulumi.String("my_resource_link"),
				ResourceId: exampleConfigstore.ID(),
			},
		},
		ForceDestroy: pulumi.Bool(true),
		//ProductEnablement: &fastly.ServiceComputeProductEnablementArgs{
		//	Fanout:     pulumi.Bool(true),
		//	Websockets: pulumi.Bool(true),
		//},
		//ResourceLinks: fastly.ServiceComputeResourceLinkArray{
		//	&fastly.ServiceComputeResourceLinkArgs{
		//		Name:       pulumi.String("string"),
		//		ResourceId: pulumi.String("string"),
		//		LinkId:     pulumi.String("string"),
		//	},
		//},
		//VersionComment: pulumi.String("Managed with Pulumi: Organization: " + ctx.Organization() + "; Project: " + ctx.Project() + "; Stack: " + ctx.Stack()),
	})
	if err != nil {
		return err
	}
	return nil
}
