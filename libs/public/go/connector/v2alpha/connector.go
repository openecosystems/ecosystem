package connectorv2alphalib

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
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/protobuf/go/protobuf/gen/platform/type/v2"
	v2alpha "libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	"libs/public/go/sdk/v2alpha"
)

var quit = make(chan os.Signal, 1)

type Method struct {
	ProcedureName string
	Input         protoreflect.MessageDescriptor
	Output        protoreflect.MessageDescriptor
	Schema        protoreflect.MethodDescriptor
}

type Connector struct {
	Bindings      *sdkv2alphalib.Bindings
	Bounds        []sdkv2alphalib.Binding
	MeshSocket    *service.Service
	ProcedureName string
	Name          string
	Err           error
	Schema        protoreflect.ServiceDescriptor
	Methods       []*Method
	MethodsByPath map[string]*Method
	Handler       http.Handler
	Opts          []ConnectorOption

	options *connectorOptions
	err     error
}

func NewConnector(ctx context.Context, bounds []sdkv2alphalib.Binding, opts ...ConnectorOption) *Connector {
	c := Configuration{}
	c.ResolveConfiguration()
	err := c.ValidateConfiguration()
	if err != nil {
		fmt.Println("validate connector configuration error: ", err)
		panic(err)
	}

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

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

func NewDynamicConnectorWithSchema(ctx context.Context, service protoreflect.ServiceDescriptor, bounds []sdkv2alphalib.Binding, opts ...ConnectorOption) *Connector {
	procedureName := "/" + string(service.FullName()) + "/"
	methods := make([]*Method, 0, service.Methods().Len())
	for j := 0; j < service.Methods().Len(); j++ {
		method := service.Methods().Get(j)
		// fmt.Printf("  Method Name: %s\n", method.Name())

		methodProcedureName := procedureName + string(method.Name())
		methods = append(methods, &Method{
			ProcedureName: methodProcedureName,
			Input:         method.Input(),
			Output:        method.Output(),
			Schema:        method,
		})
	}

	mbp := make(map[string]*Method)
	for _, method := range methods {
		mbp[method.ProcedureName] = method
	}

	c := Configuration{}
	c.ResolveConfiguration()

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

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

func NewDynamicConnector(ctx context.Context, servicePath string, bounds []sdkv2alphalib.Binding, opts ...ConnectorOption) *Connector {
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

func (connector *Connector) ListenAndProcess() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// connector.ConfigureMeshSocket()

	connector.ListenAndProcessWithCtx(ctx)
}

func (connector *Connector) ListenAndProcessWithCtx(_ context.Context) {
	var specListenableErr chan sdkv2alphalib.SpecListenableErr
	if connector.Bindings.RegisteredListenableChannels != nil {
		go func() {
			specListenableErr = connector.ListenAndProcessSpecListenable()
		}()
	}

	fmt.Println("Connector started successfully.")

	/*
	 * Graceful Shutdown Management
	 */
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	select {
	case err := <-specListenableErr:
		if err.Error != nil {
			fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err.Error))
		}
	case <-quit:
		fmt.Printf("Stopping connector gracefully. Draining connections for up to %v seconds", 30)
		fmt.Println()

		_, cancel := context.WithTimeout(context.Background(), 30)
		defer cancel()

		sdkv2alphalib.ShutdownBindings(connector.Bindings)

	}
}

func (connector *Connector) ListenAndProcessSpecListenable() chan sdkv2alphalib.SpecListenableErr {
	listeners := connector.Bindings.RegisteredListenableChannels
	listenerErr := make(chan sdkv2alphalib.SpecListenableErr, len(listeners))

	for key, listener := range listeners {

		ctx := context.Background()
		go listener.Listen(ctx, listenerErr)

		fmt.Println("Registered Listenable: " + key)

	}
	return listenerErr
}
