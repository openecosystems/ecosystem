package serverv2alphalib

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/vanguard"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	nebulav1 "libs/partner/go/nebula/v1"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// quit is a channel used to handle OS signals for graceful shutdown of the server.
var quit = make(chan os.Signal, 1)

// Server represents a configurable HTTP/2 server with bindings and service handlers.
type Server struct {
	Bindings                *sdkv2alphalib.Bindings
	PublicConnectHTTPServer *http2.Server
	MeshConnectHTTPServer   *http2.Server
	PublicHTTPServerHandler *http.ServeMux
	MeshHTTPServerHandler   *http.ServeMux
	Bounds                  []sdkv2alphalib.Binding
	ServicePath             string
	PublicServiceHandler    *vanguard.Transcoder
	MeshServiceHandler      *vanguard.Transcoder
	RawServiceHandler       *http.Handler
	ConfigurationProvider   *sdkv2alphalib.SpecConfigurationProvider

	options *serverOptions
	err     error
}

// NewServer creates and initializes a new multiplexed server with bindings, services, and server options.
func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	options, _ := newServerOptions("", opts)

	server := &Server{
		Bounds:      options.Bounds,
		ServicePath: "/",

		options: options,
	}

	if options.ConfigurationProvider != nil {
		server.ConfigurationProvider = options.ConfigurationProvider

		t := *options.ConfigurationProvider
		t.ResolveConfiguration()
		err := t.ValidateConfiguration()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	bindings := sdkv2alphalib.RegisterBindings(ctx, options.Bounds)
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

	return server
}

// NewMultiplexedServer creates and initializes a new multiplexed server with bindings, services, and server options.
func NewMultiplexedServer(ctx context.Context, bounds []sdkv2alphalib.Binding, meshServices []*vanguard.Service, publicServices []*vanguard.Service, opts ...ServerOption) *Server {
	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	publicHTTPServer := &http2.Server{
		IdleTimeout:      15 * time.Second,
		WriteByteTimeout: 10 * time.Second,
		ReadIdleTimeout:  5 * time.Second,
	}

	meshHTTPServer := &http2.Server{
		IdleTimeout:      15 * time.Second,
		WriteByteTimeout: 10 * time.Second,
		ReadIdleTimeout:  5 * time.Second,
	}

	options, _ := newServerOptions("", opts)

	publicTranscoder, err2 := vanguard.NewTranscoder(publicServices)
	if err2 != nil {
		fmt.Println(err2)
	}

	meshTranscoder, err2 := vanguard.NewTranscoder(meshServices)
	if err2 != nil {
		fmt.Println(err2)
	}

	return &Server{
		Bindings:                bindings,
		PublicConnectHTTPServer: publicHTTPServer,
		MeshConnectHTTPServer:   meshHTTPServer,
		Bounds:                  bounds,
		ServicePath:             "/",
		PublicServiceHandler:    publicTranscoder,
		MeshServiceHandler:      meshTranscoder,

		options: options,
		err:     err2,
	}
}

// NewRawServer initializes and returns a new Server instance with provided bindings, path, handler, and options.
// It resolves and validates configuration, registers bindings, and sets up the HTTP/2 server.
func NewRawServer(ctx context.Context, bounds []sdkv2alphalib.Binding, path string, handler *http.Handler, opts ...ServerOption) *Server {
	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	httpServer := &http2.Server{
		IdleTimeout:      15 * time.Second,
		WriteByteTimeout: 10 * time.Second,
		ReadIdleTimeout:  5 * time.Second,
		// MaxConcurrentStreams:         0,
		// PermitProhibitedCipherSuites: false,
		// MaxUploadBufferPerConnection: 0,
		// MaxUploadBufferPerStream:     0,
	}

	options, _ := newServerOptions(path, opts)

	return &Server{
		Bindings:              bindings,
		MeshConnectHTTPServer: httpServer,
		Bounds:                bounds,
		ServicePath:           path,
		RawServiceHandler:     handler,

		options: options,
	}
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

	var specListenableErr chan sdkv2alphalib.SpecListenableErr
	if server.Bindings.RegisteredListenableChannels != nil {
		go func() {
			specListenableErr = server.ListenAndServeSpecListenable()
		}()
	}

	/*
	 * Graceful Shutdown Management
	 */
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	select {
	case err := <-specListenableErr:
		if err.Error != nil {
			fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err.Error, errors.New("received a specListenableErr")).Error())
		}
	case err := <-httpServerErr:
		fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err, errors.New("received an httpServerError")).Error())
	case <-quit:
		fmt.Printf("Stopping edged gracefully. Draining connections for up to %v seconds", 30)
		fmt.Println()

		_, cancel := context.WithTimeout(context.Background(), 30)
		defer cancel()

		sdkv2alphalib.ShutdownBindings(server.Bindings)
	}
}

// ListenAndServeWithProvidedSocket starts serving HTTP requests using the provided net.Listener and returns an error channel.
func (server *Server) ListenAndServeWithProvidedSocket(ln net.Listener) (httpServerErr chan error) {
	return server.listenAndServe(ln)
}

// ListenAndServeMultiplexedHTTP starts an HTTP server supporting HTTP/2 without TLS over a multiplexed handler function.
// It returns a channel for listening to server errors during execution.
func (server *Server) ListenAndServeMultiplexedHTTP() (httpServerErr chan error) {
	return server.listenAndServe(nil)
}

// listenAndServe starts an HTTP/2-compatible server, optionally on a given listener, and returns a channel for errors.
// It configures the server with service handlers and supports HTTP/2 without TLS using h2c.
func (server *Server) listenAndServe(ln net.Listener) (httpServerErr chan error) {
	// u := *server.ConfigurationProvider

	// settings := u.GetDefaultConfiguration().(specv2pb.SpecSettings) //nolint:govet,copylocks

	// publicEndpoint := settings.Platform.GetEndpoint()
	// meshEndpoint := settings.Platform.Mesh.GetEndpoint()

	publicEndpoint := "0.0.0.0:8080"
	meshEndpoint := "0.0.0.0:8081"

	publicMux := http.NewServeMux()
	if server.PublicHTTPServerHandler != nil {
		publicMux = server.PublicHTTPServerHandler
	}
	server.PublicHTTPServerHandler = publicMux
	publicMux.Handle("/", server.PublicServiceHandler)

	meshMux := http.NewServeMux()
	if server.MeshHTTPServerHandler != nil {
		meshMux = server.MeshHTTPServerHandler
	}
	server.MeshHTTPServerHandler = meshMux
	if server.RawServiceHandler != nil {
		meshMux.Handle(server.ServicePath, *server.RawServiceHandler)
	} else {
		meshMux.Handle("/", server.MeshServiceHandler)
	}

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

	go func() {
		_httpServerErr <- publicHTTPServer.ListenAndServe()
	}()
	// fmt.Println("Public HTTP1.1/HTTP2.0/gRPC/gRPC-Web/Connect listening on " + settings.Platform.Endpoint)
	fmt.Println("Public HTTP1.1/HTTP2.0/gRPC/gRPC-Web/Connect listening on " + publicEndpoint)

	// if settings.Platform.Mesh.Enabled {
	if false {
		_ln, err3 := nebulav1.Bound.GetMeshListener(meshEndpoint)
		if err3 != nil {
			fmt.Println("get socket error: ", err3)
			return _httpServerErr
		}
		ln = *_ln
		fmt.Println("Mesh traffic routing enabled")
	}

	go func() {
		if ln != nil {
			_httpServerErr <- meshHTTPServer.Serve(ln)
		} else {
			_httpServerErr <- meshHTTPServer.ListenAndServe()
		}
	}()
	// fmt.Println("Mesh HTTP1.1/HTTP2.0/gRPC/gRPC-Web/Connect listening on " + settings.Platform.Mesh.Endpoint)
	fmt.Println("Mesh HTTP1.1/HTTP2.0/gRPC/gRPC-Web/Connect listening on " + meshEndpoint)

	return _httpServerErr
}

// ListenAndServeSpecListenable starts listening on all registered SpecListenable channels and returns a channel for errors.
func (server *Server) ListenAndServeSpecListenable() chan sdkv2alphalib.SpecListenableErr {
	listeners := server.Bindings.RegisteredListenableChannels
	listenerErr := make(chan sdkv2alphalib.SpecListenableErr, len(listeners))

	for key, listener := range listeners {
		ctx := context.Background()
		go listener.Listen(ctx, listenerErr)

		fmt.Println("Registered Embedded Connector: " + key)
	}
	return listenerErr
}
