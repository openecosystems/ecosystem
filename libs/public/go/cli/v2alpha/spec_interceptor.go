package cliv2alphalib

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// DecorateRequest modifies request headers with metadata values like timestamps, device details, configuration, and API context.
func DecorateRequest(_ context.Context, req connect.AnyRequest, settings *Configuration, overrides *sdkv2alphalib.RuntimeConfigurationOverrides) {
	req.Header().Set(sdkv2alphalib.SentAtKey, timestamppb.Now().String())
	req.Header().Set(sdkv2alphalib.RequestIdKey, "")
	req.Header().Set(sdkv2alphalib.DeviceIdKey, "mac")
	req.Header().Set(sdkv2alphalib.DeviceAdvertisingIdKey, "749393")
	req.Header().Set(sdkv2alphalib.DeviceManufacturerKey, sdkv2alphalib.OSData.Family)
	req.Header().Set(sdkv2alphalib.DeviceModelKey, "")
	req.Header().Set(sdkv2alphalib.DeviceNameKey, "")
	req.Header().Set(sdkv2alphalib.DeviceTypeKey, "")
	req.Header().Set(sdkv2alphalib.DeviceTokenKey, "")
	req.Header().Set(sdkv2alphalib.OsNameKey, sdkv2alphalib.OSData.Platform)
	req.Header().Set(sdkv2alphalib.OsVersionKey, sdkv2alphalib.OSData.PlatformVersion)

	if overrides != nil {
		req.Header().Set(sdkv2alphalib.FieldMask, overrides.FieldMask)
		req.Header().Set(sdkv2alphalib.ValidateOnlyKey, strconv.FormatBool(overrides.ValidateOnly))
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
func NewCLIInterceptor(settings *Configuration, overrides *sdkv2alphalib.RuntimeConfigurationOverrides) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			DecorateRequest(ctx, req, settings, overrides)
			return next(ctx, req)
		}
	}
	return interceptor
}

// NewApplyHeadersInterceptor creates a unary interceptor to apply headers to requests using provided SpecSettings.
func NewApplyHeadersInterceptor(settings *Configuration) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			DecorateRequest(ctx, req, settings, &sdkv2alphalib.RuntimeConfigurationOverrides{})
			return next(ctx, req)
		}
	}
	return interceptor
}
