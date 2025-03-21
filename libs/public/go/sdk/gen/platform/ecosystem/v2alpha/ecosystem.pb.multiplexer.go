// Code generated by protoc-gen-platform go/multiplexer. DO NOT EDIT.
// source: platform/ecosystem/v2alpha/ecosystem.proto

package ecosystemv2alphapb

import (
	"connectrpc.com/connect"
	"errors"
	"github.com/openecosystems/ecosystem/libs/partner/go/nats"
	"github.com/openecosystems/ecosystem/libs/partner/go/opentelemetry"
	"github.com/openecosystems/ecosystem/libs/partner/go/protovalidate"
	"github.com/openecosystems/ecosystem/libs/partner/go/zap"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	"github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"

	_ "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"context"
)

// EcosystemServiceHandler is the domain level implementation of the server API for mutations of the EcosystemService service
type EcosystemServiceHandler struct{}

func (s *EcosystemServiceHandler) CreateEcosystem(ctx context.Context, req *connect.Request[CreateEcosystemRequest]) (*connect.Response[CreateEcosystemResponse], error) {

	tracer := *opentelemetryv1.Bound.Tracer
	log := *zaploggerv1.Bound.Logger

	// Executes top level validation, no business domain validation
	validationCtx, validationSpan := tracer.Start(ctx, "request-validation", trace.WithSpanKind(trace.SpanKindInternal))
	v := *protovalidatev0.Bound.Validator
	if err := v.Validate(req.Msg); err != nil {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(err)
	}
	validationSpan.End()

	// Spec Propagation
	specCtx, specSpan := tracer.Start(validationCtx, "spec-propagation", trace.WithSpanKind(trace.SpanKindInternal))
	spec, ok := ctx.Value(sdkv2alphalib.SpecContextKey).(*specv2pb.Spec)
	if !ok {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("Cannot propagate spec to context"))
	}
	specSpan.End()

	// Validate field mask
	if spec.SpecData.FieldMask != nil && len(spec.SpecData.FieldMask.Paths) > 0 {
		spec.SpecData.FieldMask.Normalize()
		if !spec.SpecData.FieldMask.IsValid(&CreateEcosystemResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := EcosystemSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexCommandSync(handlerCtx, spec, &natsnodev1.SpecCommand{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		CommandName:    "",
		CommandTopic:   CommandDataEcosystemTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd CreateEcosystemResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *EcosystemServiceHandler) UpdateEcosystem(ctx context.Context, req *connect.Request[UpdateEcosystemRequest]) (*connect.Response[UpdateEcosystemResponse], error) {

	tracer := *opentelemetryv1.Bound.Tracer
	log := *zaploggerv1.Bound.Logger

	// Executes top level validation, no business domain validation
	validationCtx, validationSpan := tracer.Start(ctx, "request-validation", trace.WithSpanKind(trace.SpanKindInternal))
	v := *protovalidatev0.Bound.Validator
	if err := v.Validate(req.Msg); err != nil {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(err)
	}
	validationSpan.End()

	// Spec Propagation
	specCtx, specSpan := tracer.Start(validationCtx, "spec-propagation", trace.WithSpanKind(trace.SpanKindInternal))
	spec, ok := ctx.Value(sdkv2alphalib.SpecContextKey).(*specv2pb.Spec)
	if !ok {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("Cannot propagate spec to context"))
	}
	specSpan.End()

	// Validate field mask
	if spec.SpecData.FieldMask != nil && len(spec.SpecData.FieldMask.Paths) > 0 {
		spec.SpecData.FieldMask.Normalize()
		if !spec.SpecData.FieldMask.IsValid(&UpdateEcosystemResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := EcosystemSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexCommandSync(handlerCtx, spec, &natsnodev1.SpecCommand{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		CommandName:    "",
		CommandTopic:   CommandDataEcosystemTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd UpdateEcosystemResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *EcosystemServiceHandler) DeleteEcosystem(ctx context.Context, req *connect.Request[DeleteEcosystemRequest]) (*connect.Response[DeleteEcosystemResponse], error) {

	tracer := *opentelemetryv1.Bound.Tracer
	log := *zaploggerv1.Bound.Logger

	// Executes top level validation, no business domain validation
	validationCtx, validationSpan := tracer.Start(ctx, "request-validation", trace.WithSpanKind(trace.SpanKindInternal))
	v := *protovalidatev0.Bound.Validator
	if err := v.Validate(req.Msg); err != nil {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(err)
	}
	validationSpan.End()

	// Spec Propagation
	specCtx, specSpan := tracer.Start(validationCtx, "spec-propagation", trace.WithSpanKind(trace.SpanKindInternal))
	spec, ok := ctx.Value(sdkv2alphalib.SpecContextKey).(*specv2pb.Spec)
	if !ok {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("Cannot propagate spec to context"))
	}
	specSpan.End()

	// Validate field mask
	if spec.SpecData.FieldMask != nil && len(spec.SpecData.FieldMask.Paths) > 0 {
		spec.SpecData.FieldMask.Normalize()
		if !spec.SpecData.FieldMask.IsValid(&DeleteEcosystemResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := EcosystemSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexCommandSync(handlerCtx, spec, &natsnodev1.SpecCommand{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		CommandName:    "",
		CommandTopic:   CommandDataEcosystemTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd DeleteEcosystemResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *EcosystemServiceHandler) ListEcosystems(ctx context.Context, req *connect.Request[ListEcosystemsRequest]) (*connect.Response[ListEcosystemsResponse], error) {

	tracer := *opentelemetryv1.Bound.Tracer
	log := *zaploggerv1.Bound.Logger

	// Executes top level validation, no business domain validation
	validationCtx, validationSpan := tracer.Start(ctx, "request-validation", trace.WithSpanKind(trace.SpanKindInternal))
	v := *protovalidatev0.Bound.Validator
	if err := v.Validate(req.Msg); err != nil {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(err)
	}
	validationSpan.End()

	// Spec Propagation
	specCtx, specSpan := tracer.Start(validationCtx, "spec-propagation", trace.WithSpanKind(trace.SpanKindInternal))
	spec, ok := ctx.Value(sdkv2alphalib.SpecContextKey).(*specv2pb.Spec)
	if !ok {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("Cannot propagate spec to context"))
	}
	specSpan.End()

	// Validate field mask
	if spec.SpecData.FieldMask != nil && len(spec.SpecData.FieldMask.Paths) > 0 {
		spec.SpecData.FieldMask.Normalize()
		if !spec.SpecData.FieldMask.IsValid(&ListEcosystemsResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := EcosystemSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexEventSync(handlerCtx, spec, &natsnodev1.SpecEvent{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		EventName:      "",
		EventTopic:     EventDataEcosystemTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd ListEcosystemsResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *EcosystemServiceHandler) GetEcosystem(ctx context.Context, req *connect.Request[GetEcosystemRequest]) (*connect.Response[GetEcosystemResponse], error) {

	tracer := *opentelemetryv1.Bound.Tracer
	log := *zaploggerv1.Bound.Logger

	// Executes top level validation, no business domain validation
	validationCtx, validationSpan := tracer.Start(ctx, "request-validation", trace.WithSpanKind(trace.SpanKindInternal))
	v := *protovalidatev0.Bound.Validator
	if err := v.Validate(req.Msg); err != nil {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(err)
	}
	validationSpan.End()

	// Spec Propagation
	specCtx, specSpan := tracer.Start(validationCtx, "spec-propagation", trace.WithSpanKind(trace.SpanKindInternal))
	spec, ok := ctx.Value(sdkv2alphalib.SpecContextKey).(*specv2pb.Spec)
	if !ok {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("Cannot propagate spec to context"))
	}
	specSpan.End()

	// Validate field mask
	if spec.SpecData.FieldMask != nil && len(spec.SpecData.FieldMask.Paths) > 0 {
		spec.SpecData.FieldMask.Normalize()
		if !spec.SpecData.FieldMask.IsValid(&GetEcosystemResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := EcosystemSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexEventSync(handlerCtx, spec, &natsnodev1.SpecEvent{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		EventName:      "",
		EventTopic:     EventDataEcosystemTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd GetEcosystemResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}
