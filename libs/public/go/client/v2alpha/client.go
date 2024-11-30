package v2alpha

import (
	"context"
	"net/http"
	"net/url"

	"connectrpc.com/connect"

	"github.com/slackhq/nebula/service"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

var ResolvedClientConfiguration *specv2pb.SpecClientSettings

type ClientConfiguration struct {
	SpecClientSettings specv2pb.SpecClientSettings `json:"spec_client_settings" yaml:"spec_client_settings"`
}

type SpecClient[Req, Res any] struct {
	UnderlyingClient *connect.Client[Req, Res]
	Config           *specv2pb.SpecClientSettings
	MeshSocket       *service.Service

	config *specClientConfig
	err    error
}

func NewSpecClient[Req, Res any](url string, options ...SpecClientOption) *SpecClient[Req, Res] {
	client := &SpecClient[Req, Res]{}
	config, err := newSpecClientConfig(url, options)
	if err != nil {
		client.err = err
		return client
	}
	client.config = config

	client.UnderlyingClient = connect.NewClient[Req, Res](client.config.HttpClient, url, config.UnderlyingClientOptions...)

	return client
}

// CallSpecUnary calls a request-response procedure.
func (c *SpecClient[Req, Res]) CallSpecUnary(ctx context.Context, request *connect.Request[Req]) (*connect.Response[Res], error) {
	return c.UnderlyingClient.CallUnary(ctx, request)
}

// CallSpecClientStream calls a client streaming procedure.
func (c *SpecClient[Req, Res]) CallSpecClientStream(ctx context.Context) *connect.ClientStreamForClient[Req, Res] {
	return c.UnderlyingClient.CallClientStream(ctx)
}

// CallSpecServerStream calls a server streaming procedure.
func (c *SpecClient[Req, Res]) CallSpecServerStream(ctx context.Context, request *connect.Request[Req]) (*connect.ServerStreamForClient[Res], error) {
	return c.UnderlyingClient.CallServerStream(ctx, request)
}

// CallSpecBidiStream calls a bidirectional streaming procedure.
func (c *SpecClient[Req, Res]) CallSpecBidiStream(ctx context.Context) *connect.BidiStreamForClient[Req, Res] {
	return c.UnderlyingClient.CallBidiStream(ctx)
}

type specClientConfig struct {
	URL                     *url.URL
	MeshVPN                 bool
	MeshUrl                 string
	MeshSocket              *service.Service
	HttpClient              *http.Client
	UnderlyingClientOptions []connect.ClientOption
	// Flags                   *RuntimeConfigurationOverrides
	// Filesystem              *io.FileSystem
}

func newSpecClientConfig(rawURL string, options []SpecClientOption) (*specClientConfig, *connect.Error) {
	uri, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, &connect.Error{}
	}

	config := specClientConfig{
		URL:        uri,
		MeshVPN:    false,
		HttpClient: http.DefaultClient,
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

func (c *specClientConfig) validate() *connect.Error {
	return nil
}

// A SpecClientOption configures a [SpecClient].
//
// In addition to any options grouped in the documentation below, remember that
// any [SpecOption] is also a valid SpecClientOption.
type SpecClientOption interface {
	applyToClient(*specClientConfig)
}

// WithSpecClientOptions composes multiple ClientOptions into one.
func WithSpecClientOptions(options ...SpecClientOption) SpecClientOption {
	return &specClientOptionsOption{options}
}

//// WithMesh configures clients to use an Ecosystem Mesh
//func WithMesh(url string) SpecClientOption {
//	return &meshOption{vpn: true, Url: url}
//}

// WithUnderlyingClientOptions composes multiple ClientOptions into one.
func WithUnderlyingClientOptions(options ...connect.ClientOption) connect.ClientOption {
	return connect.WithClientOptions(options...)
}

type specClientOptionsOption struct {
	options []SpecClientOption
}

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
//	config.MeshUrl = o.Url
//
//	config.HttpClient = http.DefaultClient
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
//			config.HttpClient = &http.Client{
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
