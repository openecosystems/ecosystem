package sdkv2alphalib

import (
	"context"

	"connectrpc.com/connect"
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
