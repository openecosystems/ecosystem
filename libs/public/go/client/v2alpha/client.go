package v2alpha

import (
	"context"
	"net/http"
	"net/url"

	"connectrpc.com/connect"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"github.com/slackhq/nebula/service"
)

// ResolvedClientConfiguration holds the resolved client settings used for configuring the client behavior and context.
var ResolvedClientConfiguration *specv2pb.SpecClientSettings

// ClientConfiguration represents the configuration for a client, including settings defined in SpecClientSettings.
type ClientConfiguration struct {
	SpecClientSettings specv2pb.SpecClientSettings `json:"spec_client_settings" yaml:"spec_client_settings"`
}

// SpecClient is a generic client for handling requests and responses with underlying connection and configuration support.
type SpecClient[Req, Res any] struct {
	UnderlyingClient *connect.Client[Req, Res]
	Config           *specv2pb.SpecClientSettings
	MeshSocket       *service.Service

	config *specClientConfig
	err    error
}

// NewSpecClient creates a new instance of SpecClient with the specified URL and optional configuration options.
// It initializes the client with provided settings and returns it. If an error occurs during initialization,
// the error is stored within the SpecClient instance.
func NewSpecClient[Req, Res any](url string, options ...SpecClientOption) *SpecClient[Req, Res] {
	client := &SpecClient[Req, Res]{}
	config, err := newSpecClientConfig(url, options)
	if err != nil {
		client.err = err
		return client
	}
	client.config = config

	client.UnderlyingClient = connect.NewClient[Req, Res](client.config.HTTPClient, url, config.UnderlyingClientOptions...)

	return client
}

// CallSpecUnary invokes a unary RPC method using the underlying client and returns the server's response or an error.
func (c *SpecClient[Req, Res]) CallSpecUnary(ctx context.Context, request *connect.Request[Req]) (*connect.Response[Res], error) {
	return c.UnderlyingClient.CallUnary(ctx, request)
}

// CallSpecClientStream creates and returns a new client-side streaming connection using the underlying client.
func (c *SpecClient[Req, Res]) CallSpecClientStream(ctx context.Context) *connect.ClientStreamForClient[Req, Res] {
	return c.UnderlyingClient.CallClientStream(ctx)
}

// CallSpecServerStream initiates a server-streaming RPC call using the underlying client with the given request.
func (c *SpecClient[Req, Res]) CallSpecServerStream(ctx context.Context, request *connect.Request[Req]) (*connect.ServerStreamForClient[Res], error) {
	return c.UnderlyingClient.CallServerStream(ctx, request)
}

// CallSpecBidiStream creates a bidirectional stream using the underlying client and the provided context.
func (c *SpecClient[Req, Res]) CallSpecBidiStream(ctx context.Context) *connect.BidiStreamForClient[Req, Res] {
	return c.UnderlyingClient.CallBidiStream(ctx)
}

// specClientConfig provides configuration settings for the spec client, including network details and client options.
type specClientConfig struct {
	URL                     *url.URL
	MeshVPN                 bool
	MeshURL                 string
	MeshSocket              *service.Service
	HTTPClient              *http.Client
	UnderlyingClientOptions []connect.ClientOption
	// Flags                   *RuntimeConfigurationOverrides
	// Filesystem              *io.FileSystem
}

// newSpecClientConfig initializes and returns a specClientConfig object and a potential error.
// It parses the provided rawURL and applies the given SpecClientOption configurations.
// Returns a connect.Error if URL parsing or validation fails.
func newSpecClientConfig(rawURL string, options []SpecClientOption) (*specClientConfig, *connect.Error) {
	uri, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, &connect.Error{}
	}

	config := specClientConfig{
		URL:        uri,
		MeshVPN:    false,
		HTTPClient: http.DefaultClient,
	}
	for _, opt := range options {
		opt.applyToClient(&config)
	}
	if err := config.validate(); err != nil {
		return nil, err
	}

	// cfg := ResolveConfiguration(config.Flags, config.Filesystem)
	// ResolvedConfiguration = cfg

	return &config, nil
}

// validate checks the integrity and correctness of the specClientConfig structure and returns a connect.Error if invalid.
func (c *specClientConfig) validate() *connect.Error {
	return nil
}

// SpecClientOption defines an interface for configuring a specClientConfig instance.
type SpecClientOption interface {
	applyToClient(*specClientConfig)
}

// WithSpecClientOptions combines multiple SpecClientOption items into a single SpecClientOption for batch application.
func WithSpecClientOptions(options ...SpecClientOption) SpecClientOption {
	return &specClientOptionsOption{options}
}

//// WithMesh configures clients to use an Ecosystem Mesh
//func WithMesh(url string) SpecClientOption {
//	return &meshOption{vpn: true, Url: url}
//}

// WithUnderlyingClientOptions applies the given client options to the underlying connect client configuration.
func WithUnderlyingClientOptions(options ...connect.ClientOption) connect.ClientOption {
	return connect.WithClientOptions(options...)
}

// specClientOptionsOption is a composite SpecClientOption that bundles multiple SpecClientOptions for configuration.
type specClientOptionsOption struct {
	options []SpecClientOption
}

// applyToClient applies each SpecClientOption from the options slice to the given specClientConfig instance.
func (o *specClientOptionsOption) applyToClient(config *specClientConfig) {
	for _, option := range o.options {
		option.applyToClient(config)
	}
}

//
//
//
//type meshOption struct {
//	vpn bool
//	Url string
//}
//
//func (o *meshOption) applyToClient(config *specClientConfig) {
//	config.MeshVPN = o.vpn
//	config.MeshURL = o.Url
//
//	config.HTTPClient = http.DefaultClient
//
//	go func() {
//		if nebulav1.IsBound {
//
//			configBytes, err := yaml.Marshal(nebulav1.ResolvedConfiguration.Nebula)
//			if err != nil {
//				fmt.Printf("Error resolving Nebula configuration: %v\n", err)
//				fmt.Printf(err.Error())
//			}
//
//			var cfg nebulaConfig.C
//			if err = cfg.LoadString(string(configBytes)); err != nil {
//				fmt.Println("ERROR loading config:", err)
//			}
//
//			svc, err := service.New(&cfg)
//			if err != nil {
//				fmt.Printf("Error creating service: %v\n", err)
//				fmt.Printf(err.Error())
//			}
//
//			config.MeshSocket = svc
//
//			config.HTTPClient = &http.Client{
//				Transport: &http.Transport{
//					DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
//						return svc.Dial("tcp", o.Url)
//					},
//				},
//			}
//
//		}
//	}()
//
//}
