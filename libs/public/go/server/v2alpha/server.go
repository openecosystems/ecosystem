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

	"connectrpc.com/vanguard"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"libs/public/go/sdk/v2alpha"
)

var quit = make(chan os.Signal, 1)

type Server struct {
	Bindings          *sdkv2alphalib.Bindings
	ConnectHttpServer *http2.Server
	HttpServerHandler *http.ServeMux
	Bounds            []sdkv2alphalib.Binding
	ServicePath       string
	ServiceHandler    *vanguard.Transcoder
	RawServiceHandler *http.Handler

	options *serverOptions
	err     error
}

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
		// MaxConcurrentStreams:         0,
		// PermitProhibitedCipherSuites: false,
		// MaxUploadBufferPerConnection: 0,
		// MaxUploadBufferPerStream:     0,
	}

	options, err := newServerOptions(path, opts)
	if err != nil {
		fmt.Println("new server options error: ", err)
	}

	s := vanguard.NewService(path, *handler)
	transcoder, err2 := vanguard.NewTranscoder([]*vanguard.Service{s})
	if err2 != nil {
		fmt.Println(err2)
	}

	return &Server{
		Bindings:          bindings,
		ConnectHttpServer: httpServer,
		Bounds:            bounds,
		ServicePath:       path,
		ServiceHandler:    transcoder,

		options: options,
		err:     errors.Join(err, err2),
	}
}

func NewMultiplexedServer(ctx context.Context, bounds []sdkv2alphalib.Binding, services []*vanguard.Service, opts ...ServerOption) *Server {
	c := Configuration{}
	c.ResolveConfiguration()
	err := c.ValidateConfiguration()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	httpServer := &http2.Server{}

	options, err := newServerOptions("", opts)
	if err != nil {
		fmt.Println(err)
	}

	transcoder, err2 := vanguard.NewTranscoder(services)
	if err2 != nil {
		fmt.Println(err2)
	}

	return &Server{
		Bindings:          bindings,
		ConnectHttpServer: httpServer,
		Bounds:            bounds,
		ServicePath:       "/",
		ServiceHandler:    transcoder,

		options: options,
		err:     errors.Join(err, err2),
	}
}

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
		// MaxConcurrentStreams:         0,
		// PermitProhibitedCipherSuites: false,
		// MaxUploadBufferPerConnection: 0,
		// MaxUploadBufferPerStream:     0,
	}

	options, err := newServerOptions(path, opts)
	if err != nil {
		fmt.Println(err)
	}

	return &Server{
		Bindings:          bindings,
		ConnectHttpServer: httpServer,
		Bounds:            bounds,
		ServicePath:       path,
		RawServiceHandler: handler,

		options: options,
		err:     errors.Join(err),
	}
}

func (server *Server) ListenAndServe() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server.ListenAndServeWithCtx(ctx)
}

func (server *Server) ListenAndServeWithCtx(_ context.Context) {
	httpServerErr := server.ListenAndServeMultiplexedHttp()

	var specListenableErr chan sdkv2alphalib.SpecListenableErr
	if server.Bindings.RegisteredListenableChannels != nil {
		go func() {
			specListenableErr = server.ListenAndServeSpecListenable()
		}()
	}

	fmt.Println("Server started successfully. HTTP listening on " + ResolvedConfiguration.Http.Port)

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

func (server *Server) ListenAndServeWithProvidedSocket(ln net.Listener) (httpServerErr chan error) {
	return server.listenAndServe(ln)
}

func (server *Server) ListenAndServeMultiplexedHttp() (httpServerErr chan error) {
	return server.listenAndServe(nil)
}

func (server *Server) listenAndServe(ln net.Listener) (httpServerErr chan error) {
	httpPort, _ := strconv.Atoi(ResolvedConfiguration.Http.Port)
	mux := http.NewServeMux()
	if server.HttpServerHandler != nil {
		mux = server.HttpServerHandler
	}

	server.HttpServerHandler = mux

	if server.RawServiceHandler != nil {
		mux.Handle(server.ServicePath, *server.RawServiceHandler)
	} else {
		mux.Handle("/", server.ServiceHandler)
	}

	_httpServerErr := make(chan error)
	go func() {
		if ln != nil {
			_httpServerErr <- http.Serve(
				ln,
				// Use h2c so we can serve HTTP/2 without TLS.
				h2c.NewHandler(mux, server.ConnectHttpServer),
			)
		} else {
			_httpServerErr <- http.ListenAndServe(
				fmt.Sprintf("0.0.0.0:%d", httpPort),
				// Use h2c so we can serve HTTP/2 without TLS.
				h2c.NewHandler(mux, server.ConnectHttpServer),
			)
		}
	}()

	return _httpServerErr
}

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
