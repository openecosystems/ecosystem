package serverv2alphalib

import (
	"net/http"
	"net/url"

	"connectrpc.com/connect"
)

// A ServerOption configures a [Server].
type ServerOption interface {
	apply(*serverOptions)
}

// WithHttpServer provides a low level [http.Server]
func WithHttpServer(httpServer *http.ServeMux) ServerOption {
	return &httpServerOption{
		HttpServer: httpServer,
	}
}

type serverOptions struct {
	URL        *url.URL
	MeshVPN    bool
	HttpServer *http.ServeMux
	// UnderlyingHandlerOptions []connect.HandlerOption
}

func newServerOptions(rawURL string, options []ServerOption) (*serverOptions, *connect.Error) {
	uri, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, &connect.Error{}
	}

	config := serverOptions{
		URL:        uri,
		MeshVPN:    false,
		HttpServer: http.NewServeMux(),
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

type httpServerOption struct {
	HttpServer *http.ServeMux
}

func (o *httpServerOption) apply(config *serverOptions) {
	o.apply(config)
}
