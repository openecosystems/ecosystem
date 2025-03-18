package sdkv2alphalib

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"connectrpc.com/connect"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Infrastructure represents the core structure for managing configuration, bindings, and their associated instances.
type Infrastructure struct {
	Bindings              *Bindings
	Bounds                []Binding
	ConfigurationProvider *BaseSpecConfigurationProvider

	options *infrastructureOptions
}

// NewInfrastructure initializes a new Infrastructure instance with specified bindings and resolved configuration.
func NewInfrastructure(ctx context.Context, opts ...InfrastructureOption) *Infrastructure {
	options, err := newInfrastructureOptions(opts)
	if err != nil {
		fmt.Println("new connector options error: ")
		fmt.Println(err)
	}

	infrastructure := &Infrastructure{
		Bounds:  options.Bounds,
		options: options,
	}

	provider := options.ConfigurationProvider
	if provider == nil {
		panic("configuration provider is nil. Please provide a configuration provider to the server.")
	}

	infrastructure.ConfigurationProvider = &provider
	t := options.ConfigurationProvider

	configurer, cerr := t.ResolveConfiguration()
	if cerr != nil {
		return nil
	}
	cerr = t.ValidateConfiguration()
	if cerr != nil {
		fmt.Println(cerr)
		panic(cerr)
	}

	bindings := RegisterBindings(ctx, options.Bounds, WithConfigurer(configurer))
	infrastructure.Bindings = bindings

	return infrastructure
}

// Run executes a Pulumi program using the provided RunFunc and optional runtime configuration options.
func (infrastructure *Infrastructure) Run(runFunc pulumi.RunFunc, opts ...pulumi.RunOption) {
	pulumi.Run(runFunc, opts...)
}

// ShortenString truncates the input string `s` to the specified `limit` while ensuring it remains a valid UTF-8 string.
func ShortenString(s string, limit int) string {
	if len(s) < limit {
		return s
	}

	if utf8.ValidString(s[:limit]) {
		return s[:limit]
	}
	return s[:limit+1]
}

// WriteIndentedMultilineText takes a multiline string and returns it with each line prefixed by an 8-space indentation.
func WriteIndentedMultilineText(text string) string {
	indent := "        "
	lines := strings.Split(text, "\n")

	var builder strings.Builder

	for _, line := range lines {
		_, err := builder.WriteString(indent + line + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	return builder.String()
}

// InfrastructureOption defines an interface for applying custom configuration to a infrastructureOptions object.
type InfrastructureOption interface {
	apply(*infrastructureOptions)
}

// infrastructureOptions defines the configuration options for a infrastructure, including supported protocols and codecs.
type infrastructureOptions struct {
	Bounds                []Binding
	ConfigurationProvider BaseSpecConfigurationProvider

	protocols map[typev2pb.Protocol]struct{}
	// codecNames     map[string]struct{}
	// preferredCodec string
}

// infrastructureOptionFunc is a function type that modifies the settings of a infrastructureOptions instance.
type infrastructureOptionFunc func(*infrastructureOptions)

// apply applies the infrastructureOptionFunc to the given infrastructureOptions.
func (f infrastructureOptionFunc) apply(opts *infrastructureOptions) {
	f(opts)
}

// newInfrastructureOptions creates and configures a new infrastructureOptions instance using the provided InfrastructureOption slice.
// Returns the configured infrastructureOptions and an error if validation fails.
func newInfrastructureOptions(options []InfrastructureOption) (*infrastructureOptions, *connect.Error) {
	config := infrastructureOptions{
		protocols: nil,
	}

	for _, opt := range options {
		opt.apply(&config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

// validate checks the integrity and consistency of the infrastructureOptions fields.
// Returns a *connect.Error if validation fails or nil if successful.
func (c *infrastructureOptions) validate() *connect.Error {
	return nil
}

// WithInfrastructureOptions composes multiple Options into one.
func WithInfrastructureOptions(opts ...InfrastructureOption) InfrastructureOption {
	return infrastructureOptionFunc(func(cfg *infrastructureOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithInfrastructureBounds configures the infrastructure with the specified bounds, overriding the default bindings list in server options.
func WithInfrastructureBounds(bounds []Binding) InfrastructureOption {
	return infrastructureOptionFunc(func(cfg *infrastructureOptions) {
		cfg.Bounds = bounds
	})
}

// WithInfrastructureConfigurationProvider sets the SpecConfigurationProvider for the server configuration and applies it as a ServerOption.
func WithInfrastructureConfigurationProvider(settings BaseSpecConfigurationProvider) InfrastructureOption {
	return infrastructureOptionFunc(func(cfg *infrastructureOptions) {
		cfg.ConfigurationProvider = settings
	})
}
