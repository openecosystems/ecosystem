package sdkv2betalib

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"connectrpc.com/connect"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SpecInterceptor represents a structure that contains the platform specification defined by specv2pb.Spec.
type SpecInterceptor struct{}

// NewSpecInterceptor create an interceptor
func NewSpecInterceptor() *SpecInterceptor {
	return &SpecInterceptor{}
}

// WrapUnary wrap interceptor
func (i *SpecInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		// Before Call
		ctx = DecorateContext(ctx, req.Header(), req.Spec().Procedure)
		conn, err := next(ctx, req)
		// After Call
		if err != nil {
			err2 := i.HumanizeResponse(ctx, err)
			return nil, &err2
		}

		return conn, nil
	}
}

// WrapStreamingClient wrap interceptor for streaming clients
func (*SpecInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		// Before Call
		conn := next(ctx, spec)
		// After Call
		return conn
	}
}

// WrapStreamingHandler server side
func (i *SpecInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		// Before Call
		ctx = DecorateContext(ctx, conn.RequestHeader(), conn.Spec().Procedure)
		err := next(ctx, conn)
		// After Call
		return err
	}
}

// DecorateContext adds a "spec" value to the provided context based on the information extracted from the given request.
func DecorateContext(ctx context.Context, h http.Header, procedure string) context.Context {
	factory := NewFactory(ctx, h, procedure)
	s := factory.Spec

	ctx = context.WithValue(ctx, SpecContextKey, s)

	return ctx
}

// HumanizeResponse humanize the response
func (i *SpecInterceptor) HumanizeResponse(ctx context.Context, err error) connect.Error {
	var specErr SpecError
	var requestInfo *errdetails.RequestInfo

	val := ctx.Value(SpecContextKey)
	spec, ok := val.(*specv2pb.Spec)
	if ok {
		if spec != nil && spec.GetSpanContext() != nil && spec.GetSpanContext().GetTraceId() != "" {
			requestInfo = &errdetails.RequestInfo{
				RequestId:   spec.GetSpanContext().GetTraceId(),
				ServingData: "build: ; version: v2.0",
			}
		}
	}

	if errors.As(err, &specErr) {
		specErr = specErr.WithRequestInfo(requestInfo)

		return specErr.ConnectErr
	}

	specErr = ErrServerInternal.WithRequestInfo(requestInfo)
	return specErr.ConnectErr

	//if merr := multierr.Errors(err); len(merr) > 1 {
	//	if specErr != nil && specErr.ConnectErr != nil {
	//		return specErr.ConnectErr
	//	}
	//}
}

//// NewSpecInterceptor creates a unary interceptor that decorates the context with a specification derived from the request.
//func NewSpecInterceptor() connect.UnaryInterceptorFunc {
//	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
//		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
//			ctx = DecorateContext(ctx, req)
//			return next(ctx, req)
//		}
//	}
//	return interceptor
//}

// DecorateCLIRequest modifies request headers with metadata values like timestamps, device details, configuration, and API context.
func DecorateCLIRequest(_ context.Context, req connect.AnyRequest, settings *CLIConfiguration, overrides *RuntimeConfigurationOverrides) {
	req.Header().Set(SentAtKey, timestamppb.Now().String())
	req.Header().Set(RequestIdKey, "")
	req.Header().Set(DeviceIdKey, "mac")
	req.Header().Set(DeviceAdvertisingIdKey, "749393")
	req.Header().Set(DeviceManufacturerKey, OSData.Family)
	req.Header().Set(DeviceModelKey, "")
	req.Header().Set(DeviceNameKey, "")
	req.Header().Set(DeviceTypeKey, "")
	req.Header().Set(DeviceTokenKey, "")
	req.Header().Set(OsNameKey, OSData.Platform)
	req.Header().Set(OsVersionKey, OSData.PlatformVersion)

	if overrides != nil {
		req.Header().Set(FieldMask, overrides.FieldMask)
		req.Header().Set(ValidateOnlyKey, strconv.FormatBool(overrides.ValidateOnly))
	}

	if settings != nil {
		// req.Header().Set(ApiKey, settings.Context.ApiKey)
		for _, header := range settings.Context.Headers {
			for _, v := range header.Values {
				req.Header().Add(header.Key, v)
			}
		}
	}
}

// NewCLIInterceptor creates a unary interceptor function to decorate requests with headers based on given settings and overrides.
func NewCLIInterceptor(settings *CLIConfiguration, overrides *RuntimeConfigurationOverrides) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Before Call
			DecorateCLIRequest(ctx, req, settings, overrides)
			conn, err := next(ctx, req)
			// After Call
			return conn, err
		}
	}
	return interceptor
}

// NewApplyHeadersInterceptor creates a unary interceptor to apply headers to requests using provided SpecSettings.
func NewApplyHeadersInterceptor(settings *CLIConfiguration) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Before Call
			DecorateCLIRequest(ctx, req, settings, &RuntimeConfigurationOverrides{})
			conn, err := next(ctx, req)
			// After Call
			return conn, err
		}
	}
	return interceptor
}
