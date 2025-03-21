// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: platform/system/v2alpha/system.proto

package systemv2alphapbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2alpha "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/system/v2alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SystemServiceName is the fully-qualified name of the SystemService service.
	SystemServiceName = "platform.system.v2alpha.SystemService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SystemServiceEnableProcedure is the fully-qualified name of the SystemService's Enable RPC.
	SystemServiceEnableProcedure = "/platform.system.v2alpha.SystemService/Enable"
	// SystemServiceDisableProcedure is the fully-qualified name of the SystemService's Disable RPC.
	SystemServiceDisableProcedure = "/platform.system.v2alpha.SystemService/Disable"
)

// SystemServiceClient is a client for the platform.system.v2alpha.SystemService service.
type SystemServiceClient interface {
	// Method to Subscribe to events based on scopes
	Enable(context.Context, *connect.Request[v2alpha.EnableRequest]) (*connect.ServerStreamForClient[v2alpha.EnableResponse], error)
	// Method to Unsubscribe to an event scope
	Disable(context.Context, *connect.Request[v2alpha.DisableRequest]) (*connect.Response[v2alpha.DisableResponse], error)
}

// NewSystemServiceClient constructs a client for the platform.system.v2alpha.SystemService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSystemServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SystemServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	systemServiceMethods := v2alpha.File_platform_system_v2alpha_system_proto.Services().ByName("SystemService").Methods()
	return &systemServiceClient{
		enable: connect.NewClient[v2alpha.EnableRequest, v2alpha.EnableResponse](
			httpClient,
			baseURL+SystemServiceEnableProcedure,
			connect.WithSchema(systemServiceMethods.ByName("Enable")),
			connect.WithClientOptions(opts...),
		),
		disable: connect.NewClient[v2alpha.DisableRequest, v2alpha.DisableResponse](
			httpClient,
			baseURL+SystemServiceDisableProcedure,
			connect.WithSchema(systemServiceMethods.ByName("Disable")),
			connect.WithClientOptions(opts...),
		),
	}
}

// systemServiceClient implements SystemServiceClient.
type systemServiceClient struct {
	enable  *connect.Client[v2alpha.EnableRequest, v2alpha.EnableResponse]
	disable *connect.Client[v2alpha.DisableRequest, v2alpha.DisableResponse]
}

// Enable calls platform.system.v2alpha.SystemService.Enable.
func (c *systemServiceClient) Enable(ctx context.Context, req *connect.Request[v2alpha.EnableRequest]) (*connect.ServerStreamForClient[v2alpha.EnableResponse], error) {
	return c.enable.CallServerStream(ctx, req)
}

// Disable calls platform.system.v2alpha.SystemService.Disable.
func (c *systemServiceClient) Disable(ctx context.Context, req *connect.Request[v2alpha.DisableRequest]) (*connect.Response[v2alpha.DisableResponse], error) {
	return c.disable.CallUnary(ctx, req)
}

// SystemServiceHandler is an implementation of the platform.system.v2alpha.SystemService service.
type SystemServiceHandler interface {
	// Method to Subscribe to events based on scopes
	Enable(context.Context, *connect.Request[v2alpha.EnableRequest], *connect.ServerStream[v2alpha.EnableResponse]) error
	// Method to Unsubscribe to an event scope
	Disable(context.Context, *connect.Request[v2alpha.DisableRequest]) (*connect.Response[v2alpha.DisableResponse], error)
}

// NewSystemServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSystemServiceHandler(svc SystemServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	systemServiceMethods := v2alpha.File_platform_system_v2alpha_system_proto.Services().ByName("SystemService").Methods()
	systemServiceEnableHandler := connect.NewServerStreamHandler(
		SystemServiceEnableProcedure,
		svc.Enable,
		connect.WithSchema(systemServiceMethods.ByName("Enable")),
		connect.WithHandlerOptions(opts...),
	)
	systemServiceDisableHandler := connect.NewUnaryHandler(
		SystemServiceDisableProcedure,
		svc.Disable,
		connect.WithSchema(systemServiceMethods.ByName("Disable")),
		connect.WithHandlerOptions(opts...),
	)
	return "/platform.system.v2alpha.SystemService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SystemServiceEnableProcedure:
			systemServiceEnableHandler.ServeHTTP(w, r)
		case SystemServiceDisableProcedure:
			systemServiceDisableHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSystemServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSystemServiceHandler struct{}

func (UnimplementedSystemServiceHandler) Enable(context.Context, *connect.Request[v2alpha.EnableRequest], *connect.ServerStream[v2alpha.EnableResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("platform.system.v2alpha.SystemService.Enable is not implemented"))
}

func (UnimplementedSystemServiceHandler) Disable(context.Context, *connect.Request[v2alpha.DisableRequest]) (*connect.Response[v2alpha.DisableResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("platform.system.v2alpha.SystemService.Disable is not implemented"))
}
