package cryptographyv2alphasrv

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"go.opentelemetry.io/otel/trace"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	"libs/partner/go/opentelemetry/v2"
	tinkv2 "libs/partner/go/tink/v2"
	zaploggerv1 "libs/partner/go/zap/v1"
	_ "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
)

type EncryptionServiceHandler struct{}

func (s *EncryptionServiceHandler) Encrypt(ctx context.Context, req *connect.Request[cryptographyv2alphapb.EncryptRequest]) (*connect.Response[cryptographyv2alphapb.EncryptResponse], error) {
	tracer := *opentelemetryv2.Bound.Tracer
	log := *zaploggerv1.Bound.Logger
	_ = *tinkv2.Bound

	fmt.Println(req)

	log.Info("Encrypting...")

	ctx, span := tracer.Start(ctx, "encrypt", trace.WithSpanKind(trace.SpanKindInternal))

	response := cryptographyv2alphapb.EncryptResponse{
		Result: &cryptographyv2alphapb.EncryptResponse_Err{Err: "error response"},
	}

	span.End()

	return connect.NewResponse(&response), nil
}

func (s *EncryptionServiceHandler) Decrypt(ctx context.Context, req *connect.Request[cryptographyv2alphapb.DecryptRequest]) (*connect.Response[cryptographyv2alphapb.DecryptResponse], error) {
	tracer := *opentelemetryv2.Bound.Tracer
	// log := *zaploggerv1.Bound.Logger

	fmt.Println(req)

	ctx, span := tracer.Start(ctx, "encrypt", trace.WithSpanKind(trace.SpanKindInternal))

	response := cryptographyv2alphapb.DecryptResponse{
		Result: &cryptographyv2alphapb.DecryptResponse_Err{Err: "error response"},
	}

	span.End()

	return connect.NewResponse(&response), nil
}
