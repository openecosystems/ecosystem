package main

import (
	"context"
	"encoding/base64"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	bounds := []sdkv2alphalib.Binding{}

	infrastructure := sdkv2alphalib.NewInfrastructure(context.Background(), sdkv2alphalib.WithInfrastructureBounds(bounds))

	cnf := infrastructure.Config
	name := sdkv2alphalib.ShortenString(cnf.App.EnvironmentName+"-"+cnf.App.Name, 63)

	infrastructure.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		version := cfg.Require("version")
		publicKey := cfg.Require("publicKey")
		caCrt := cfg.Require("caCrt")
		hostCrt := cfg.Require("hostCrt")
		hostKey := cfg.Require("hostKey")

		script, err := vultr.NewStartupScript(ctx, name, &vultr.StartupScriptArgs{
			Name:   pulumi.String(name + "-startup-script"),
			Script: pulumi.String(base64.StdEncoding.EncodeToString([]byte(cloudinit(publicKey, caCrt, hostCrt, hostKey, version)))),
			Type:   pulumi.String("boot"),
		})
		if err != nil {
			return err
		}

		firewallGroup, err := vultr.NewFirewallGroup(ctx, "EcosystemMultiplexerInbound", &vultr.FirewallGroupArgs{
			Description: pulumi.String("Ecosystem Multiplexer Group"),
		})
		if err != nil {
			return err
		}

		_, err = vultr.NewFirewallRule(ctx, "22/tcp", &vultr.FirewallRuleArgs{
			FirewallGroupId: firewallGroup.ID(),
			Protocol:        pulumi.String("tcp"),
			IpType:          pulumi.String("v4"),
			Subnet:          pulumi.String("0.0.0.0"),
			SubnetSize:      pulumi.Int(0),
			Port:            pulumi.String("22"),
			Notes:           pulumi.String("22/tcp/v4"),
		})
		if err != nil {
			return err
		}

		// Only allow 4222 on Tun device
		_, err = vultr.NewFirewallRule(ctx, "4222/tcp", &vultr.FirewallRuleArgs{
			FirewallGroupId: firewallGroup.ID(),
			Protocol:        pulumi.String("tcp"),
			IpType:          pulumi.String("v4"),
			Subnet:          pulumi.String("0.0.0.0"),
			SubnetSize:      pulumi.Int(0),
			Port:            pulumi.String("4222"),
			Notes:           pulumi.String("4222/tcp/v4"),
		})
		if err != nil {
			return err
		}

		// Only allow 6477 from Fastly
		_, err = vultr.NewFirewallRule(ctx, "6477/tcp", &vultr.FirewallRuleArgs{
			FirewallGroupId: firewallGroup.ID(),
			Protocol:        pulumi.String("tcp"),
			IpType:          pulumi.String("v4"),
			Subnet:          pulumi.String("0.0.0.0"),
			SubnetSize:      pulumi.Int(0),
			Port:            pulumi.String("6477"),
			Notes:           pulumi.String("6477/tcp/v4"),
		})
		if err != nil {
			return err
		}

		// Only allow 7999 from Fastly
		_, err = vultr.NewFirewallRule(ctx, "7999/tcp", &vultr.FirewallRuleArgs{
			FirewallGroupId: firewallGroup.ID(),
			Protocol:        pulumi.String("tcp"),
			IpType:          pulumi.String("v4"),
			Subnet:          pulumi.String("0.0.0.0"),
			SubnetSize:      pulumi.Int(0),
			Port:            pulumi.String("7999"),
			Notes:           pulumi.String("7999/tcp/v4"),
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
			DdosProtection:    pulumi.Bool(false),
			DisablePublicIpv4: pulumi.Bool(false),
			EnableIpv6:        pulumi.Bool(false),
			FirewallGroupId:   firewallGroup.ID(),
			Hostname:          pulumi.String(name),
			Label:             pulumi.String(name),
			OsId:              pulumi.Int(2136),                                      // "Debian 12 x64 (bookworm)"
			ReservedIpId:      pulumi.String("624af82d-cd58-4c7a-b7fb-848affafd7fe"), // Multiplexer IP address 149.28.81.51
			Plan:              pulumi.String("vhp-1c-1gb-amd"),                       // AMD High Performance
			Region:            pulumi.String("lax"),
			ScriptId:          script.ID(),
			Tags: pulumi.StringArray{
				pulumi.String("system:event"),
				pulumi.String("language:golang"),
				pulumi.String("cycle:public"),
				pulumi.String("version:v2alpha"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
