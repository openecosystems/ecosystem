package sdkv2betalib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/vanguard"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

// serverQuit is a channel used to handle OS signals for graceful shutdown of the server.
var serverQuit = make(chan os.Signal, 1)

// Server represents a configurable HTTP/2 server with bindings and service handlers.
type Server struct {
	Bindings                *Bindings
	PublicConnectHTTPServer *http2.Server
	MeshConnectHTTPServer   *http2.Server
	PublicHTTPServerHandler *http.ServeMux
	MeshHTTPServerHandler   *http.ServeMux
	Bounds                  []Binding
	ServicePath             string
	PublicServiceHandler    *vanguard.Transcoder
	MeshServiceHandler      *vanguard.Transcoder
	RawServiceHandler       *http.ServeMux
	ConfigurationProvider   *BaseSpecConfigurationProvider
	NetListener             *net.Listener

	options *serverOptions
	// err     error
}

// NewServer creates and initializes a new multiplexed server with bindings, services, and server options.
func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	options, _ := newServerOptions(opts)

	server := &Server{
		Bounds:      options.Bounds,
		ServicePath: "/",

		options: options,
	}

	provider := options.ConfigurationProvider
	if provider == nil {
		fmt.Println("configuration provider is nil. Please provide a configuration provider to the server.")
		os.Exit(1)
	}

	server.ConfigurationProvider = &provider
	t := options.ConfigurationProvider

	configurer, err := t.ResolveConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = t.ValidateConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bindings := RegisterBindings(ctx, options.Bounds, WithConfigurer(configurer))
	server.Bindings = bindings

	if options.PublicServices != nil {
		publicHTTPServer := &http2.Server{
			IdleTimeout:      15 * time.Second,
			WriteByteTimeout: 10 * time.Second,
			ReadIdleTimeout:  5 * time.Second,
		}

		publicTranscoder, err := vanguard.NewTranscoder(options.PublicServices)
		if err != nil {
			fmt.Println(err)
		}

		server.PublicConnectHTTPServer = publicHTTPServer
		server.PublicServiceHandler = publicTranscoder
	}

	if options.MeshServices != nil {
		meshHTTPServer := &http2.Server{
			IdleTimeout:      15 * time.Second,
			WriteByteTimeout: 10 * time.Second,
			ReadIdleTimeout:  5 * time.Second,
		}

		meshTranscoder, err2 := vanguard.NewTranscoder(options.MeshServices)
		if err2 != nil {
			fmt.Println(err2)
		}

		server.MeshConnectHTTPServer = meshHTTPServer
		server.MeshServiceHandler = meshTranscoder
	}

	if options.RawServerOptions != nil {
		rawOptions := options.RawServerOptions
		server.RawServiceHandler = rawOptions.Handler

		httpServer := &http2.Server{
			IdleTimeout:      15 * time.Second,
			WriteByteTimeout: 10 * time.Second,
			ReadIdleTimeout:  5 * time.Second,
			// MaxConcurrentStreams:         0,
			// PermitProhibitedCipherSuites: false,
			// MaxUploadBufferPerConnection: 0,
			// MaxUploadBufferPerStream:     0,
		}

		server.ServicePath = rawOptions.Path
		server.PublicConnectHTTPServer = httpServer
	}

	if options.NetListener != nil {
		server.NetListener = options.NetListener
	}

	return server
}

// ListenAndServe starts the server and listens for incoming HTTP requests on the configured address and port.
func (server *Server) ListenAndServe() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server.ListenAndServeWithCtx(ctx)
}

// ListenAndServeWithCtx starts the server with a provided context, handling HTTP and specialized bindings listeners.
// It supports graceful shutdown management upon receiving termination signals.
func (server *Server) ListenAndServeWithCtx(_ context.Context) {
	httpServerErr := server.ListenAndServeMultiplexedHTTP()

	var specListenableErr chan SpecListenableErr
	if server.Bindings.RegisteredListenableChannels != nil {
		go func() {
			specListenableErr = server.ListenAndServeSpecListenable()
		}()
	}

	/*
	 * Graceful Shutdown Management
	 */
	signal.Notify(serverQuit, syscall.SIGTERM)
	signal.Notify(serverQuit, os.Interrupt)
	select {
	case err := <-specListenableErr:
		if err.Error != nil {
			fmt.Println(ErrServerInternal.WithInternalErrorDetail(err.Error, errors.New("received a specListenableErr")).Error())
		}
	case err := <-httpServerErr:
		fmt.Println(ErrServerInternal.WithInternalErrorDetail(err, errors.New("received an httpServerError")).Error())
	case <-serverQuit:
		server.Shutdown()
	}
}

// ListenAndServeWithProvidedSocket starts serving HTTP requests using the provided net.Listener and returns an error channel.
func (server *Server) ListenAndServeWithProvidedSocket(ln *net.Listener) {
	server.NetListener = ln
	server.ListenAndServe()
}

// ListenAndServeMultiplexedHTTP starts an HTTP server supporting HTTP/2 without TLS over a multiplexed handler function.
// It returns a channel for listening to server errors during execution.
func (server *Server) ListenAndServeMultiplexedHTTP() (httpServerErr chan error) {
	var ln *net.Listener
	if server.NetListener != nil {
		ln = server.NetListener
	}

	return server.listenAndServe(ln)
}

// listenAndServe starts an HTTP/2-compatible server, optionally on a given listener, and returns a channel for errors.
// It configures the server with service handlers and supports HTTP/2 without TLS using h2c.
func (server *Server) listenAndServe(ln *net.Listener) (httpServerErr chan error) {
	cp := *server.ConfigurationProvider

	bytes, err := cp.GetConfigurationBytes()
	if err != nil {
		fmt.Println("GetConfigurationBytes error:")
		return nil
	}

	var settings specv2pb.SpecSettings
	err = json.Unmarshal(bytes, &settings)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	publicEndpoint := settings.Platform.GetEndpoint()
	meshEndpoint := settings.Platform.Mesh.GetEndpoint()

	publicMux := http.NewServeMux()
	if server.PublicHTTPServerHandler != nil {
		publicMux = server.PublicHTTPServerHandler
	}

	server.PublicHTTPServerHandler = publicMux
	if server.RawServiceHandler != nil {
		publicMux.Handle(server.ServicePath, server.RawServiceHandler)
	} else {
		publicMux.Handle("/", server.PublicServiceHandler)
	}

	meshMux := http.NewServeMux()
	if server.MeshHTTPServerHandler != nil {
		meshMux = server.MeshHTTPServerHandler
	}

	server.MeshHTTPServerHandler = meshMux
	meshMux.Handle("/", server.MeshServiceHandler)

	publicHTTPServer := &http.Server{
		Addr: publicEndpoint,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler:      h2c.NewHandler(edgeRouter(publicMux), server.PublicConnectHTTPServer),
		ReadTimeout:  5 * time.Second,  // Time allowed to read the request
		WriteTimeout: 10 * time.Second, // Time allowed to write the response
		IdleTimeout:  15 * time.Second, // Time for keep-alive connections
	}

	meshHTTPServer := &http.Server{
		Addr: meshEndpoint,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler:      h2c.NewHandler(edgeRouter(meshMux), server.MeshConnectHTTPServer),
		ReadTimeout:  5 * time.Second,  // Time allowed to read the request
		WriteTimeout: 10 * time.Second, // Time allowed to write the response
		IdleTimeout:  15 * time.Second, // Time for keep-alive connections
	}

	_httpServerErr := make(chan error)

	if server.PublicConnectHTTPServer != nil {
		go func() {
			_httpServerErr <- publicHTTPServer.ListenAndServe()
		}()
		fmt.Println("Public HTTP1.1/HTTP2.0/gRPC/gRPC-Web/Connect listening on " + settings.Platform.Endpoint)
	}

	if server.MeshConnectHTTPServer != nil {
		if settings.Platform.Mesh.Enabled {
			go func() {
				if ln != nil {
					_httpServerErr <- meshHTTPServer.Serve(*ln)
				} else {
					_httpServerErr <- meshHTTPServer.ListenAndServe()
				}
			}()
			fmt.Println("Mesh HTTP1.1/HTTP2.0/gRPC/gRPC-Web/Connect listening on " + settings.Platform.Mesh.Endpoint)
		}
	}

	return _httpServerErr
}

// ListenAndServeSpecListenable starts listening on all registered SpecListenable channels and returns a channel for errors.
func (server *Server) ListenAndServeSpecListenable() chan SpecListenableErr {
	listeners := server.Bindings.RegisteredListenableChannels
	listenerErr := make(chan SpecListenableErr, len(listeners))

	for key, listener := range listeners {
		ctx := context.Background()
		go listener.Listen(ctx, listenerErr)

		fmt.Println("Registered Embedded Connector: " + key)
	}
	return listenerErr
}

// Shutdown gracefully stopping the connector
func (server *Server) Shutdown() {
	fmt.Printf("Stopping server gracefully. Draining connections for up to %v seconds", 30)
	fmt.Println()

	_, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	ShutdownBindings(server.Bindings)
}

// A ServerOption configures a [Server].
type ServerOption interface {
	apply(*serverOptions)
}

// RawServerOptions defines options for setting up a raw HTTP server, including the server's path and its handler.
type RawServerOptions struct {
	Path    string
	Handler *http.ServeMux
}

type serverOptions struct {
	URL                   *url.URL
	MeshVPN               bool
	HTTPServer            *http.ServeMux
	PublicServices        []*vanguard.Service
	MeshServices          []*vanguard.Service
	RawServerOptions      *RawServerOptions
	Bounds                []Binding
	PlatformContext       string
	ConfigPath            string
	ConfigurationProvider BaseSpecConfigurationProvider
	NetListener           *net.Listener
}

type serverOptionFunc func(*serverOptions)

func (f serverOptionFunc) apply(cfg *serverOptions) { f(cfg) }

//nolint:unparam
func newServerOptions(options []ServerOption) (*serverOptions, *connect.Error) {
	// Defaults
	config := serverOptions{
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

// WithServerOptions composes multiple Options into one.
func WithServerOptions(opts ...ServerOption) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithPublicServices sets the public services for the server configuration and returns a ServerOption.
func WithPublicServices(services []*vanguard.Service) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.PublicServices = services
	})
}

// WithMeshServices sets the mesh services for the server configuration.
func WithMeshServices(services []*vanguard.Service) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.MeshServices = services
	})
}

// WithBounds configures the server with the specified bounds, overriding the default bindings list in server options.
func WithBounds(bounds []Binding) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.Bounds = bounds
	})
}

// WithServerPlatformContext sets the platform context in the server options configuration.
func WithServerPlatformContext(context string) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.PlatformContext = context
	})
}

// WithServerConfigPath sets the configuration file path in the server options.
func WithServerConfigPath(path string) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.ConfigPath = path
	})
}

// WithConfigurationProvider sets the SpecConfigurationProvider for the server configuration and applies it as a ServerOption.
func WithConfigurationProvider(settings BaseSpecConfigurationProvider) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.ConfigurationProvider = settings
	})
}

// WithRawServer creates a ServerOption to configure raw server options with the provided RawServerOptions parameter.
func WithRawServer(options *RawServerOptions) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		uri, _ := url.ParseRequestURI(options.Path)
		cfg.URL = uri
		cfg.RawServerOptions = options
	})
}

// WithNetListener with a mesh or other net.Listener
func WithNetListener(ln *net.Listener) ServerOption {
	return serverOptionFunc(func(cfg *serverOptions) {
		cfg.NetListener = ln
	})
}
