package main

import (
	"context"
	"fmt"

	opentelemetryv1 "github.com/openecosystems/ecosystem/libs/partner/go/opentelemetry"
	cryptographyv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/cryptography/v2alpha"

	"connectrpc.com/connect"

	tinkv2 "github.com/openecosystems/ecosystem/libs/partner/go/tink"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"

	"go.opentelemetry.io/otel/trace"
)

// EncryptionServiceHandler provides encryption and decryption functionality through predefined service methods.
type EncryptionServiceHandler struct{}

// Encrypt handles the encryption of plaintext and associated data, returning an EncryptResponse with the result or an error.
func (s *EncryptionServiceHandler) Encrypt(ctx context.Context, req *connect.Request[cryptographyv2alphapb.EncryptRequest]) (*connect.Response[cryptographyv2alphapb.EncryptResponse], error) {
	tracer := *opentelemetryv1.Bound.Tracer
	log := *zaploggerv1.Bound.Logger
	_ = *tinkv2.Bound

	fmt.Println(req)

	log.Info("Encrypting...")

	_, span := tracer.Start(ctx, "encrypt", trace.WithSpanKind(trace.SpanKindInternal))

	response := cryptographyv2alphapb.EncryptResponse{
		Result: &cryptographyv2alphapb.EncryptResponse_Err{Err: "error response"},
	}

	span.End()

	return connect.NewResponse(&response), nil
}

// Decrypt processes a request to decrypt the provided ciphertext and returns a response containing either plaintext or an error.
func (s *EncryptionServiceHandler) Decrypt(ctx context.Context, req *connect.Request[cryptographyv2alphapb.DecryptRequest]) (*connect.Response[cryptographyv2alphapb.DecryptResponse], error) {
	tracer := *opentelemetryv1.Bound.Tracer
	// log := *zaploggerv1.Bound.Logger

	fmt.Println(req)

	_, span := tracer.Start(ctx, "encrypt", trace.WithSpanKind(trace.SpanKindInternal))

	response := cryptographyv2alphapb.DecryptResponse{
		Result: &cryptographyv2alphapb.DecryptResponse_Err{Err: "error response"},
	}

	span.End()

	return connect.NewResponse(&response), nil
}
