package natsnodev1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"dario.cat/mergo"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	natsd "github.com/nats-io/nats-server/v2/server"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
)

// NatsdServerName defines the name of the NATS server.
// NatsdServerHost specifies the host address of the NATS server.
// NatsdServerPort represents the port number for the NATS server.
// NatsdServerJetstreamStoreDir sets the directory for JetStream data storage.
const (
	NatsdServerName              = "natsv2"
	NatsdServerHost              = "localhost"
	NatsdServerPort              = 4222
	NatsdServerJetstreamStoreDir = "./nats-jetstream-data"
)

// ResolvedConfiguration holds the resolved configuration for the application.
// NatsServers defines the default NATS server addresses.
var (
	ResolvedConfiguration *Configuration
	// NatsServers           = []string{"nats://127.0.0.1:4222"}
)

// Nats represents the configuration for NATS connectivity including mesh and specific connection options.
type Nats struct {
	Options nats.Options
}

// Natsd configures an embedded NATS server with customizable options for use in the application.
// Enabled specifies whether the embedded NATS server is active or not.
// Options defines the configuration settings for the embedded NATS server through natsd.Options.
type Natsd struct {
	Enabled bool
	Options natsd.Options
}

// EventStreamRegistry is a structure that holds a list of jetstream stream configurations.
type EventStreamRegistry struct {
	Streams []jetstream.StreamConfig
}

// Configuration represents the overall settings structure comprising NATS, NATS server options, and stream registry configurations.
type Configuration struct {
	App                 specv2pb.App `yaml:"app,omitempty"`
	Platform            specv2pb.Platform
	Nats                Nats
	Natsd               Natsd
	EventStreamRegistry EventStreamRegistry

	err error
}

// ResolveConfiguration resolves the binding's configuration using the default configuration as a base and assigns it.
func (b *Binding) ResolveConfiguration(opts ...sdkv2betalib.ConfigurationProviderOption) (*sdkv2betalib.Configurer, error) {
	var c Configuration
	configurer, err := sdkv2betalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	sdkv2betalib.Resolve(configurer, &c, b.GetDefaultConfiguration())
	b.configuration = &c
	ResolvedConfiguration = &c

	dsc := jetstream.StreamConfig{
		MaxMsgs:           -1,
		MaxBytes:          -1,
		Discard:           0,
		MaxAge:            9151516080000000000, // 290 years is the max Nats supports,
		MaxMsgsPerSubject: -1,
		MaxMsgSize:        -1,
		Storage:           0,
		Replicas:          1, // TODO: Review this default
		NoAck:             false,
		Duplicates:        60 * time.Second, //"2m0s"
		DenyDelete:        true,
		DenyPurge:         true,
		AllowRollup:       false,
	}

	var mergedJsc []jetstream.StreamConfig
	var errs []error
	for _, r := range b.configuration.EventStreamRegistry.Streams {
		if err := mergo.Merge(&r, dsc); err != nil {
			fmt.Println("SpecError merging nats stream configuration:", err)
			errs = append(errs, err)
		}
		mergedJsc = append(mergedJsc, r)
	}
	if len(errs) > 0 {
		fmt.Println("nats configuration error: ", errs)
		b.configuration.err = errors.Join(errs...)
	}

	c.EventStreamRegistry = EventStreamRegistry{mergedJsc}
	b.configuration.EventStreamRegistry = EventStreamRegistry{mergedJsc}
	ResolvedConfiguration.EventStreamRegistry = EventStreamRegistry{mergedJsc}

	return configurer, nil
}

// ValidateConfiguration checks if NATS and stream configurations are valid, returning an error if validation fails.
func (b *Binding) ValidateConfiguration() error {
	if !b.configuration.Natsd.Enabled {
		return nil
	}

	var errs []error

	//if len(b.configuration.Natsd.Options.LeafNode.Remotes) == 0 {
	//	errs = append(errs, errors.New(`missing leaf node remotes configuration. An example is:
	//natsd:
	// options:
	//   leafNode:
	//     remotes:
	//       - urls:
	//           scheme: "tls"
	//           host:   "connect.ngs.global"
	//         credentials: "./example.creds"`))
	//}

	for i, s := range b.configuration.EventStreamRegistry.Streams {
		if s.Name == "" {
			errs = append(errs, errors.New("missing stream name for item with index: "+strconv.Itoa(i)))
		}

		if len(s.Subjects) == 0 {
			errs = append(errs, errors.New("missing array of subjects for item with index: "+strconv.Itoa(i)))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// GetDefaultConfiguration returns the default configuration object for the Binding, including NATS and JetStream settings.
func (b *Binding) GetDefaultConfiguration() *Configuration {
	// cfg := sdkv2betalib.ResolvedConfiguration

	return &Configuration{
		Platform: specv2pb.Platform{
			Mesh: &specv2pb.Mesh{
				Enabled: false,
			},
		},
		Nats: Nats{
			Options: nats.Options{
				//Servers: NatsServers,
				//TLSConfig: &tls.Config{},
				//Dialer: &net.Dialer{
				//	Timeout:   0,
				//	Deadline:  time.Time{},
				//	LocalAddr: nil,
				//	KeepAliveConfig: net.KeepAliveConfig{
				//		Enable:   false,
				//		Idle:     0,
				//		Interval: 0,
				//		Count:    0,
				//	},
				//},
				//CustomDialer:                nil,
				Compression: false,
			},
		},
		Natsd: Natsd{
			Enabled: false,
			Options: natsd.Options{
				ServerName:   NatsdServerName,
				Host:         NatsdServerHost,
				Port:         NatsdServerPort,
				DontListen:   false,
				Trace:        true,
				Debug:        true,
				TraceVerbose: true,
				// Debug:      cfg.App.Debug,
				MaxConn: -1,
				MaxSubs: -1,
				//LeafNode: natsd.LeafNodeOpts{
				//	Remotes: nil,
				//},
				JetStream:              true,
				JetStreamMaxMemory:     -1,
				JetStreamMaxStore:      -1,
				StoreDir:               NatsdServerJetstreamStoreDir,
				DisableJetStreamBanner: true,
				// TLSConfig:              &tls.Config{},
				// AllowNonTLS:            true,
				// TLSHandshakeFirst:      true,
				Routes: []*url.URL{},
				//Cluster: natsd.ClusterOpts{
				// ConnectRetries: 3,
				//PoolSize: 3,
				//Compression: natsd.CompressionOpts{
				//	Mode:          "s2_auto",
				//	RTTThresholds: []time.Duration{10 * time.Millisecond, 50 * time.Millisecond, 100 * time.Millisecond},
				//},
				//},
			},
		},
		EventStreamRegistry: EventStreamRegistry{
			// Streams: mergedJsc,
		},
	}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (b *Binding) CreateConfiguration() (*Configuration, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfiguration() *Configuration {
	return b.configuration
}

// GetConfigurationBytes retrieves the configuration of the binding instance. Returns the configuration as an *Configuration.
func (b *Binding) GetConfigurationBytes() ([]byte, error) {
	byteArray, err := json.Marshal(*b.GetConfiguration()) //nolint:staticcheck
	if err != nil {
		fmt.Println("SpecError:", err)
		return nil, err
	}
	return byteArray, nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (b *Binding) WatchConfigurations(directories ...string) error {
	fmt.Println("Watch settings ecosystem internal directories:", directories)
	return nil
}
