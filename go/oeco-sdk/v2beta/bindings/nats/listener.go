package natsnodev1

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	protopb "google.golang.org/protobuf/proto"

	apexlog "github.com/apex/log"
	"github.com/mennanov/fmutils"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
	optionv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/anypb"
)

// SpecEventListener is an interface for handling event streaming, listening, and processing for specific configurations.
// It defines methods for retrieving listener configuration, listening to events, and processing event messages.
type SpecEventListener interface {
	GetConfiguration() *ListenerConfiguration
	Listen(ctx context.Context, listenerErr chan sdkv2betalib.SpecListenableErr)
	Validate(ctx context.Context, message *ListenerMessage)
	Process(ctx context.Context, message *ListenerMessage)
}

// ListenerConfiguration defines the configuration for a listener, including stream type, entity, subject, queue, and Jetstream settings.
type ListenerConfiguration struct {
	StreamType             Stream
	Entity                 sdkv2betalib.Entity
	TypeName               string
	Procedure              string
	Topic                  string
	CQRS                   optionv2pb.CQRSType
	JetstreamConfiguration *jetstream.ConsumerConfig
}

// ListenerMessage represents a message delivered to a consumer along with its associated metadata and configuration.
type ListenerMessage struct {
	SpecKey               *specv2pb.SpecKey
	Spec                  *specv2pb.Spec
	Subscription          *jetstream.Consumer
	Message               *jetstream.Msg
	NatsMessage           *nats.Msg
	ListenerConfiguration *ListenerConfiguration
	EventResponseChannel  string
	Request               protopb.Message
	Response              protopb.Message
	SpecError             sdkv2betalib.SpecErrorable
}

// ListenerErr represents an error encountered by a listener, including the related subscription for context.
type ListenerErr struct {
	Error        error
	Subscription *nats.Subscription
}

// ListenForMultiplexedRequests subscribes to a NATS subject to process multiplexed spec events synchronously.
func ListenForMultiplexedRequests(_ context.Context, listener SpecEventListener) {
	configuration := listener.GetConfiguration()

	if configuration == nil || configuration.Procedure == "" || configuration.StreamType == nil || configuration.Entity == nil {
		fmt.Println("Configuration is missing. Entity, Procedure, StreamType are required when configuring a SpecListener")
		panic("Configuration is missing")
	}

	n := Bound.Nats
	subject := ""

	switch configuration.CQRS {
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_CREATE:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_UPDATE:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_CLIENT_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_SERVER_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_BIDI_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_DELETE:
		subject = GetMultiplexedRequestSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.CommandTopic(), configuration.Procedure)
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_LIST:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_EXISTS:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_CLIENT_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_BIDI_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_GET:
		subject = GetMultiplexedRequestSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.EventTopic(), configuration.Procedure)
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_SERVER_STREAM:
		// subject = GetMultiplexedRequestSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.EventTopic(), configuration.Procedure+".>")
		subject = GetMultiplexedRequestSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.EventTopic(), configuration.Procedure)
	case optionv2pb.CQRSType_CQRS_TYPE_NONE:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_UNSPECIFIED:
		fallthrough
	default:
		panic("Cannot start queue subscriber without a proper CQRS type")
	}

	queue := GetQueueGroupName(configuration.StreamType.StreamPrefix(), configuration.Entity.TypeName(), configuration.Procedure)

	fmt.Println("Listening for multiplexed spec events on subject: " + subject)

	_, err := n.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		messageCtx, message, _ := convertNatsToListenerMessage(configuration, msg)

		listener.Validate(messageCtx, message)

		if message.Spec.Context.Validation.GetValidateOnly() {
			RespondToListenerProcess(messageCtx, message)
			return
		}

		if message.SpecError != nil || message.Response != nil {
			RespondToListenerProcess(messageCtx, message)
			return
		}

		// The Processor is responsible for replying to the Reply subject and responding with any errors
		listener.Process(messageCtx, message)
		RespondToListenerProcess(messageCtx, message)
	})
	if err != nil {
		fmt.Println("Found error in queue subscribing to nats subject: " + err.Error())
		panic("Cannot start queue subscriber")
	}
}

// RespondToMultiplexedRequest processes an inbound message, modifies the provided message, and sends a response through NATS.
func RespondToMultiplexedRequest(_ context.Context, message *ListenerMessage) {
	js := *Bound.JetStream
	nm := *message.NatsMessage

	if message.Spec == nil {
		spec := specv2pb.Spec{SpecError: sdkv2betalib.ErrServerInternal.ToStatus()}
		respond(&nm, &spec)
		return
	}

	if message.Spec.SpecData != nil && message.Spec.SpecData.Data != nil {
		fmutils.Filter(message.Spec.SpecData.Data, message.Spec.SpecData.FieldMask.GetPaths())
	}

	specBytes, err := protopb.Marshal(message.Spec)
	if err != nil {
		message.Spec.SpecError = sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(err).ToStatus()
		respond(&nm, message.Spec)
		return
	}

	configuration := message.ListenerConfiguration

	subject := ""

	switch configuration.CQRS {
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_CREATE:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_UPDATE:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_CLIENT_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_SERVER_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_BIDI_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_MUTATION_DELETE:
		subject = GetSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.CommandTopic(), configuration.Procedure)
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_LIST:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_EXISTS:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_CLIENT_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_SERVER_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_BIDI_STREAM:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_QUERY_GET:
		subject = GetSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.EventTopic(), configuration.Procedure)
	case optionv2pb.CQRSType_CQRS_TYPE_NONE:
		fallthrough
	case optionv2pb.CQRSType_CQRS_TYPE_UNSPECIFIED:
		fallthrough
	default:
		message.Spec.SpecError = sdkv2betalib.ErrServerInternal.WithSpecDetail(message.Spec).WithInternalErrorDetail(errors.New("cannot respond to multiplexed requests, as CQRS type is invalid. This should have been caught at startup. Bad")).ToStatus()
		respond(&nm, message.Spec)
		return
	}

	go func() {
		_, err = js.Publish(context.Background(), subject, specBytes)
		if err != nil {
			message.Spec.SpecError = sdkv2betalib.ErrServerInternal.WithSpecDetail(message.Spec).WithInternalErrorDetail(errors.New("found error when publishing"), err).ToStatus()
			respond(&nm, message.Spec)
			return
		}
	}()

	respond(&nm, message.Spec)
}

// respond reply to a NATS message
func respond(msg *nats.Msg, spec *specv2pb.Spec) {
	if msg == nil {
		apexlog.Warn("Received nil message, ignoring")
		return
	}

	if spec == nil {
		e := sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("the spec is nil: ")).ToStatus()
		spec = &specv2pb.Spec{
			SpecError: e,
		}
	}

	marshal, err := protopb.Marshal(spec)
	if err != nil {
		apexlog.Error(sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("cannot marshal spec: "), err).Error())
		err = msg.Respond(nil)
		if err != nil {
			apexlog.Error(sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("error responding to NATS: "), err).Error())
		}
	}

	err = msg.Respond(marshal)
	if err != nil {
		apexlog.Error(sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("error responding to NATS: "), err).Error())
	}
}

// ListenForJetStreamEvents subscribes to a Jetstream subject.
func ListenForJetStreamEvents(ctx context.Context, env string, listener SpecEventListener) {
	configuration := listener.GetConfiguration()

	if configuration == nil || configuration.Procedure == "" || configuration.StreamType == nil {
		apexlog.Error("Configuration is missing. Procedure, StreamType are required when configuring a SpecListener")
		panic("Configuration is missing")
	}

	log := *zaploggerv1.Bound.Logger
	js := *Bound.JetStream
	streamName := GetStreamName(env, configuration.StreamType.StreamPrefix(), configuration.TypeName)

	stream, err := js.Stream(ctx, streamName)
	if err != nil {
		apexlog.Error("Could not find stream: " + streamName + err.Error())
		return
	}

	c, err := stream.CreateOrUpdateConsumer(ctx, *configuration.JetstreamConfiguration)
	if err != nil {
		log.Error("SpecError creating consumer", zap.Error(err))
		panic("Cannot start consumer")
	}

	// TODO: This must be closed
	_, err = c.Consume(func(msg jetstream.Msg) {
		messageCtx, message, _ := convertJetstreamToListenerMessage(configuration, &msg)

		listener.Validate(messageCtx, message)
		if message.Response != nil {
			RespondToJetstreamEvent(messageCtx, message)
			return
		}

		// The Processor is responsible for replying to the Reply subject and responding with any errors
		listener.Process(messageCtx, message)
		RespondToJetstreamEvent(messageCtx, message)
	}, jetstream.ConsumeErrHandler(func(_ jetstream.ConsumeContext, err error) {
		fmt.Println(err)
	}))
	if err != nil {
		log.Fatal("consume error", zap.Error(err))
	}

	apexlog.Info("Listening for stream spec events on subject: " + streamName)

	//
}

// RespondToJetstreamEvent processes an inbound message, modifies the provided message, and sends a response through NATS.
func RespondToJetstreamEvent(_ context.Context, message *ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	jm := *message.Message

	err4 := jm.Ack()
	if err4 != nil {
		log.Error("SpecError acknowledging message", zap.Error(err4))
		return
	}
}

func RespondToListenerProcess(ctx context.Context, message *ListenerMessage) {
	serr := message.SpecError
	response := message.Response

	if serr != nil {
		message.Spec.SpecError = serr.ToStatus()
		RespondToMultiplexedRequest(ctx, message)
		return
	}

	if response == nil {
		message.Spec.SpecError = sdkv2betalib.ErrServerInternal.WithSpecDetail(message.Spec).WithInternalErrorDetail(errors.New("response is nil: ")).ToStatus()
	}

	a, err := anypb.New(response)
	if err != nil {
		message.Spec.SpecError = sdkv2betalib.ErrServerInternal.WithSpecDetail(message.Spec).WithInternalErrorDetail(err).ToStatus()
		RespondToMultiplexedRequest(ctx, message)
		return
	}

	if message.Spec.SpecData != nil {
		message.Spec.SpecData = &specv2pb.SpecData{}
	}

	message.Spec.SpecData.Data = a
	RespondToMultiplexedRequest(ctx, message)
}

// convertNatsToListenerMessage transforms a NATS message into a ListenerMessage while setting up the required context.
// It unmarshals data from the NATS message into a Spec object and attaches it to the ListenerMessage.
// Returns a context, ListenerMessage populated with details from the input, and an error if unmarshalling fails.
//
//nolint:unparam
func convertNatsToListenerMessage(config *ListenerConfiguration, msg *nats.Msg) (context.Context, *ListenerMessage, sdkv2betalib.SpecErrorable) {
	// Start with a new ctx here because it must remain transaction safe
	ctx := context.Background()
	s := &specv2pb.Spec{}
	m := *msg
	err := protopb.Unmarshal(m.Data, s)
	if err != nil {
		return ctx, nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec"), err)
	}

	parentSpanCtx := convertSpecSpanContextToContext(s)
	ctx = trace.ContextWithRemoteSpanContext(ctx, parentSpanCtx)

	responseSubject := GetStreamResponseSubjectName(
		config.StreamType.StreamPrefix(),
		config.Topic,
		config.Procedure,
		s.MessageId,
	)

	return ctx, &ListenerMessage{
		Spec:                  s,
		Subscription:          nil,
		NatsMessage:           &m,
		ListenerConfiguration: config,
		EventResponseChannel:  responseSubject,
	}, nil
}

func convertJetstreamToListenerMessage(config *ListenerConfiguration, msg *jetstream.Msg) (context.Context, *ListenerMessage, sdkv2betalib.SpecErrorable) {
	// Start with a new ctx here because it must remain transaction safe
	ctx := context.Background()
	s := &specv2pb.Spec{}
	m := *msg
	err := protopb.Unmarshal(m.Data(), s)
	if err != nil {
		return ctx, nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec"), err)
	}

	parentSpanCtx := convertSpecSpanContextToContext(s)
	ctx = trace.ContextWithRemoteSpanContext(ctx, parentSpanCtx)

	return ctx, &ListenerMessage{
		Spec:                  s,
		Subscription:          nil,
		Message:               &m,
		ListenerConfiguration: config,
	}, nil
}

func convertSpecSpanContextToContext(spec *specv2pb.Spec) trace.SpanContext {
	if spec == nil || spec.SpanContext == nil {
		return trace.SpanContext{}
	}

	traceID, err := trace.TraceIDFromHex(spec.SpanContext.TraceId)
	if err != nil {
		return trace.SpanContext{}
	}

	return trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceID,
		Remote:  true,
	})
}
