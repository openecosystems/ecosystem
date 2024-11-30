package sdkv2alphalib

import (
	"connectrpc.com/connect"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	specproto "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"strconv"
)

type SpecInterceptor struct {
	spec specproto.Spec // The Platform Spec
}

func DecorateContext(ctx context.Context, req connect.AnyRequest) context.Context {

	factory := NewFactory(req)
	s := factory.Spec

	ctx = context.WithValue(ctx, "spec", s)

	return ctx
}

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

func NewSpecInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = DecorateContext(ctx, req)
			return next(ctx, req)
		}
	}
	return interceptor
}

func NewApplyHeadersInterceptor(settings *specv2pb.SpecSettings) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			DecorateRequest(ctx, req, settings, &RuntimeConfigurationOverrides{})
			return next(ctx, req)
		}
	}
	return interceptor
}

func NewCLIInterceptor(settings *specv2pb.SpecSettings, overrides *RuntimeConfigurationOverrides) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			DecorateRequest(ctx, req, settings, overrides)
			return next(ctx, req)
		}
	}
	return interceptor
}
