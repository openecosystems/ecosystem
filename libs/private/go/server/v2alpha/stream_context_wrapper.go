package serverv2alpha

import (
	"context"

	"google.golang.org/grpc"
)

// StreamContextWrapper When using server streams, it isn't clear how to update the context for propagation. Simply creating a wrapper instead
type StreamContextWrapper interface {
	grpc.ServerStream
	SetContext(context.Context)
}

type wrapper struct {
	grpc.ServerStream
	info *grpc.StreamServerInfo
	ctx  context.Context
}

func (w *wrapper) Context() context.Context {
	return w.ctx
}

func (w *wrapper) SetContext(ctx context.Context) {
	w.ctx = ctx
}

func newStreamContextWrapper(inner grpc.ServerStream, info *grpc.StreamServerInfo) StreamContextWrapper {
	ctx := inner.Context()
	return &wrapper{
		inner,
		info,
		ctx,
	}
}
