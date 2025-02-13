// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: kevel/advertisement/v1/decision.proto

package advertisementv1pbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "libs/partner/go/protobuf/gen/kevel/advertisement/v1"
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
	// DecisionServiceName is the fully-qualified name of the DecisionService service.
	DecisionServiceName = "kevel.advertisement.v1.DecisionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DecisionServiceGetDecisionsProcedure is the fully-qualified name of the DecisionService's
	// GetDecisions RPC.
	DecisionServiceGetDecisionsProcedure = "/kevel.advertisement.v1.DecisionService/GetDecisions"
	// DecisionServiceAddCustomPropertiesProcedure is the fully-qualified name of the DecisionService's
	// AddCustomProperties RPC.
	DecisionServiceAddCustomPropertiesProcedure = "/kevel.advertisement.v1.DecisionService/AddCustomProperties"
	// DecisionServiceAddInterestsProcedure is the fully-qualified name of the DecisionService's
	// AddInterests RPC.
	DecisionServiceAddInterestsProcedure = "/kevel.advertisement.v1.DecisionService/AddInterests"
	// DecisionServiceAddRetargetingSegmentProcedure is the fully-qualified name of the
	// DecisionService's AddRetargetingSegment RPC.
	DecisionServiceAddRetargetingSegmentProcedure = "/kevel.advertisement.v1.DecisionService/AddRetargetingSegment"
	// DecisionServiceOptOutProcedure is the fully-qualified name of the DecisionService's OptOut RPC.
	DecisionServiceOptOutProcedure = "/kevel.advertisement.v1.DecisionService/OptOut"
	// DecisionServiceReadProcedure is the fully-qualified name of the DecisionService's Read RPC.
	DecisionServiceReadProcedure = "/kevel.advertisement.v1.DecisionService/Read"
	// DecisionServiceIpOverrideProcedure is the fully-qualified name of the DecisionService's
	// IpOverride RPC.
	DecisionServiceIpOverrideProcedure = "/kevel.advertisement.v1.DecisionService/IpOverride"
	// DecisionServiceForgetProcedure is the fully-qualified name of the DecisionService's Forget RPC.
	DecisionServiceForgetProcedure = "/kevel.advertisement.v1.DecisionService/Forget"
	// DecisionServiceGdprConsentProcedure is the fully-qualified name of the DecisionService's
	// GdprConsent RPC.
	DecisionServiceGdprConsentProcedure = "/kevel.advertisement.v1.DecisionService/GdprConsent"
	// DecisionServiceMatchUserProcedure is the fully-qualified name of the DecisionService's MatchUser
	// RPC.
	DecisionServiceMatchUserProcedure = "/kevel.advertisement.v1.DecisionService/MatchUser"
)

// DecisionServiceClient is a client for the kevel.advertisement.v1.DecisionService service.
type DecisionServiceClient interface {
	GetDecisions(context.Context, *connect.Request[v1.GetDecisionsRequest]) (*connect.Response[v1.GetDecisionsResponse], error)
	AddCustomProperties(context.Context, *connect.Request[v1.AddCustomPropertiesRequest]) (*connect.Response[v1.AddCustomPropertiesResponse], error)
	AddInterests(context.Context, *connect.Request[v1.AddInterestsRequest]) (*connect.Response[v1.AddInterestsResponse], error)
	AddRetargetingSegment(context.Context, *connect.Request[v1.AddRetargetingSegmentRequest]) (*connect.Response[v1.AddRetargetingSegmentResponse], error)
	OptOut(context.Context, *connect.Request[v1.OptOutRequest]) (*connect.Response[v1.OptOutResponse], error)
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	IpOverride(context.Context, *connect.Request[v1.IpOverrideRequest]) (*connect.Response[v1.IpOverrideResponse], error)
	Forget(context.Context, *connect.Request[v1.ForgetRequest]) (*connect.Response[v1.ForgetResponse], error)
	GdprConsent(context.Context, *connect.Request[v1.GdprConsentRequest]) (*connect.Response[v1.GdprConsentResponse], error)
	MatchUser(context.Context, *connect.Request[v1.MatchUserRequest]) (*connect.Response[v1.MatchUserResponse], error)
}

// NewDecisionServiceClient constructs a client for the kevel.advertisement.v1.DecisionService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDecisionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DecisionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	decisionServiceMethods := v1.File_kevel_advertisement_v1_decision_proto.Services().ByName("DecisionService").Methods()
	return &decisionServiceClient{
		getDecisions: connect.NewClient[v1.GetDecisionsRequest, v1.GetDecisionsResponse](
			httpClient,
			baseURL+DecisionServiceGetDecisionsProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("GetDecisions")),
			connect.WithClientOptions(opts...),
		),
		addCustomProperties: connect.NewClient[v1.AddCustomPropertiesRequest, v1.AddCustomPropertiesResponse](
			httpClient,
			baseURL+DecisionServiceAddCustomPropertiesProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("AddCustomProperties")),
			connect.WithClientOptions(opts...),
		),
		addInterests: connect.NewClient[v1.AddInterestsRequest, v1.AddInterestsResponse](
			httpClient,
			baseURL+DecisionServiceAddInterestsProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("AddInterests")),
			connect.WithClientOptions(opts...),
		),
		addRetargetingSegment: connect.NewClient[v1.AddRetargetingSegmentRequest, v1.AddRetargetingSegmentResponse](
			httpClient,
			baseURL+DecisionServiceAddRetargetingSegmentProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("AddRetargetingSegment")),
			connect.WithClientOptions(opts...),
		),
		optOut: connect.NewClient[v1.OptOutRequest, v1.OptOutResponse](
			httpClient,
			baseURL+DecisionServiceOptOutProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("OptOut")),
			connect.WithClientOptions(opts...),
		),
		read: connect.NewClient[v1.ReadRequest, v1.ReadResponse](
			httpClient,
			baseURL+DecisionServiceReadProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("Read")),
			connect.WithClientOptions(opts...),
		),
		ipOverride: connect.NewClient[v1.IpOverrideRequest, v1.IpOverrideResponse](
			httpClient,
			baseURL+DecisionServiceIpOverrideProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("IpOverride")),
			connect.WithClientOptions(opts...),
		),
		forget: connect.NewClient[v1.ForgetRequest, v1.ForgetResponse](
			httpClient,
			baseURL+DecisionServiceForgetProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("Forget")),
			connect.WithClientOptions(opts...),
		),
		gdprConsent: connect.NewClient[v1.GdprConsentRequest, v1.GdprConsentResponse](
			httpClient,
			baseURL+DecisionServiceGdprConsentProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("GdprConsent")),
			connect.WithClientOptions(opts...),
		),
		matchUser: connect.NewClient[v1.MatchUserRequest, v1.MatchUserResponse](
			httpClient,
			baseURL+DecisionServiceMatchUserProcedure,
			connect.WithSchema(decisionServiceMethods.ByName("MatchUser")),
			connect.WithClientOptions(opts...),
		),
	}
}

// decisionServiceClient implements DecisionServiceClient.
type decisionServiceClient struct {
	getDecisions          *connect.Client[v1.GetDecisionsRequest, v1.GetDecisionsResponse]
	addCustomProperties   *connect.Client[v1.AddCustomPropertiesRequest, v1.AddCustomPropertiesResponse]
	addInterests          *connect.Client[v1.AddInterestsRequest, v1.AddInterestsResponse]
	addRetargetingSegment *connect.Client[v1.AddRetargetingSegmentRequest, v1.AddRetargetingSegmentResponse]
	optOut                *connect.Client[v1.OptOutRequest, v1.OptOutResponse]
	read                  *connect.Client[v1.ReadRequest, v1.ReadResponse]
	ipOverride            *connect.Client[v1.IpOverrideRequest, v1.IpOverrideResponse]
	forget                *connect.Client[v1.ForgetRequest, v1.ForgetResponse]
	gdprConsent           *connect.Client[v1.GdprConsentRequest, v1.GdprConsentResponse]
	matchUser             *connect.Client[v1.MatchUserRequest, v1.MatchUserResponse]
}

// GetDecisions calls kevel.advertisement.v1.DecisionService.GetDecisions.
func (c *decisionServiceClient) GetDecisions(ctx context.Context, req *connect.Request[v1.GetDecisionsRequest]) (*connect.Response[v1.GetDecisionsResponse], error) {
	return c.getDecisions.CallUnary(ctx, req)
}

// AddCustomProperties calls kevel.advertisement.v1.DecisionService.AddCustomProperties.
func (c *decisionServiceClient) AddCustomProperties(ctx context.Context, req *connect.Request[v1.AddCustomPropertiesRequest]) (*connect.Response[v1.AddCustomPropertiesResponse], error) {
	return c.addCustomProperties.CallUnary(ctx, req)
}

// AddInterests calls kevel.advertisement.v1.DecisionService.AddInterests.
func (c *decisionServiceClient) AddInterests(ctx context.Context, req *connect.Request[v1.AddInterestsRequest]) (*connect.Response[v1.AddInterestsResponse], error) {
	return c.addInterests.CallUnary(ctx, req)
}

// AddRetargetingSegment calls kevel.advertisement.v1.DecisionService.AddRetargetingSegment.
func (c *decisionServiceClient) AddRetargetingSegment(ctx context.Context, req *connect.Request[v1.AddRetargetingSegmentRequest]) (*connect.Response[v1.AddRetargetingSegmentResponse], error) {
	return c.addRetargetingSegment.CallUnary(ctx, req)
}

// OptOut calls kevel.advertisement.v1.DecisionService.OptOut.
func (c *decisionServiceClient) OptOut(ctx context.Context, req *connect.Request[v1.OptOutRequest]) (*connect.Response[v1.OptOutResponse], error) {
	return c.optOut.CallUnary(ctx, req)
}

// Read calls kevel.advertisement.v1.DecisionService.Read.
func (c *decisionServiceClient) Read(ctx context.Context, req *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return c.read.CallUnary(ctx, req)
}

// IpOverride calls kevel.advertisement.v1.DecisionService.IpOverride.
func (c *decisionServiceClient) IpOverride(ctx context.Context, req *connect.Request[v1.IpOverrideRequest]) (*connect.Response[v1.IpOverrideResponse], error) {
	return c.ipOverride.CallUnary(ctx, req)
}

// Forget calls kevel.advertisement.v1.DecisionService.Forget.
func (c *decisionServiceClient) Forget(ctx context.Context, req *connect.Request[v1.ForgetRequest]) (*connect.Response[v1.ForgetResponse], error) {
	return c.forget.CallUnary(ctx, req)
}

// GdprConsent calls kevel.advertisement.v1.DecisionService.GdprConsent.
func (c *decisionServiceClient) GdprConsent(ctx context.Context, req *connect.Request[v1.GdprConsentRequest]) (*connect.Response[v1.GdprConsentResponse], error) {
	return c.gdprConsent.CallUnary(ctx, req)
}

// MatchUser calls kevel.advertisement.v1.DecisionService.MatchUser.
func (c *decisionServiceClient) MatchUser(ctx context.Context, req *connect.Request[v1.MatchUserRequest]) (*connect.Response[v1.MatchUserResponse], error) {
	return c.matchUser.CallUnary(ctx, req)
}

// DecisionServiceHandler is an implementation of the kevel.advertisement.v1.DecisionService
// service.
type DecisionServiceHandler interface {
	GetDecisions(context.Context, *connect.Request[v1.GetDecisionsRequest]) (*connect.Response[v1.GetDecisionsResponse], error)
	AddCustomProperties(context.Context, *connect.Request[v1.AddCustomPropertiesRequest]) (*connect.Response[v1.AddCustomPropertiesResponse], error)
	AddInterests(context.Context, *connect.Request[v1.AddInterestsRequest]) (*connect.Response[v1.AddInterestsResponse], error)
	AddRetargetingSegment(context.Context, *connect.Request[v1.AddRetargetingSegmentRequest]) (*connect.Response[v1.AddRetargetingSegmentResponse], error)
	OptOut(context.Context, *connect.Request[v1.OptOutRequest]) (*connect.Response[v1.OptOutResponse], error)
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	IpOverride(context.Context, *connect.Request[v1.IpOverrideRequest]) (*connect.Response[v1.IpOverrideResponse], error)
	Forget(context.Context, *connect.Request[v1.ForgetRequest]) (*connect.Response[v1.ForgetResponse], error)
	GdprConsent(context.Context, *connect.Request[v1.GdprConsentRequest]) (*connect.Response[v1.GdprConsentResponse], error)
	MatchUser(context.Context, *connect.Request[v1.MatchUserRequest]) (*connect.Response[v1.MatchUserResponse], error)
}

// NewDecisionServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDecisionServiceHandler(svc DecisionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	decisionServiceMethods := v1.File_kevel_advertisement_v1_decision_proto.Services().ByName("DecisionService").Methods()
	decisionServiceGetDecisionsHandler := connect.NewUnaryHandler(
		DecisionServiceGetDecisionsProcedure,
		svc.GetDecisions,
		connect.WithSchema(decisionServiceMethods.ByName("GetDecisions")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceAddCustomPropertiesHandler := connect.NewUnaryHandler(
		DecisionServiceAddCustomPropertiesProcedure,
		svc.AddCustomProperties,
		connect.WithSchema(decisionServiceMethods.ByName("AddCustomProperties")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceAddInterestsHandler := connect.NewUnaryHandler(
		DecisionServiceAddInterestsProcedure,
		svc.AddInterests,
		connect.WithSchema(decisionServiceMethods.ByName("AddInterests")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceAddRetargetingSegmentHandler := connect.NewUnaryHandler(
		DecisionServiceAddRetargetingSegmentProcedure,
		svc.AddRetargetingSegment,
		connect.WithSchema(decisionServiceMethods.ByName("AddRetargetingSegment")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceOptOutHandler := connect.NewUnaryHandler(
		DecisionServiceOptOutProcedure,
		svc.OptOut,
		connect.WithSchema(decisionServiceMethods.ByName("OptOut")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceReadHandler := connect.NewUnaryHandler(
		DecisionServiceReadProcedure,
		svc.Read,
		connect.WithSchema(decisionServiceMethods.ByName("Read")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceIpOverrideHandler := connect.NewUnaryHandler(
		DecisionServiceIpOverrideProcedure,
		svc.IpOverride,
		connect.WithSchema(decisionServiceMethods.ByName("IpOverride")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceForgetHandler := connect.NewUnaryHandler(
		DecisionServiceForgetProcedure,
		svc.Forget,
		connect.WithSchema(decisionServiceMethods.ByName("Forget")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceGdprConsentHandler := connect.NewUnaryHandler(
		DecisionServiceGdprConsentProcedure,
		svc.GdprConsent,
		connect.WithSchema(decisionServiceMethods.ByName("GdprConsent")),
		connect.WithHandlerOptions(opts...),
	)
	decisionServiceMatchUserHandler := connect.NewUnaryHandler(
		DecisionServiceMatchUserProcedure,
		svc.MatchUser,
		connect.WithSchema(decisionServiceMethods.ByName("MatchUser")),
		connect.WithHandlerOptions(opts...),
	)
	return "/kevel.advertisement.v1.DecisionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DecisionServiceGetDecisionsProcedure:
			decisionServiceGetDecisionsHandler.ServeHTTP(w, r)
		case DecisionServiceAddCustomPropertiesProcedure:
			decisionServiceAddCustomPropertiesHandler.ServeHTTP(w, r)
		case DecisionServiceAddInterestsProcedure:
			decisionServiceAddInterestsHandler.ServeHTTP(w, r)
		case DecisionServiceAddRetargetingSegmentProcedure:
			decisionServiceAddRetargetingSegmentHandler.ServeHTTP(w, r)
		case DecisionServiceOptOutProcedure:
			decisionServiceOptOutHandler.ServeHTTP(w, r)
		case DecisionServiceReadProcedure:
			decisionServiceReadHandler.ServeHTTP(w, r)
		case DecisionServiceIpOverrideProcedure:
			decisionServiceIpOverrideHandler.ServeHTTP(w, r)
		case DecisionServiceForgetProcedure:
			decisionServiceForgetHandler.ServeHTTP(w, r)
		case DecisionServiceGdprConsentProcedure:
			decisionServiceGdprConsentHandler.ServeHTTP(w, r)
		case DecisionServiceMatchUserProcedure:
			decisionServiceMatchUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDecisionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDecisionServiceHandler struct{}

func (UnimplementedDecisionServiceHandler) GetDecisions(context.Context, *connect.Request[v1.GetDecisionsRequest]) (*connect.Response[v1.GetDecisionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.GetDecisions is not implemented"))
}

func (UnimplementedDecisionServiceHandler) AddCustomProperties(context.Context, *connect.Request[v1.AddCustomPropertiesRequest]) (*connect.Response[v1.AddCustomPropertiesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.AddCustomProperties is not implemented"))
}

func (UnimplementedDecisionServiceHandler) AddInterests(context.Context, *connect.Request[v1.AddInterestsRequest]) (*connect.Response[v1.AddInterestsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.AddInterests is not implemented"))
}

func (UnimplementedDecisionServiceHandler) AddRetargetingSegment(context.Context, *connect.Request[v1.AddRetargetingSegmentRequest]) (*connect.Response[v1.AddRetargetingSegmentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.AddRetargetingSegment is not implemented"))
}

func (UnimplementedDecisionServiceHandler) OptOut(context.Context, *connect.Request[v1.OptOutRequest]) (*connect.Response[v1.OptOutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.OptOut is not implemented"))
}

func (UnimplementedDecisionServiceHandler) Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.Read is not implemented"))
}

func (UnimplementedDecisionServiceHandler) IpOverride(context.Context, *connect.Request[v1.IpOverrideRequest]) (*connect.Response[v1.IpOverrideResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.IpOverride is not implemented"))
}

func (UnimplementedDecisionServiceHandler) Forget(context.Context, *connect.Request[v1.ForgetRequest]) (*connect.Response[v1.ForgetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.Forget is not implemented"))
}

func (UnimplementedDecisionServiceHandler) GdprConsent(context.Context, *connect.Request[v1.GdprConsentRequest]) (*connect.Response[v1.GdprConsentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.GdprConsent is not implemented"))
}

func (UnimplementedDecisionServiceHandler) MatchUser(context.Context, *connect.Request[v1.MatchUserRequest]) (*connect.Response[v1.MatchUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kevel.advertisement.v1.DecisionService.MatchUser is not implemented"))
}
