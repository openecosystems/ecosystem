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

// WithHTTPServer provides a low level [http.Server]
func WithHTTPServer(httpServer *http.ServeMux) ServerOption {
	return &httpServerOption{
		HTTPServer: httpServer,
	}
}

type serverOptions struct {
	URL        *url.URL
	MeshVPN    bool
	HTTPServer *http.ServeMux
	// UnderlyingHandlerOptions []connect.HandlerOption
}

//nolint:unparam
func newServerOptions(rawURL string, options []ServerOption) (*serverOptions, *connect.Error) {
	uri, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, &connect.Error{}
	}

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

type httpServerOption struct {
	HTTPServer *http.ServeMux
}

func (o *httpServerOption) apply(config *serverOptions) {
	// Delegate to a different method or perform specific logic here
	o.applyImplementation(config)
}

func (o *httpServerOption) applyImplementation(_ *serverOptions) {
	// Actual logic for modifying the serverOptions
}
