package pushpinv1

import sdkv2alphalib "libs/public/go/sdk/v2alpha"

var ResolvedConfiguration *Configuration

type Pki struct {
	Ca   string `json:"ca,omitempty" yaml:"ca,omitempty"`
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key  string `json:"key,omitempty" yaml:"key,omitempty"`
}

type Lighthouse struct {
	AmLighthouse bool     `json:"am_lighthouse" yaml:"am_lighthouse"`
	Interval     int      `json:"interval,omitempty" yaml:"interval,omitempty"`
	Hosts        []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`
}

type Listen struct {
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	Port int    `json:"port,omitempty" yaml:"port,omitempty"`
}

type Punchy struct {
	Punch        bool   `json:"punch,omitempty" yaml:"punch,omitempty"`
	Respond      bool   `json:"respond,omitempty" yaml:"respond,omitempty"`
	RespondDelay string `json:"respond_delay" yaml:"respond_delay"`
	Delay        string `json:"delay,omitempty" yaml:"delay,omitempty"`
}

type Relay struct {
	AmRelay   bool `json:"am_relay,omitempty" yaml:"am_relay,omitempty"`
	UseRelays bool `json:"use_relays,omitempty" yaml:"use_relays,omitempty"`
}

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

type Logging struct {
	Level  string `json:"level,omitempty" yaml:"level,omitempty"`
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}

type Firewall struct {
	OutboundAction string `json:"outbound_action,omitempty" yaml:"outbound_action,omitempty"`
	InboundAction  string `json:"inbound_action,omitempty" yaml:"inbound_action,omitempty"`

	Conntrack Conntrack `json:"conntrack,omitempty" yaml:"conntrack,omitempty"`
	Outbound  Outbound  `json:"outbound,omitempty" yaml:"outbound,omitempty"`
	Inbound   Inbound   `json:"inbound,omitempty" yaml:"inbound,omitempty"`
}

type Conntrack struct {
	TCPTimeout     string `json:"tcp_timeout,omitempty" yaml:"tcp_timeout,omitempty"`
	UDPTimeout     string `json:"udp_timeout,omitempty" yaml:"udp_timeout,omitempty"`
	DefaultTimeout string `json:"default_timeout,omitempty" yaml:"default_timeout,omitempty"`
}

type Outbound []struct {
	Port  string `json:"port,omitempty" yaml:"port,omitempty"`
	Proto string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Host  string `json:"host,omitempty" yaml:"host,omitempty"`
}

type Inbound []struct {
	Port      string   `json:"port,omitempty" yaml:"port,omitempty"`
	Proto     string   `json:"proto,omitempty" yaml:"proto,omitempty"`
	Host      string   `json:"host,omitempty" yaml:"host,omitempty"`
	Groups    []string `json:"groups,omitempty" yaml:"groups,omitempty"`
	Group     string   `json:"group,omitempty" yaml:"group,omitempty"`
	LocalCidr string   `json:"local_cidr,omitempty" yaml:"local_cidr,omitempty"`
}

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

type Configuration struct {
	Nebula Nebula `json:"nebula" yaml:"nebula"`
}

func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

func (b *Binding) ValidateConfiguration() error {
	return nil
}

func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{}
}
