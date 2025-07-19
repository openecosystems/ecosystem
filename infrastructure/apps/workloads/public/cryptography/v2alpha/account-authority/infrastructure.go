package cryptography

import (
	"encoding/base64"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// CryptographyInfrastructure infrastructure for cryptography
func CryptographyInfrastructure(ctx *pulumi.Context) error {
	cfg := config.New(ctx, "")
	version := cfg.Require("version")
	publicKey := cfg.Require("publicKey")
	caCrt := cfg.Require("caCrt")
	hostCrt := cfg.Require("hostCrt")
	hostKey := cfg.Require("hostKey")
	name := "cryptography"

	script, err := vultr.NewStartupScript(ctx, name, &vultr.StartupScriptArgs{
		Name:   pulumi.String(name + "-startup-script"),
		Script: pulumi.String(base64.StdEncoding.EncodeToString([]byte(cloudinit(publicKey, caCrt, hostCrt, hostKey, version)))),
		Type:   pulumi.String("boot"),
	})
	if err != nil {
		return err
	}

	firewallGroup, err := vultr.NewFirewallGroup(ctx, "AccountAuthorityInbound", &vultr.FirewallGroupArgs{
		Description: pulumi.String("Account Authority Firewall Group"),
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
		OsId:              pulumi.Int(2136),                // "Debian 12 x64 (bookworm)"
		Plan:              pulumi.String("vhp-1c-1gb-amd"), // AMD High Performance
		Region:            pulumi.String("lax"),
		ScriptId:          script.ID(),
		Tags: pulumi.StringArray{
			pulumi.String("system:cryptography"),
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
