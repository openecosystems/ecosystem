package sdkv2betalib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"connectrpc.com/connect"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
)

// CLI represents the core structure for managing CLI operations, bindings, and their configuration.
type CLI struct {
	Bindings              *Bindings
	Bounds                []Binding
	ConfigurationProvider BaseSpecConfigurationProvider

	options *cliOptions
}

// NewCLI initializes a new CLI instance by registering the provided bindings and returns the CLI object.
func NewCLI(ctx context.Context, opts ...CLIOption) *CLI {
	options, _ := newCLIOptions(opts)

	cli := CLI{
		Bounds: options.Bounds,

		options: options,
	}

	provider := options.ConfigurationProvider
	if provider == nil {
		panic("configuration provider is nil. Please provide a configuration provider to the server.")
	}

	cli.ConfigurationProvider = provider
	t := options.ConfigurationProvider

	var copts []ConfigurationProviderOption
	if options.RuntimeConfigurationOverrides.Overridden {
		copts = append(copts, WithRuntimeOverrides(options.RuntimeConfigurationOverrides))
	}

	configurer, err := t.ResolveConfiguration(copts...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = t.ValidateConfiguration()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	copts = append(copts, WithConfigurer(configurer))

	bindings := RegisterBindings(ctx, options.Bounds, copts...)
	cli.Bindings = bindings

	return &cli
}

// GracefulShutdown gracefully shuts down the CLI by cleaning up resources and closing bindings within a timeout context.
func (cli *CLI) GracefulShutdown() {
	_, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	ShutdownBindings(cli.Bindings)
}

// GetConfiguration retrieves and returns the configuration from the CLI's ConfigurationProvider if implemented.
func (cli *CLI) GetConfiguration() *CLIConfiguration {
	t := cli.ConfigurationProvider

	bytes, err := t.GetConfigurationBytes()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	c := CLIConfiguration{}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &c
}

// A CLIOption configures a [Server].
type CLIOption interface {
	apply(*cliOptions)
}

type cliOptions struct {
	URL     *url.URL
	MeshVPN bool

	Bounds                        []Binding
	PlatformContext               string
	RuntimeConfigurationOverrides RuntimeConfigurationOverrides
	ConfigurationProvider         BaseSpecConfigurationProvider
}

type cliOptionFunc func(*cliOptions)

func (f cliOptionFunc) apply(cfg *cliOptions) { f(cfg) }

//nolint:unparam
func newCLIOptions(options []CLIOption) (*cliOptions, *connect.Error) {
	// Defaults
	config := cliOptions{
		MeshVPN: false,
	}

	for _, opt := range options {
		opt.apply(&config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *cliOptions) validate() *connect.Error {
	return nil
}

// WithCLIOptions composes multiple Options into one.
func WithCLIOptions(opts ...CLIOption) CLIOption {
	return cliOptionFunc(func(cfg *cliOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithCLIBounds configures the server with the specified bounds, overriding the default bindings list in server options.
func WithCLIBounds(bounds []Binding) CLIOption {
	return cliOptionFunc(func(cfg *cliOptions) {
		cfg.Bounds = bounds
	})
}

// WithCLIPlatformContext sets the platform context in the server options configuration.
func WithCLIPlatformContext(context string) CLIOption {
	return cliOptionFunc(func(cfg *cliOptions) {
		cfg.PlatformContext = context
	})
}

// WithCLIConfigurationProvider sets the SpecConfigurationProvider for the server configuration and applies it as a CLIOption.
func WithCLIConfigurationProvider(settings BaseSpecConfigurationProvider) CLIOption {
	return cliOptionFunc(func(cfg *cliOptions) {
		cfg.ConfigurationProvider = settings
	})
}

// WithCLIRuntimeOverrides sets runtime configuration overrides for the CLI, modifying default behavior based on the provided settings.
func WithCLIRuntimeOverrides(overrides RuntimeConfigurationOverrides) CLIOption {
	overrides.Overridden = true
	return cliOptionFunc(func(cfg *cliOptions) {
		cfg.RuntimeConfigurationOverrides = overrides
	})
}

// CLIConfiguration represents a structure for application configuration settings, including app, GRPC, and HTTP details.
type CLIConfiguration struct {
	App      specv2pb.App      `yaml:"app,omitempty"`
	Platform specv2pb.Platform `yaml:"platform,omitempty"`
	Context  specv2pb.Context  `yaml:"context,omitempty"`
	Systems  []specv2pb.SpecSystem

	configuration *CLIConfiguration
	// err error
}

// ResolveConfiguration merges and resolves the environment and default configuration settings into a unified structure.
func (c *CLIConfiguration) ResolveConfiguration(opts ...ConfigurationProviderOption) (*Configurer, error) {
	var config CLIConfiguration

	opts = append(opts, WithConfigPath(ContextDirectory))
	configurer, err := NewConfigurer(opts...)
	if err != nil {
		return nil, err
	}

	Resolve(configurer, &config, c.GetDefaultConfiguration())

	config.configuration = &config
	c.configuration = &config

	return configurer, nil
}

// ValidateConfiguration checks if the configuration instance is valid and returns an error if validation fails.
func (c *CLIConfiguration) ValidateConfiguration() error {
	return nil
}

// GetDefaultConfiguration returns a default `CLIConfiguration` instance with preset values for App, Grpc, and Http fields.
func (c *CLIConfiguration) GetDefaultConfiguration() *CLIConfiguration {
	return &CLIConfiguration{
		App: specv2pb.App{
			Name:            "oeco",
			Version:         "0.0.0",
			EnvironmentName: "local-1",
			EnvironmentType: "local",
		},
		Platform: specv2pb.Platform{
			Endpoint:            "http://localhost:6577",
			Insecure:            true,
			DnsEndpoints:        []string{"45.63.49.173:4242"},
			DynamicConfigReload: false,
			Mesh: &specv2pb.Mesh{
				Enabled:     true,
				Endpoint:    "http://192.168.100.5:6477",
				Insecure:    true,
				DnsEndpoint: "192.168.100.1",
				Punchy:      true,
			},
		},
	}
}

// CreateConfiguration generates and returns a default or custom configuration for the Binding instance.
func (c *CLIConfiguration) CreateConfiguration() (*CLIConfiguration, error) {
	return nil, nil
}

// GetConfiguration retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (c *CLIConfiguration) GetConfiguration() *CLIConfiguration {
	return c.configuration
}

// GetConfigurationBytes retrieves the configuration of the binding instance. Returns the configuration as an interface{}.
func (c *CLIConfiguration) GetConfigurationBytes() ([]byte, error) {
	byteArray, err := json.Marshal(c.GetConfiguration())
	if err != nil {
		fmt.Println("SpecError:", err)
		return nil, err
	}
	return byteArray, nil
}

// WatchConfigurations observes changes in the binding's configuration and updates the internal state accordingly.
func (c *CLIConfiguration) WatchConfigurations(directories ...string) error {
	fmt.Println("Watch settings ecosystem internal directories:", directories)
	return nil
}
