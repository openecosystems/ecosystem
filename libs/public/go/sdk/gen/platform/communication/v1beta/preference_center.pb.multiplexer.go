// Code generated by protoc-gen-platform go/multiplexer. DO NOT EDIT.
// source: platform/communication/v1beta/preference_center.proto

package communicationv1betapb

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

// PreferenceCenterServiceHandler is the domain level implementation of the server API for mutations of the PreferenceCenterService service
type PreferenceCenterServiceHandler struct{}

func (s *PreferenceCenterServiceHandler) CreateOrUpdatePreference(ctx context.Context, req *connect.Request[CreateOrUpdatePreferenceRequest]) (*connect.Response[CreateOrUpdatePreferenceResponse], error) {

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
		if !spec.SpecData.FieldMask.IsValid(&CreateOrUpdatePreferenceResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := PreferenceCenterSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexCommandSync(handlerCtx, spec, &natsnodev1.SpecCommand{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		CommandName:    "",
		CommandTopic:   CommandDataPreferenceCenterTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd CreateOrUpdatePreferenceResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *PreferenceCenterServiceHandler) DeletePreference(ctx context.Context, req *connect.Request[DeletePreferenceRequest]) (*connect.Response[DeletePreferenceResponse], error) {

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
		if !spec.SpecData.FieldMask.IsValid(&DeletePreferenceResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := PreferenceCenterSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexCommandSync(handlerCtx, spec, &natsnodev1.SpecCommand{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		CommandName:    "",
		CommandTopic:   CommandDataPreferenceCenterTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd DeletePreferenceResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *PreferenceCenterServiceHandler) GetPreference(ctx context.Context, req *connect.Request[GetPreferenceRequest]) (*connect.Response[GetPreferenceResponse], error) {

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
		if !spec.SpecData.FieldMask.IsValid(&GetPreferenceResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := PreferenceCenterSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexEventSync(handlerCtx, spec, &natsnodev1.SpecEvent{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		EventName:      "",
		EventTopic:     EventDataPreferenceCenterTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd GetPreferenceResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}

func (s *PreferenceCenterServiceHandler) GetPreferenceOptions(ctx context.Context, req *connect.Request[GetPreferenceOptionsRequest]) (*connect.Response[GetPreferenceOptionsResponse], error) {

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
		if !spec.SpecData.FieldMask.IsValid(&GetPreferenceOptionsResponse{}) {
			log.Error("Invalid field mask")
			return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("Invalid field mask"))
		}
	}

	// Distributed Domain Handler
	handlerCtx, handlerSpan := tracer.Start(specCtx, "event-generation", trace.WithSpanKind(trace.SpanKindInternal))

	entity := PreferenceCenterSpecEntity{}
	reply, err2 := natsnodev1.Bound.MultiplexEventSync(handlerCtx, spec, &natsnodev1.SpecEvent{
		Request:        req.Msg,
		Stream:         natsnodev1.NewInboundStream(),
		EventName:      "",
		EventTopic:     EventDataPreferenceCenterTopic,
		EntityTypeName: entity.TypeName(),
	})
	if err2 != nil {
		log.Error(err2.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	var dd GetPreferenceOptionsResponse
	err3 := proto.Unmarshal(reply.Data, &dd)
	if err3 != nil {
		log.Error(err3.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	handlerSpan.End()

	return connect.NewResponse(&dd), nil

}
