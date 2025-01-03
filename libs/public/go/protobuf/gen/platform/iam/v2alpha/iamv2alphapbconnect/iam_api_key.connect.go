// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: platform/iam/v2alpha/iam_api_key.proto

package iamv2alphapbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v2alpha "libs/public/go/protobuf/gen/platform/iam/v2alpha"
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
	// IamApiKeyServiceName is the fully-qualified name of the IamApiKeyService service.
	IamApiKeyServiceName = "platform.iam.v2alpha.IamApiKeyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IamApiKeyServiceCreateApiKeyProcedure is the fully-qualified name of the IamApiKeyService's
	// CreateApiKey RPC.
	IamApiKeyServiceCreateApiKeyProcedure = "/platform.iam.v2alpha.IamApiKeyService/CreateApiKey"
	// IamApiKeyServiceDeleteApiKeyProcedure is the fully-qualified name of the IamApiKeyService's
	// DeleteApiKey RPC.
	IamApiKeyServiceDeleteApiKeyProcedure = "/platform.iam.v2alpha.IamApiKeyService/DeleteApiKey"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	iamApiKeyServiceServiceDescriptor            = v2alpha.File_platform_iam_v2alpha_iam_api_key_proto.Services().ByName("IamApiKeyService")
	iamApiKeyServiceCreateApiKeyMethodDescriptor = iamApiKeyServiceServiceDescriptor.Methods().ByName("CreateApiKey")
	iamApiKeyServiceDeleteApiKeyMethodDescriptor = iamApiKeyServiceServiceDescriptor.Methods().ByName("DeleteApiKey")
)

// IamApiKeyServiceClient is a client for the platform.iam.v2alpha.IamApiKeyService service.
type IamApiKeyServiceClient interface {
	// Method to create API Key
	CreateApiKey(context.Context, *connect.Request[v2alpha.CreateApiKeyRequest]) (*connect.Response[v2alpha.CreateApiKeyResponse], error)
	// Method to logout
	DeleteApiKey(context.Context, *connect.Request[v2alpha.DeleteApiKeyRequest]) (*connect.Response[v2alpha.DeleteApiKeyResponse], error)
}

// NewIamApiKeyServiceClient constructs a client for the platform.iam.v2alpha.IamApiKeyService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIamApiKeyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IamApiKeyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &iamApiKeyServiceClient{
		createApiKey: connect.NewClient[v2alpha.CreateApiKeyRequest, v2alpha.CreateApiKeyResponse](
			httpClient,
			baseURL+IamApiKeyServiceCreateApiKeyProcedure,
			connect.WithSchema(iamApiKeyServiceCreateApiKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteApiKey: connect.NewClient[v2alpha.DeleteApiKeyRequest, v2alpha.DeleteApiKeyResponse](
			httpClient,
			baseURL+IamApiKeyServiceDeleteApiKeyProcedure,
			connect.WithSchema(iamApiKeyServiceDeleteApiKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// iamApiKeyServiceClient implements IamApiKeyServiceClient.
type iamApiKeyServiceClient struct {
	createApiKey *connect.Client[v2alpha.CreateApiKeyRequest, v2alpha.CreateApiKeyResponse]
	deleteApiKey *connect.Client[v2alpha.DeleteApiKeyRequest, v2alpha.DeleteApiKeyResponse]
}

// CreateApiKey calls platform.iam.v2alpha.IamApiKeyService.CreateApiKey.
func (c *iamApiKeyServiceClient) CreateApiKey(ctx context.Context, req *connect.Request[v2alpha.CreateApiKeyRequest]) (*connect.Response[v2alpha.CreateApiKeyResponse], error) {
	return c.createApiKey.CallUnary(ctx, req)
}

// DeleteApiKey calls platform.iam.v2alpha.IamApiKeyService.DeleteApiKey.
func (c *iamApiKeyServiceClient) DeleteApiKey(ctx context.Context, req *connect.Request[v2alpha.DeleteApiKeyRequest]) (*connect.Response[v2alpha.DeleteApiKeyResponse], error) {
	return c.deleteApiKey.CallUnary(ctx, req)
}

// IamApiKeyServiceHandler is an implementation of the platform.iam.v2alpha.IamApiKeyService
// service.
type IamApiKeyServiceHandler interface {
	// Method to create API Key
	CreateApiKey(context.Context, *connect.Request[v2alpha.CreateApiKeyRequest]) (*connect.Response[v2alpha.CreateApiKeyResponse], error)
	// Method to logout
	DeleteApiKey(context.Context, *connect.Request[v2alpha.DeleteApiKeyRequest]) (*connect.Response[v2alpha.DeleteApiKeyResponse], error)
}

// NewIamApiKeyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIamApiKeyServiceHandler(svc IamApiKeyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	iamApiKeyServiceCreateApiKeyHandler := connect.NewUnaryHandler(
		IamApiKeyServiceCreateApiKeyProcedure,
		svc.CreateApiKey,
		connect.WithSchema(iamApiKeyServiceCreateApiKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	iamApiKeyServiceDeleteApiKeyHandler := connect.NewUnaryHandler(
		IamApiKeyServiceDeleteApiKeyProcedure,
		svc.DeleteApiKey,
		connect.WithSchema(iamApiKeyServiceDeleteApiKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/platform.iam.v2alpha.IamApiKeyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IamApiKeyServiceCreateApiKeyProcedure:
			iamApiKeyServiceCreateApiKeyHandler.ServeHTTP(w, r)
		case IamApiKeyServiceDeleteApiKeyProcedure:
			iamApiKeyServiceDeleteApiKeyHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIamApiKeyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIamApiKeyServiceHandler struct{}

func (UnimplementedIamApiKeyServiceHandler) CreateApiKey(context.Context, *connect.Request[v2alpha.CreateApiKeyRequest]) (*connect.Response[v2alpha.CreateApiKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("platform.iam.v2alpha.IamApiKeyService.CreateApiKey is not implemented"))
}

func (UnimplementedIamApiKeyServiceHandler) DeleteApiKey(context.Context, *connect.Request[v2alpha.DeleteApiKeyRequest]) (*connect.Response[v2alpha.DeleteApiKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("platform.iam.v2alpha.IamApiKeyService.DeleteApiKey is not implemented"))
}
