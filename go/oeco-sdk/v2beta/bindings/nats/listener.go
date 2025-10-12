package natsnodev1

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	protopb "google.golang.org/protobuf/proto"

	"github.com/mennanov/fmutils"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
	optionv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SpecEventListener is an interface for handling event streaming, listening, and processing for specific configurations.
// It defines methods for retrieving listener configuration, listening to events, and processing event messages.
type SpecEventListener interface {
	Configure() *ListenerConfiguration
	Listen(ctx context.Context, listenerErr chan sdkv2betalib.SpecListenableErr)
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
	configuration := listener.Configure()

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
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered panic in listener.Process for %s: %v\n", subject, r)
			}
		}()

		messageCtx, message, _ := convertNatsToListenerMessage(configuration, msg)

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
func RespondToMultiplexedRequest(ctx context.Context, message *ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	js := *Bound.JetStream
	nm := *message.NatsMessage

	if message.Spec == nil {
		spec := specv2pb.Spec{SpecError: sdkv2betalib.ErrServerInternal.ToStatus()}
		respond(&nm, &spec)
		return
	}

	fields := receivedFields(message.Spec, nm.Subject)
	log.Info("Received multiplexed request", fields...)

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
		_, err = js.Publish(ctx, subject, specBytes)
		if err != nil {
			message.Spec.SpecError = sdkv2betalib.ErrServerInternal.WithSpecDetail(message.Spec).WithInternalErrorDetail(errors.New("found error when publishing to jetstream subject: "+subject), err).ToStatus()
			respond(&nm, message.Spec)
			return
		}
		log.Info("Published to jetstream subject: " + subject)
	}()

	respond(&nm, message.Spec)
}

// respond reply to a NATS message
func respond(msg *nats.Msg, spec *specv2pb.Spec) {
	log := *zaploggerv1.Bound.Logger
	if msg == nil {
		log.Warn("Received nil message, ignoring")
		return
	}

	if spec == nil {
		e := sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("the spec is nil: ")).ToStatus()
		spec = &specv2pb.Spec{
			SpecError: e,
		}
	}

	spec.CompletedAt = timestamppb.Now()

	fields := completedFields(spec, msg.Subject)
	var milliseconds int64
	if spec.CompletedAt != nil && spec.ReceivedAt != nil {
		milliseconds = spec.CompletedAt.AsTime().Sub(spec.ReceivedAt.AsTime()).Milliseconds()
	}
	log.Info(fmt.Sprintf("Completed multiplexed request in %d ms\n", milliseconds), fields...)

	marshal, err := protopb.Marshal(spec)
	if err != nil {
		log.Error(sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("cannot marshal spec: "), err).Error())
		err = msg.Respond(nil)
		if err != nil {
			log.Error(sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("error responding to NATS: "), err).Error())
		}
	}

	err = msg.Respond(marshal)
	if err != nil {
		log.Error(sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("error responding to NATS: "), err).Error())
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
