package main

import (
	"encoding/base64"
	"libs/private/go/infrastructure/v2alpha"
	"libs/public/go/sdk/v2alpha"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
		// Open firewall to listen on 6477 (Must be open to the edge Router)
		// 4222 will be on the mesh
		// Add support for mTLS with connection to Edge Router

		script, err := vultr.NewStartupScript(ctx, name, &vultr.StartupScriptArgs{
			Name:   pulumi.String(name + "-startup-script"),
			Script: pulumi.String(base64.StdEncoding.EncodeToString([]byte(CloudInitStartupScript))),
			Type:   pulumi.String("boot"),
		})
		if err != nil {
			return err
		}

		firewallGroup, err := vultr.NewFirewallGroup(ctx, "EventListenerInbound", &vultr.FirewallGroupArgs{
			Description: pulumi.String("Event Listener Firewall Group"),
		})
		if err != nil {
			return err
		}
		_, err = vultr.NewFirewallRule(ctx, "4242/udp", &vultr.FirewallRuleArgs{
			FirewallGroupId: firewallGroup.ID(),
			Protocol:        pulumi.String("udp"),
			IpType:          pulumi.String("v4"),
			Subnet:          pulumi.String("0.0.0.0"),
			SubnetSize:      pulumi.Int(0),
			Port:            pulumi.String("4242"),
			Notes:           pulumi.String("4222/udp/v4"),
		})
		if err != nil {
			return err
		}

		_, err = vultr.NewFirewallRule(ctx, "ICMP", &vultr.FirewallRuleArgs{
			FirewallGroupId: firewallGroup.ID(),
			Protocol:        pulumi.String("icmp"),
			IpType:          pulumi.String("v4"),
			Subnet:          pulumi.String("0.0.0.0"),
			SubnetSize:      pulumi.Int(0),
			Notes:           pulumi.String("ICMP/v4"),
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
			FirewallGroupId:   firewallGroup.ID(),
			Label:             pulumi.String(name),
			OsId:              pulumi.Int(2136),                //"Debian 12 x64 (bookworm)"
			Plan:              pulumi.String("vhp-1c-1gb-amd"), // AMD High Performance
			Region:            pulumi.String("atl"),
			ScriptId:          script.ID(),
			Tags: pulumi.StringArray{
				pulumi.String("system:event"),
				pulumi.String("language:golang"),
				pulumi.String("cycle:private"),
				pulumi.String("version:v2alpha"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
