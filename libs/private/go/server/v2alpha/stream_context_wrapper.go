package serverv2alpha

import (
	"context"

	"google.golang.org/grpc"
)

// StreamContextWrapper is an interface that combines grpc.ServerStream with the ability to set a custom context.
type StreamContextWrapper interface {
	grpc.ServerStream
	SetContext(context.Context)
}

// wrapper is a struct implementing the StreamContextWrapper interface for managing gRPC stream context and metadata.
//
//nolint:unused
type wrapper struct {
	grpc.ServerStream
	info *grpc.StreamServerInfo
	ctx  context.Context
}

// Context returns the context associated with the wrapper instance.
//
//nolint:unused
func (w *wrapper) Context() context.Context {
	return w.ctx
}

// SetContext sets a new context for the wrapper. This method updates the wrapper's internal context value.
//
//nolint:unused
func (w *wrapper) SetContext(ctx context.Context) {
	w.ctx = ctx
}

// newStreamContextWrapper wraps a grpc.ServerStream with additional context and server information.
// Returns a StreamContextWrapper, allowing context to be set and retrieved.
//
//nolint:unused
func newStreamContextWrapper(inner grpc.ServerStream, info *grpc.StreamServerInfo) StreamContextWrapper {
	ctx := inner.Context()
	return &wrapper{
		inner,
		info,
		ctx,
	}
}
