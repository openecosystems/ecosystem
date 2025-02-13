package cliv2alphalib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"connectrpc.com/connect"
)

// CLI represents the core structure for managing CLI operations, bindings, and their configuration.
type CLI struct {
	Bindings              *sdkv2alphalib.Bindings
	Bounds                []sdkv2alphalib.Binding
	ConfigurationProvider sdkv2alphalib.BaseSpecConfigurationProvider

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

	configurer, err := t.ResolveConfiguration()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = t.ValidateConfiguration()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	bindings := sdkv2alphalib.RegisterBindings(ctx, options.Bounds, sdkv2alphalib.WithConfigurer(configurer))
	cli.Bindings = bindings

	return &cli
}

// GracefulShutdown gracefully shuts down the CLI by cleaning up resources and closing bindings within a timeout context.
func (cli *CLI) GracefulShutdown() {
	_, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	sdkv2alphalib.ShutdownBindings(cli.Bindings)
}

// GetConfiguration retrieves and returns the configuration from the CLI's ConfigurationProvider if implemented.
func (cli *CLI) GetConfiguration() *Configuration {
	t := cli.ConfigurationProvider

	bytes, err := t.GetConfigurationBytes()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	c := Configuration{}
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

	Bounds                        []sdkv2alphalib.Binding
	PlatformContext               string
	RuntimeConfigurationOverrides *sdkv2alphalib.RuntimeConfigurationOverrides
	ConfigurationProvider         sdkv2alphalib.BaseSpecConfigurationProvider
}

type optionFunc func(*cliOptions)

func (f optionFunc) apply(cfg *cliOptions) { f(cfg) }

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

// WithOptions composes multiple Options into one.
func WithOptions(opts ...CLIOption) CLIOption {
	return optionFunc(func(cfg *cliOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithBounds configures the server with the specified bounds, overriding the default bindings list in server options.
func WithBounds(bounds []sdkv2alphalib.Binding) CLIOption {
	return optionFunc(func(cfg *cliOptions) {
		cfg.Bounds = bounds
	})
}

// WithPlatformContext sets the platform context in the server options configuration.
func WithPlatformContext(context string) CLIOption {
	return optionFunc(func(cfg *cliOptions) {
		cfg.PlatformContext = context
	})
}

// WithConfigurationProvider sets the SpecConfigurationProvider for the server configuration and applies it as a CLIOption.
func WithConfigurationProvider(settings sdkv2alphalib.BaseSpecConfigurationProvider) CLIOption {
	return optionFunc(func(cfg *cliOptions) {
		cfg.ConfigurationProvider = settings
	})
}

// WithRuntimeOverrides sets runtime configuration overrides for the CLI, modifying default behavior based on the provided settings.
func WithRuntimeOverrides(overrides *sdkv2alphalib.RuntimeConfigurationOverrides) CLIOption {
	return optionFunc(func(cfg *cliOptions) {
		cfg.RuntimeConfigurationOverrides = overrides
	})
}
