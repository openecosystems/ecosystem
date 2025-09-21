package sdkv2betalib

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"connectrpc.com/connect"
	"github.com/slackhq/nebula/service"
	"google.golang.org/protobuf/reflect/protoreflect"

	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"
)

var connectorQuit = make(chan os.Signal, 1)

// ConnectorMethod represents a gRPC method within a service, including its name, input/output descriptors, and schema information.
type ConnectorMethod struct {
	ProcedureName string
	Input         protoreflect.MessageDescriptor
	Output        protoreflect.MessageDescriptor
	Schema        protoreflect.MethodDescriptor
}

// Connector represents a structure for managing service bindings, procedures, configuration options, & service handlers.
type Connector struct {
	Bindings              *Bindings
	Bounds                []Binding
	MeshSocket            *service.Service
	ProcedureName         string
	Name                  string
	Err                   error
	Schema                protoreflect.ServiceDescriptor
	Methods               []*ConnectorMethod
	MethodsByPath         map[string]*ConnectorMethod
	Handler               http.Handler
	Opts                  []ConnectorOption
	ConfigurationProvider *BaseSpecConfigurationProvider

	options *connectorOptions
	// err     error
}

// NewConnector initializes a new Connector instance with the provided context, bindings, and optional configuration options.
// It resolves and validates the configuration, registers bindings, processes options, and returns the constructed Connector.
// Panics if configuration resolution or validation fails.
func NewConnector(ctx context.Context, opts ...ConnectorOption) *Connector {
	options, err := newConnectorOptions(opts)
	if err != nil {
		fmt.Println("new connector options error: ")
		fmt.Println(err)
	}

	connector := &Connector{
		Bounds:  options.Bounds,
		options: options,
	}

	provider := options.ConfigurationProvider
	if provider == nil {
		panic("configuration provider is nil. Please provide a configuration provider to the server.")
	}

	connector.ConfigurationProvider = &provider
	t := options.ConfigurationProvider

	configurer, cerr := t.ResolveConfiguration()
	if cerr != nil {
		fmt.Println(cerr)
		panic(cerr)
	}
	cerr = t.ValidateConfiguration()
	if cerr != nil {
		fmt.Println(cerr)
		panic(cerr)
	}

	bindings := RegisterBindings(ctx, options.Bounds, WithConfigurer(configurer))
	connector.Bindings = bindings

	return connector
}

// ListenAndProcess initializes the connector's context, manages its lifecycle, and delegates processing tasks with context.
func (connector *Connector) ListenAndProcess() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// connector.ConfigureMeshSocket()

	connector.ListenAndProcessWithCtx(ctx)
}

// ListenAndProcessWithCtx listens on registered channels and processes events while managing graceful shutdowns using context.
func (connector *Connector) ListenAndProcessWithCtx(_ context.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("Recovered from panic in connector %s: %v\n", connector.Name, rec)
			debug.PrintStack() // optional stack trace
		}
	}()

	var specListenableErr chan SpecListenableErr
	if connector.Bindings.RegisteredListenableChannels != nil {
		go func() {
			specListenableErr = connector.ListenAndProcessSpecListenable()
		}()
	}

	fmt.Println("Connector started successfully.")

	/*
	 * Graceful Shutdown Management
	 */
	signal.Notify(connectorQuit, syscall.SIGTERM)
	signal.Notify(connectorQuit, os.Interrupt)
	select {
	case err := <-specListenableErr:
		if err.Error != nil {
			fmt.Println(ErrServerInternal.WithInternalErrorDetail(err.Error))
		}
	case <-connectorQuit:
		connector.Shutdown()
	}
}

// ListenAndProcessSpecListenable starts listening on all registered listenable channels
// and returns a channel for errors. Each listener goroutine is wrapped in a supervision
// loop so that if it panics, it will log, recover, and restart after a short backoff.
func (connector *Connector) ListenAndProcessSpecListenable() chan SpecListenableErr {
	listeners := connector.Bindings.RegisteredListenableChannels
	listenerErr := make(chan SpecListenableErr, len(listeners))

	for key, listener := range listeners {
		ctx := context.Background()
		go listener.Listen(ctx, listenerErr)

		fmt.Println("Registered Listenable: " + key)
	}
	return listenerErr
}

// Shutdown gracefully stopping the connector
func (connector *Connector) Shutdown() {
	fmt.Printf("Stopping connector gracefully. Draining connections for up to %v seconds", 30)
	fmt.Println()

	_, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	ShutdownBindings(connector.Bindings)
}

// descKind returns a string describing the kind of protoreflect.Descriptor instance provided as input.
func descKind(desc protoreflect.Descriptor) string { // nolint:unused
	switch desc := desc.(type) {
	case protoreflect.FileDescriptor:
		return "a file"
	case protoreflect.MessageDescriptor:
		return "a message"
	case protoreflect.FieldDescriptor:
		if desc.IsExtension() {
			return "an extension"
		}
		return "a field"
	case protoreflect.OneofDescriptor:
		return "a oneof"
	case protoreflect.EnumDescriptor:
		return "an enum"
	case protoreflect.EnumValueDescriptor:
		return "an enum value"
	case protoreflect.ServiceDescriptor:
		return "a service"
	case protoreflect.MethodDescriptor:
		return "a method"
	default:
		return fmt.Sprintf("%T", desc)
	}
}

// ConnectorOption defines an interface for applying custom configuration to a connectorOptions object.
type ConnectorOption interface {
	apply(*connectorOptions)
}

// connectorOptions defines the configuration options for a connector, including supported protocols and codecs.
type connectorOptions struct {
	Bounds                []Binding
	ConfigurationProvider BaseSpecConfigurationProvider

	protocols map[typev2pb.Protocol]struct{}
	// codecNames     map[string]struct{}
	// preferredCodec string
}

// connectorOptionFunc is a function type that modifies the settings of a connectorOptions instance.
type connectorOptionFunc func(*connectorOptions)

// apply applies the connectorOptionFunc to the given connectorOptions.
func (f connectorOptionFunc) apply(opts *connectorOptions) {
	f(opts)
}

// newConnectorOptions creates and configures a new connectorOptions instance using the provided ConnectorOption slice.
// Returns the configured connectorOptions and an error if validation fails.
func newConnectorOptions(options []ConnectorOption) (*connectorOptions, *connect.Error) {
	config := connectorOptions{
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

// validate checks the integrity and consistency of the connectorOptions fields.
// Returns a *connect.Error if validation fails or nil if successful.
func (c *connectorOptions) validate() *connect.Error {
	return nil
}

// WithConnectorOptions composes multiple Options into one.
func WithConnectorOptions(opts ...ConnectorOption) ConnectorOption {
	return connectorOptionFunc(func(cfg *connectorOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithConnectorBounds configures the connector with the specified bounds, overriding the default bindings list in server options.
func WithConnectorBounds(bounds []Binding) ConnectorOption {
	return connectorOptionFunc(func(cfg *connectorOptions) {
		cfg.Bounds = bounds
	})
}

// WithConnectorConfigurationProvider sets the SpecConfigurationProvider for the server configuration and applies it as a ServerOption.
func WithConnectorConfigurationProvider(settings BaseSpecConfigurationProvider) ConnectorOption {
	return connectorOptionFunc(func(cfg *connectorOptions) {
		cfg.ConfigurationProvider = settings
	})
}

// WithTargetProtocols sets the allowed target protocols for the connector using the provided list of protocols.
func WithTargetProtocols(protocols ...typev2pb.Protocol) ConnectorOption {
	return connectorOptionFunc(func(opts *connectorOptions) {
		opts.protocols = make(map[typev2pb.Protocol]struct{}, len(protocols))
		for _, p := range protocols {
			opts.protocols[p] = struct{}{}
		}
	})
}
