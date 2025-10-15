package natsnodev1

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
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
	NatsServers           = []string{"nats://127.0.0.1:4222"}
)

// Nats represents the configuration for NATS connectivity including mesh and specific connection options.
type Nats struct {
	// TODO: Replace this dynamic accounts when we get to Super Clusters. Keep this simple for now until stable.
	Username string
	Password string
	Options  nats.Options
}

// Natsd configures an embedded NATS server with customizable options for use in the application.
// Enabled specifies whether the embedded NATS server is active or not.
// Options defines the configuration settings for the embedded NATS server through natsd.Options.
type Natsd struct {
	Enabled bool
	// Options    natsd.Options
	ServerName string
	Host       string
	Port       int
	HTTPPort   int
	RoutesStr  string
	Username   string
	Password   string
	Replicas   int
	Clustered  bool
	StoreDir   string
	Cluster    Cluster
}

type Cluster struct {
	Host string
	Port int
	Name string
}

// Configuration represents the overall settings structure comprising NATS, NATS server options, and stream registry configurations.
type Configuration struct {
	App      specv2pb.App `yaml:"app,omitempty"`
	Platform specv2pb.Platform
	Nats     Nats
	Natsd    Natsd

	err error
}

// ResolveConfiguration resolves the binding's configuration using the default configuration as a base and assigns it.
func (b *Binding) ResolveConfiguration(opts ...sdkv2betalib.ConfigurationProviderOption) (*sdkv2betalib.Configurer, error) {
	if b.configuration != nil {
		configurer, err := sdkv2betalib.NewConfigurer(opts...)
		if err != nil {
			return nil, err
		}
		return configurer, nil
	}

	var c Configuration
	configurer, err := sdkv2betalib.NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	sdkv2betalib.Resolve(configurer, &c, b.GetDefaultConfiguration())
	b.configuration = &c
	ResolvedConfiguration = &c

	return configurer, nil
}

// ValidateConfiguration checks if NATS and stream configurations are valid, returning an error if validation fails.
func (b *Binding) ValidateConfiguration() error {
	if !b.configuration.Natsd.Enabled {
		return nil
	}

	var errs []error

	if b.configuration.Natsd.Clustered {

		if b.configuration.Natsd.ServerName == "" {
			if podName, ok := os.LookupEnv("POD_NAME"); ok && podName != "" {
				b.configuration.Natsd.ServerName = podName
			} else {
				errs = append(errs, errors.New("Natsd.ServerName is required (missing and POD_NAME not set)"))
			}
		}

		if b.configuration.Natsd.Cluster.Host == "" {
			errs = append(errs, errors.New("Natsd.Cluster.Host is required"))
		}

		if b.configuration.Natsd.Cluster.Port == 0 {
			errs = append(errs, errors.New("Natsd.Cluster.Port is required"))
		}

		if b.configuration.Natsd.Cluster.Name == "" {
			errs = append(errs, errors.New("Natsd.Cluster.Name is required"))
		}

		if b.configuration.Natsd.Replicas == 0 {
			errs = append(errs, errors.New("Natsd.Replicas is required and must be greater than 0"))
		}

		if b.configuration.Natsd.RoutesStr == "" {
			errs = append(errs, errors.New("Natsd.RoutesStr is required"))
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
				Servers: NatsServers,
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
			//Enabled: false,
			//Options: natsd.Options{
			//	ServerName: NatsdServerName,
			//	Host:       NatsdServerHost,
			//	Port:       NatsdServerPort,
			//	DontListen: false,
			//	// Trace:      true,
			//	// Debug:      cfg.App.Debug,
			//	MaxConn: -1,
			//	MaxSubs: -1,
			//	LeafNode: natsd.LeafNodeOpts{
			//		Remotes: nil,
			//	},
			//	JetStream:              true,
			//	JetStreamMaxMemory:     -1,
			//	JetStreamMaxStore:      -1,
			//	StoreDir:               NatsdServerJetstreamStoreDir,
			//	DisableJetStreamBanner: true,
			//	// TLSConfig:              &tls.Config{},
			//	// AllowNonTLS:            true,
			//	// TLSHandshakeFirst:      true,
			//},
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
