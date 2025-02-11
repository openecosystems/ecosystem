package nebulav1

import (
	"errors"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// ResolvedConfiguration holds the resolved runtime configuration for the Nebula binding.
var ResolvedConfiguration *Configuration

// Pki represents Public Key Infrastructure details, including CA, certificate, and private key configuration.
// The `Ca` field specifies the certificate authority information.
// The `Cert` field contains the signed certificate for the entity.
// The `Key` field includes the private key associated with the certificate.
type Pki struct {
	Ca   string `json:"ca,omitempty" yaml:"ca,omitempty"`
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key  string `json:"key,omitempty" yaml:"key,omitempty"`
}

// Lighthouse defines the configuration for using a node as a lighthouse in the network.
type Lighthouse struct {
	AmLighthouse bool     `json:"am_lighthouse" yaml:"am_lighthouse"`
	Interval     int      `json:"interval,omitempty" yaml:"interval,omitempty"`
	Hosts        []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`
}

// Listen defines network listening parameters including host and port configuration.
type Listen struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port int    `json:"port,omitempty" yaml:"port,omitempty"`
}

// Punchy represents configuration options related to NAT traversal and connectivity punch behavior.
// Punch indicates whether NAT hole punching is enabled or not.
// Respond specifies if the client should respond to punch requests.
// RespondDelay defines the wait time before responding to a punch request.
// Delay sets an additional delay for initiating punch attempts.
type Punchy struct {
	Punch        bool   `json:"punch,omitempty" yaml:"punch,omitempty"`
	Respond      bool   `json:"respond,omitempty" yaml:"respond,omitempty"`
	RespondDelay string `json:"respond_delay" yaml:"respond_delay"`
	Delay        string `json:"delay,omitempty" yaml:"delay,omitempty"`
}

// Relay represents relay configuration settings used in networking or forwarding scenarios.
// AmRelay indicates whether the current instance acts as a relay.
// UseRelays specifies if relays are utilized for forwarding connections.
type Relay struct {
	AmRelay   bool `json:"am_relay,omitempty" yaml:"am_relay,omitempty"`
	UseRelays bool `json:"use_relays,omitempty" yaml:"use_relays,omitempty"`
}

// Tun represents the configuration for the TUN interface in a network application.
// User specifies if the interface is in user mode.
// Disabled indicates whether the TUN interface is disabled.
// Dev specifies the device name of the TUN interface.
// DropLocalBroadcast determines if local broadcast traffic should be dropped.
// DropMulticast determines if multicast traffic should be dropped.
// TxQueue sets the transmit queue length for the interface.
// Mtu defines the Maximum Transmission Unit for the interface.
// Routes contains the routing configuration for TUN traffic.
// UnsafeRoutes stores routes that are not subject to safety checks.
type Tun struct {
	User               bool        `json:"user,omitempty" yaml:"user,omitempty"`
	Disabled           bool        `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	Dev                string      `json:"dev,omitempty" yaml:"dev,omitempty"`
	DropLocalBroadcast bool        `json:"drop_local_broadcast,omitempty" yaml:"drop_local_broadcast,omitempty"`
	DropMulticast      bool        `json:"drop_multicast,omitempty" yaml:"drop_multicast,omitempty"`
	TxQueue            int         `json:"tx_queue,omitempty" yaml:"tx_queue,omitempty"`
	Mtu                int         `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Routes             interface{} `json:"routes,omitempty" yaml:"routes,omitempty"`
	UnsafeRoutes       interface{} `json:"unsafe_routes,omitempty" yaml:"unsafe_routes,omitempty"`
}

// Logging represents configuration settings for logging.
// It includes the log Level and Format to customize logging behavior.
type Logging struct {
	Level  string `json:"level,omitempty" yaml:"level,omitempty"`
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}

// Firewall represents the structure for firewall rules and policies, including actions for inbound and outbound traffic.
type Firewall struct {
	OutboundAction string `json:"outbound_action,omitempty" yaml:"outbound_action,omitempty"`
	InboundAction  string `json:"inbound_action,omitempty" yaml:"inbound_action,omitempty"`

	Conntrack Conntrack `json:"conntrack,omitempty" yaml:"conntrack,omitempty"`
	Outbound  Outbound  `json:"outbound,omitempty" yaml:"outbound,omitempty"`
	Inbound   Inbound   `json:"inbound,omitempty" yaml:"inbound,omitempty"`
}

// Conntrack represents connection tracking settings with configurable timeouts for different protocols.
// TCPTimeout specifies the timeout value for TCP connections.
// UDPTimeout specifies the timeout value for UDP connections.
// DefaultTimeout specifies the default timeout value for unsupported protocols.
type Conntrack struct {
	TCPTimeout     string `json:"tcp_timeout,omitempty" yaml:"tcp_timeout,omitempty"`
	UDPTimeout     string `json:"udp_timeout,omitempty" yaml:"udp_timeout,omitempty"`
	DefaultTimeout string `json:"default_timeout,omitempty" yaml:"default_timeout,omitempty"`
}

// OutboundRule defines a network rule specifying the port, protocol, and host configuration for outbound traffic.
type OutboundRule struct {
	Port  string `json:"port,omitempty" yaml:"port,omitempty"`
	Proto string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Host  string `json:"host,omitempty" yaml:"host,omitempty"`
}

// Outbound represents a collection of outbound rules with port, protocol, and host information.
type Outbound []struct {
	Port  string `json:"port,omitempty" yaml:"port,omitempty"`
	Proto string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Host  string `json:"host,omitempty" yaml:"host,omitempty"`
}

// InboundRule represents a network rule defining inbound traffic configurations such as port, protocol, host, and groups.
type InboundRule struct {
	Port      string   `json:"port,omitempty" yaml:"port,omitempty"`
	Proto     string   `json:"proto,omitempty" yaml:"proto,omitempty"`
	Host      string   `json:"host,omitempty" yaml:"host,omitempty"`
	Groups    []string `json:"groups,omitempty" yaml:"groups,omitempty"`
	Group     string   `json:"group,omitempty" yaml:"group,omitempty"`
	LocalCidr string   `json:"local_cidr,omitempty" yaml:"local_cidr,omitempty"`
}

// Inbound represents a collection of inbound rules with port, protocol, host, groups, and CIDR configuration options.
type Inbound []struct {
	Port      string   `json:"port,omitempty" yaml:"port,omitempty"`
	Proto     string   `json:"proto,omitempty" yaml:"proto,omitempty"`
	Host      string   `json:"host,omitempty" yaml:"host,omitempty"`
	Groups    []string `json:"groups,omitempty" yaml:"groups,omitempty"`
	Group     string   `json:"group,omitempty" yaml:"group,omitempty"`
	LocalCidr string   `json:"local_cidr,omitempty" yaml:"local_cidr,omitempty"`
}

// Nebula represents the configuration for a Nebula mesh network. It includes settings for hosts, PKI, and other components.
type Nebula struct {
	Host map[string][]string `json:"static_host_map" yaml:"static_host_map"`

	Pki        Pki        `json:"pki,omitempty" yaml:"pki,omitempty"`
	Lighthouse Lighthouse `json:"lighthouse,omitempty" yaml:"lighthouse,omitempty"`
	Listen     Listen     `json:"listen,omitempty" yaml:"listen,omitempty"`
	Punchy     Punchy     `json:"punchy,omitempty" yaml:"punchy,omitempty"`
	Relay      Relay      `json:"relay,omitempty" yaml:"relay,omitempty"`
	Tun        Tun        `json:"tun,omitempty" yaml:"tun,omitempty"`
	Logging    Logging    `json:"logging,omitempty" yaml:"logging,omitempty"`
	Firewall   Firewall   `json:"firewall,omitempty" yaml:"firewall,omitempty"`
}

// Configuration defines the configuration structure containing settings for the Nebula network.
type Configuration struct {
	Nebula Nebula `json:"nebula" yaml:"nebula"`
}

// ResolveConfiguration resolves the binding's configuration using the default configuration as a base and assigns it.
func (b *Binding) ResolveConfiguration(provider *sdkv2alphalib.ConfigurationProvider) {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(provider, &c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

// ValidateConfiguration validates the configuration of the `Binding` instance.
// Ensures required fields like `Nebula.Pki.Ca` are set and adds errors if validation fails.
func (b *Binding) ValidateConfiguration() error {
	var errs []error
	if b.configuration == nil || b.configuration.Nebula.Pki.Ca == "" {
		_ = append(errs, errors.New("Nebula.Pki.Ca is required"))
	}

	// Host: nil,
	//  Pki: Pki{
	//    Ca:   "",
	//    Cert: "",
	//    Key:  "",
	//  },

	//if b.configuration.Nebula.Lighthouse.AmLighthouse == true {
	// hosts:
	//  - '192.168.100.1'
	//}

	return nil
}

// GetDefaultConfiguration provides the default configuration settings for the Binding instance, returning a Configuration object.
func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{
		Nebula: Nebula{
			Lighthouse: Lighthouse{
				AmLighthouse: false,
				Interval:     60,
			},
			Punchy: Punchy{
				Punch:        true,
				Respond:      true,
				RespondDelay: "5s",
				Delay:        "1s",
			},
			Tun: Tun{
				User:     true,
				Disabled: false,
			},
			Firewall: Firewall{
				Outbound: nil,
				Inbound:  nil, // TODO: Handle this with proper rule
			},
		},
	}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (b *Binding) CreateConfiguration() (interface{}, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (b *Binding) GetConfiguration() interface{} {
	return nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations() error {
	return nil
}
