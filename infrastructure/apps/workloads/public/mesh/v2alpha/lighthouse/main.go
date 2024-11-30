package main

import (
	"encoding/base64"
	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"libs/private/go/infrastructure/v2alpha"
)

func main() {

	bounds := []sdkv2alphalib.Binding{}

	infrastructure := infrastructurev2alphalib.NewInfrastructure(bounds)

	config := infrastructure.Config
	name := infrastructurev2alphalib.ShortenString(config.App.EnvironmentName+"-"+config.App.Name, 63)

	infrastructure.Run(func(ctx *pulumi.Context) error {

		// Install Tun device
		// Route all traffic through Tun device
		// Start services
		// Open firewall to listen on 4242 (Must be open to the edge Router)
		// Add support for mTLS with connection to Edge Router
		// Create Server instance

		script, err := vultr.NewStartupScript(ctx, name, &vultr.StartupScriptArgs{
			Name:   pulumi.String(name + "-startup-script"),
			Script: pulumi.String(base64.StdEncoding.EncodeToString([]byte(CloudInitStartupScript))),
			Type:   pulumi.String("boot"),
		})
		if err != nil {
			return err
		}

		_, err = vultr.NewInstance(ctx, name, &vultr.InstanceArgs{
			ActivationEmail: pulumi.Bool(false),
			Backups:         pulumi.String("enabled"),
			BackupsSchedule: &vultr.InstanceBackupsScheduleArgs{
				Type: pulumi.String("daily"),
			},
			DdosProtection:    pulumi.Bool(true),
			DisablePublicIpv4: pulumi.Bool(false),
			EnableIpv6:        pulumi.Bool(true),
			Hostname:          pulumi.String(name),
			Label:             pulumi.String(name),
			OsId:              pulumi.Int(2136),                //"Debian 12 x64 (bookworm)"
			Plan:              pulumi.String("vhp-1c-1gb-amd"), //AMD High Performance
			Region:            pulumi.String("atl"),
			ScriptId:          script.ID(),
			Tags: pulumi.StringArray{
				pulumi.String("system:event"),
				pulumi.String("language:golang"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
