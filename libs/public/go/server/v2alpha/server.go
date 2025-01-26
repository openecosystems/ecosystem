package serverv2alphalib

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"connectrpc.com/vanguard"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// quit is a channel used to handle OS signals for graceful shutdown of the server.
var quit = make(chan os.Signal, 1)

// Server represents a configurable HTTP/2 server with bindings and service handlers.
type Server struct {
	Bindings          *sdkv2alphalib.Bindings
	ConnectHTTPServer *http2.Server
	HTTPServerHandler *http.ServeMux
	Bounds            []sdkv2alphalib.Binding
	ServicePath       string
	ServiceHandler    *vanguard.Transcoder
	RawServiceHandler *http.Handler

	options *serverOptions
	err     error
}

// NewServer initializes and returns a new Server instance with specified context, bindings, path, handler, and options.
func NewServer(ctx context.Context, bounds []sdkv2alphalib.Binding, path string, handler *http.Handler, opts ...ServerOption) *Server {
	c := Configuration{}
	c.ResolveConfiguration()
	err := c.ValidateConfiguration()
	if err != nil {
		fmt.Println("validate server configuration error: ", err)
		panic(err)
	}

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

	s := vanguard.NewService(path, *handler)
	transcoder, err2 := vanguard.NewTranscoder([]*vanguard.Service{s})
	if err2 != nil {
		fmt.Println(err2)
	}

	return &Server{
		Bindings:          bindings,
		ConnectHTTPServer: httpServer,
		Bounds:            bounds,
		ServicePath:       path,
		ServiceHandler:    transcoder,

		options: options,
		err:     errors.Join(err, err2),
	}
}

// NewMultiplexedServer creates and initializes a new multiplexed server with bindings, services, and server options.
func NewMultiplexedServer(ctx context.Context, bounds []sdkv2alphalib.Binding, services []*vanguard.Service, opts ...ServerOption) *Server {
	c := Configuration{}
	c.ResolveConfiguration()
	err := c.ValidateConfiguration()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	httpServer := &http2.Server{
		IdleTimeout:      15 * time.Second,
		WriteByteTimeout: 10 * time.Second,
		ReadIdleTimeout:  5 * time.Second,
	}

	options, _ := newServerOptions("", opts)

	transcoder, err2 := vanguard.NewTranscoder(services)
	if err2 != nil {
		fmt.Println(err2)
	}

	return &Server{
		Bindings:          bindings,
		ConnectHTTPServer: httpServer,
		Bounds:            bounds,
		ServicePath:       "/",
		ServiceHandler:    transcoder,

		options: options,
		err:     errors.Join(err, err2),
	}
}

// NewRawServer initializes and returns a new Server instance with provided bindings, path, handler, and options.
// It resolves and validates configuration, registers bindings, and sets up the HTTP/2 server.
func NewRawServer(ctx context.Context, bounds []sdkv2alphalib.Binding, path string, handler *http.Handler, opts ...ServerOption) *Server {
	c := Configuration{}
	c.ResolveConfiguration()
	err := c.ValidateConfiguration()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

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
		Bindings:          bindings,
		ConnectHTTPServer: httpServer,
		Bounds:            bounds,
		ServicePath:       path,
		RawServiceHandler: handler,

		options: options,
		err:     errors.Join(err),
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

	fmt.Println("Server started successfully. HTTP listening on " + ResolvedConfiguration.HTTP.Port)

	/*
	 * Graceful Shutdown Management
	 */
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	select {
	case err := <-specListenableErr:
		if err.Error != nil {
			fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err.Error).Error())
		}
	case err := <-httpServerErr:
		fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err).Error())
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
	httpPort, _ := strconv.Atoi(ResolvedConfiguration.HTTP.Port)
	mux := http.NewServeMux()
	if server.HTTPServerHandler != nil {
		mux = server.HTTPServerHandler
	}

	server.HTTPServerHandler = mux

	if server.RawServiceHandler != nil {
		mux.Handle(server.ServicePath, *server.RawServiceHandler)
	} else {
		mux.Handle("/", server.ServiceHandler)
	}

	httpServer := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", httpPort),
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler:      h2c.NewHandler(edgeRouter(mux), server.ConnectHTTPServer),
		ReadTimeout:  5 * time.Second,  // Time allowed to read the request
		WriteTimeout: 10 * time.Second, // Time allowed to write the response
		IdleTimeout:  15 * time.Second, // Time for keep-alive connections
	}

	_httpServerErr := make(chan error)
	go func() {
		if ln != nil {
			_httpServerErr <- httpServer.Serve(ln)
		} else {
			_httpServerErr <- httpServer.ListenAndServe()
		}
	}()

	return _httpServerErr
}

// ListenAndServeSpecListenable starts listening on all registered SpecListenable channels and returns a channel for errors.
func (server *Server) ListenAndServeSpecListenable() chan sdkv2alphalib.SpecListenableErr {
	listeners := server.Bindings.RegisteredListenableChannels
	listenerErr := make(chan sdkv2alphalib.SpecListenableErr, len(listeners))

	for key, listener := range listeners {
		ctx := context.Background()
		go listener.Listen(ctx, listenerErr)

		fmt.Println("Registered Listenable: " + key)
	}
	return listenerErr
}

//&http.Client{
//Transport: &http.Transport{
//DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
//return svc.Dial("tcp", "192.168.100.5:4222")
//},
//},
//}
//get, err := hc.Get("/")
//if err != nil {
//fmt.Println("GET ERROR")
//fmt.Println(err)
//return
//}
//
//fmt.Println(get.Status)
//
//conn, err := svc.Dial("tcp", "192.168.100.5:4222")
//
//if err != nil {
//fmt.Println("NEBULA BOUND ERROR")
//fmt.Println(err)
//return
//}
//
//// buffer to get data
//received := make([]byte, 1024)
//_, err = conn.Read(received)
//if err != nil {
//println("Read data failed:", err.Error())
//os.Exit(1)
//}
//
//println("Received message:", string(received))
//defer conn.Close()
//
//fmt.Println("NEBULA BOUND3")
//
//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
//_, err = bufio.NewReader(conn).ReadString('\n')
