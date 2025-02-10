package serverv2alphalib

import (
	"net/http"
	"net/url"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"connectrpc.com/connect"
	"connectrpc.com/vanguard"
)

// A ServerOption configures a [Server].
type ServerOption interface {
	apply(*serverOptions)
}

type serverOptions struct {
	URL                   *url.URL
	MeshVPN               bool
	HTTPServer            *http.ServeMux
	PublicServices        []*vanguard.Service
	MeshServices          []*vanguard.Service
	Bounds                []sdkv2alphalib.Binding
	PlatformContext       string
	ConfigPath            string
	ConfigurationProvider *sdkv2alphalib.SpecConfigurationProvider
}

type optionFunc func(*serverOptions)

func (f optionFunc) apply(cfg *serverOptions) { f(cfg) }

//nolint:unparam
func newServerOptions(rawURL string, options []ServerOption) (*serverOptions, *connect.Error) {
	uri, _ := url.ParseRequestURI(rawURL)

	config := serverOptions{
		URL:        uri,
		MeshVPN:    false,
		HTTPServer: http.NewServeMux(),
	}

	for _, opt := range options {
		opt.apply(&config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *serverOptions) validate() *connect.Error {
	return nil
}

// WithOptions composes multiple Options into one.
func WithOptions(opts ...ServerOption) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithPublicServices sets the public services for the server configuration and returns a ServerOption.
func WithPublicServices(services []*vanguard.Service) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		cfg.PublicServices = services
	})
}

// WithMeshServices sets the mesh services for the server configuration.
func WithMeshServices(services []*vanguard.Service) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		cfg.MeshServices = services
	})
}

// WithBounds configures the server with the specified bounds, overriding the default bindings list in server options.
func WithBounds(bounds []sdkv2alphalib.Binding) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		cfg.Bounds = bounds
	})
}

// WithPlatformContext sets the platform context in the server options configuration.
func WithPlatformContext(context string) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		cfg.PlatformContext = context
	})
}

// WithConfigPath sets the configuration file path in the server options.
func WithConfigPath(path string) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		cfg.ConfigPath = path
	})
}

// WithConfigurationProvider sets the SpecConfigurationProvider for the server configuration and applies it as a ServerOption.
func WithConfigurationProvider(settings sdkv2alphalib.SpecConfigurationProvider) ServerOption {
	return optionFunc(func(cfg *serverOptions) {
		cfg.ConfigurationProvider = &settings
	})
}
