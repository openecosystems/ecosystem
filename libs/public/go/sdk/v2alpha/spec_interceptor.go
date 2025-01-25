package sdkv2alphalib

import (
	"context"
	"strconv"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"connectrpc.com/connect"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// SpecInterceptor represents a structure that contains the platform specification defined by specv2pb.Spec.
type SpecInterceptor struct {
	// spec specv2pb.Spec // The Platform Spec
}

// DecorateContext adds a "spec" value to the provided context based on the information extracted from the given request.
func DecorateContext(ctx context.Context, req connect.AnyRequest) context.Context {
	factory := NewFactory(req)
	s := factory.Spec

	ctx = context.WithValue(ctx, SpecContextKey, s)

	return ctx
}

// DecorateRequest modifies request headers with metadata values like timestamps, device details, configuration, and API context.
func DecorateRequest(_ context.Context, req connect.AnyRequest, settings *specv2pb.SpecSettings, overrides *RuntimeConfigurationOverrides) {
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

	if settings != nil && settings.Context != nil {
		req.Header().Set(ApiKey, settings.Context.ApiKey)
		for _, header := range settings.Context.Headers {
			for _, v := range header.Values {
				req.Header().Add(header.Key, v)
			}
		}
	}
}

// NewSpecInterceptor creates a unary interceptor that decorates the context with a specification derived from the request.
func NewSpecInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = DecorateContext(ctx, req)
			return next(ctx, req)
		}
	}
	return interceptor
}

// NewApplyHeadersInterceptor creates a unary interceptor to apply headers to requests using provided SpecSettings.
func NewApplyHeadersInterceptor(settings *specv2pb.SpecSettings) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			DecorateRequest(ctx, req, settings, &RuntimeConfigurationOverrides{})
			return next(ctx, req)
		}
	}
	return interceptor
}

// NewCLIInterceptor creates a unary interceptor function to decorate requests with headers based on given settings and overrides.
func NewCLIInterceptor(settings *specv2pb.SpecSettings, overrides *RuntimeConfigurationOverrides) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			DecorateRequest(ctx, req, settings, overrides)
			return next(ctx, req)
		}
	}
	return interceptor
}
