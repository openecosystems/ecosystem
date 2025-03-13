package sdkv2alphalib

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"connectrpc.com/connect"
	"github.com/slackhq/nebula/service"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	v2alpha "github.com/openecosystems/ecosystem/libs/public/go/protobuf/gen/platform/configuration/v2alpha"
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
	Bindings      *Bindings
	Bounds        []Binding
	MeshSocket    *service.Service
	ProcedureName string
	Name          string
	Err           error
	Schema        protoreflect.ServiceDescriptor
	Methods       []*ConnectorMethod
	MethodsByPath map[string]*ConnectorMethod
	Handler       http.Handler
	Opts          []ConnectorOption

	options *connectorOptions
	// err     error
}

// NewConnector initializes a new Connector instance with the provided context, bindings, and optional configuration options.
// It resolves and validates the configuration, registers bindings, processes options, and returns the constructed Connector.
// Panics if configuration resolution or validation fails.
func NewConnector(ctx context.Context, bounds []Binding, opts ...ConnectorOption) *Connector {
	//c := Configuration{}
	//c.ResolveConfiguration()
	//err := c.ValidateConfiguration()
	//if err != nil {
	//	fmt.Println("validate connector configuration error: ", err)
	//	panic(err)
	//}

	bindings := RegisterBindings(ctx, bounds)

	options, err := newConnectorOptions(opts)
	if err != nil {
		fmt.Println("new connector options error: ")
		fmt.Println(err)
	}

	return &Connector{
		Bindings: bindings,
		Bounds:   bounds,
		options:  options,
	}
}

// NewDynamicConnectorWithSchema creates a dynamically configured Connector using the provided schema, bindings, and options.
func NewDynamicConnectorWithSchema(ctx context.Context, service protoreflect.ServiceDescriptor, bounds []Binding, opts ...ConnectorOption) *Connector {
	procedureName := "/" + string(service.FullName()) + "/"
	methods := make([]*ConnectorMethod, 0, service.Methods().Len())
	for j := 0; j < service.Methods().Len(); j++ {
		method := service.Methods().Get(j)
		// fmt.Printf("  Method Name: %s\n", method.Name())

		methodProcedureName := procedureName + string(method.Name())
		methods = append(methods, &ConnectorMethod{
			ProcedureName: methodProcedureName,
			Input:         method.Input(),
			Output:        method.Output(),
			Schema:        method,
		})
	}

	mbp := make(map[string]*ConnectorMethod)
	for _, method := range methods {
		mbp[method.ProcedureName] = method
	}

	// c := Configuration{}
	// c.ResolveConfiguration()

	bindings := RegisterBindings(ctx, bounds)

	options, err := newConnectorOptions(opts)
	if err != nil {
		fmt.Println(err)
	}

	connector := &Connector{
		Bindings:      bindings,
		ProcedureName: procedureName,
		Name:          string(service.FullName()),
		Err:           nil,
		Schema:        service,
		Methods:       methods,
		MethodsByPath: mbp,
		Opts:          opts,

		options: options,
	}

	// TODO create a WithConnectOption option to allow to pass data directly to connect
	// connector.Handler = NewDynamicConnectorHandler(connector)

	return connector
}

// NewDynamicConnector creates a new instance of Connector based on the given service path, bindings, and optional configurations.
// It resolves the service schema dynamically and initializes the Connector with methods and bindings information.
// Returns a Connector, which may include an error if schema resolution fails.
func NewDynamicConnector(ctx context.Context, servicePath string, bounds []Binding, opts ...ConnectorOption) *Connector {
	serviceName := strings.TrimSuffix(strings.TrimPrefix(servicePath, "/"), "/")
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(serviceName))
	if err != nil {
		return &Connector{Err: fmt.Errorf("could not resolve schema for service at path %q: %w", servicePath, err)}
	}
	svcDesc, ok := desc.(protoreflect.ServiceDescriptor)
	if !ok {
		return &Connector{
			Err: fmt.Errorf("could not resolve schema for service at path %q: resolved descriptor is %s, not a service", servicePath, descKind(desc)),
		}
	}
	return NewDynamicConnectorWithSchema(ctx, svcDesc, bounds, opts...)
}

// func (ImplementedDynamicServiceHandler) DynamicUnary(context.Context, *connect.Request[dynamicpb.Message]) (*connect.Response[dynamicpb.Message], error) {

// DynamicUnary processes a CreateConfigurationRequest and returns a CreateConfigurationResponse or an error.
func (connector *Connector) DynamicUnary(_ context.Context, req *connect.Request[v2alpha.CreateConfigurationRequest]) (*connect.Response[v2alpha.CreateConfigurationResponse], error) {
	// fmt.Println(req.HTTPMethod())
	// fmt.Println(req.Spec().Schema)
	// fmt.Println(req.Spec().StreamType)
	fmt.Println(req.Spec().Procedure)
	// fmt.Println(req.Spec().IdempotencyLevel)
	// fmt.Println(req.Spec().IsClient)

	fmt.Println(req.Msg)

	return connect.NewResponse(&v2alpha.CreateConfigurationResponse{
		SpecContext: &specv2pb.SpecResponseContext{
			ResponseValidation: &typev2pb.ResponseValidation{
				ValidateOnly: true,
			},
			OrganizationSlug: "hello",
			WorkspaceSlug:    "world",
			WorkspaceJan:     1,
			RoutineId:        "123",
		},
		Configuration: &v2alpha.Configuration{
			Id:               "123",
			OrganizationSlug: "hello",
			WorkspaceSlug:    "world",
		},
	}), nil
	//
	//tracer := *opentelemetryv2.Bound.Tracer
	//log := *zaploggerv1.Bound.Logger
	//
	//// Get it from the GlobalSystem Registry
	//_, err := GlobalSystems.GetSystemByName(req.Spec().Procedure)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Executes top level validation, no business domain validation
	//validationCtx, validationSpan := tracer.Start(ctx, "request-validation", trace.WithSpanKind(trace.SpanKindInternal))
	//v := *protovalidatev0.Bound.Validator
	//if err := v.Validate(req.Msg); err != nil {
	//	return nil, ErrServerPreconditionFailed.WithInternalErrorDetail(err)
	//}
	//validationSpan.End()
	//
	//// Spec Propagation
	//specCtx, specSpan := tracer.Start(validationCtx, "spec-propagation", trace.WithSpanKind(trace.SpanKindInternal))
	//spec, ok := ctx.Value("spec").(*specv2pb.Spec)
	//if !ok {
	//	return nil, ErrServerInternal.WithInternalErrorDetail(errors.New("cannot propagate spec to context"))
	//}
	//specSpan.End()
	//
	//// Distributed Domain Handler
	//handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))
	//
	//entity := DynamicSpecEntity{}
	//reply, err2 := natsnodev2.Bound.MultiplexCommandSync(handlerCtx, spec, &natsnodev2.SpecCommand{
	//	Request:        req.Msg,
	//	Stream:         natsnodev2.NewInboundStream(),
	//	CommandName:    "",
	//	CommandTopic:   EventDataDynamicTopic,
	//	EntityTypeName: entity.TypeName(),
	//})
	//if err2 != nil {
	//	log.Error(err2.Error())
	//	return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	//}
	//
	//var dd v2alpha.CreateConfigurationResponse
	//err3 := proto.Unmarshal(reply.Data, &dd)
	//if err3 != nil {
	//	log.Error(err3.Error())
	//	return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	//}
	//
	//handlerSpan.End()
	//
	//return connect.NewResponse(&dd), nil
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
		fmt.Printf("Stopping connector gracefully. Draining connections for up to %v seconds", 30)
		fmt.Println()

		_, cancel := context.WithTimeout(context.Background(), 30)
		defer cancel()

		ShutdownBindings(connector.Bindings)
	}
}

// ListenAndProcessSpecListenable starts listening on all registered listenable channels and returns a channel for errors.
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
