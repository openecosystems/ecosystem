package lighthouse

import (
	"encoding/base64"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// LighthouseInfrastructure infrastructure for lighthouse
func LighthouseInfrastructure(ctx *pulumi.Context) error {
	name := "lighthouse"

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

	firewallGroup, err := vultr.NewFirewallGroup(ctx, "LighthouseInbound", &vultr.FirewallGroupArgs{
		Description: pulumi.String("Lighthouse Firewall Group"),
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

	_, err = vultr.NewFirewallRule(ctx, "4242/udp", &vultr.FirewallRuleArgs{
		FirewallGroupId: firewallGroup.ID(),
		Protocol:        pulumi.String("udp"),
		IpType:          pulumi.String("v4"),
		Subnet:          pulumi.String("0.0.0.0"),
		SubnetSize:      pulumi.Int(0),
		Port:            pulumi.String("4242"),
		Notes:           pulumi.String("4242/udp/v4"),
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
		ReservedIpId:      pulumi.String("50e066c5-d640-4f7d-92da-6426c5b340cb"), // Lighthouse IP address 45.63.49.173
		Plan:              pulumi.String("vhp-1c-1gb-amd"),                       // AMD High Performance
		Region:            pulumi.String("lax"),
		ScriptId:          script.ID(),
		Tags: pulumi.StringArray{
			pulumi.String("system:mesh"),
			pulumi.String("language:golang"),
			pulumi.String("cycle:public"),
			pulumi.String("version:v2alpha"),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
