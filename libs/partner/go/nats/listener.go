package natsnodev1

import (
	"context"
	"errors"
	"fmt"

	"github.com/mennanov/fmutils"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	protopb "google.golang.org/protobuf/proto"

	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	optionv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
	specproto "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// SpecEventListener is an interface for handling event streaming, listening, and processing for specific configurations.
// It defines methods for retrieving listener configuration, listening to events, and processing event messages.
type SpecEventListener interface {
	GetConfiguration() *ListenerConfiguration
	Listen(ctx context.Context, listenerErr chan sdkv2alphalib.SpecListenableErr)
	Process(ctx context.Context, request *ListenerMessage)
}

// ListenerConfiguration defines the configuration for a listener, including stream type, entity, subject, queue, and Jetstream settings.
type ListenerConfiguration struct {
	StreamType             Stream
	Entity                 sdkv2alphalib.Entity
	TypeName               string
	Procedure              string
	Topic                  string
	CQRS                   optionv2pb.CQRSType
	JetstreamConfiguration *jetstream.ConsumerConfig
}

// ListenerMessage represents a message delivered to a consumer along with its associated metadata and configuration.
type ListenerMessage struct {
	SpecKey               *specproto.SpecKey
	Spec                  *specproto.Spec
	Subscription          *jetstream.Consumer
	Message               *jetstream.Msg
	NatsMessage           *nats.Msg
	ListenerConfiguration *ListenerConfiguration
	EventResponseChannel  string
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
		listener.Process(messageCtx, &message)
	})
	if err != nil {
		fmt.Println("Found error in queue subscribing to nats subject: " + err.Error())
		panic("Cannot start queue subscriber")
	}
}

// RespondToMultiplexedRequest processes an inbound request, modifies the provided message, and sends a response through NATS.
func RespondToMultiplexedRequest(_ context.Context, request *ListenerMessage, m protopb.Message) {
	log := *zaploggerv1.Bound.Logger
	js := *Bound.JetStream
	nm := *request.NatsMessage

	fmutils.Filter(m, request.Spec.SpecData.FieldMask.GetPaths())

	specBytes, err := protopb.Marshal(request.Spec)
	if err != nil {
		return
	}

	configuration := request.ListenerConfiguration

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
		log.Error("Cannot respond to multiplexed requests, as CQRS type is invalid. This should have been caught at startup. Bad.")
	}

	go func() {
		_, err2 := js.Publish(context.Background(), subject, specBytes)
		if err2 != nil {
			log.Error("Found error when publishing", zap.Error(err2))
		}
	}()

	marshal, err3 := protopb.Marshal(m)
	if err3 != nil {
		fmt.Println("Cannot marshal Message", zap.Error(err3))
		return
	}

	err4 := nm.Respond(marshal)
	if err4 != nil {
		log.Error("Error acknowledging message", zap.Error(err4))
		return
	}
}

// ListenForJetStreamEvents subscribes to a Jetstream subject.
func ListenForJetStreamEvents(ctx context.Context, env string, listener SpecEventListener) {
	configuration := listener.GetConfiguration()

	if configuration == nil || configuration.Procedure == "" || configuration.StreamType == nil {
		fmt.Println("Configuration is missing. Procedure, StreamType are required when configuring a SpecListener")
		panic("Configuration is missing")
	}

	log := *zaploggerv1.Bound.Logger
	js := *Bound.JetStream
	streamName := GetStreamName(env, configuration.StreamType.StreamPrefix(), configuration.TypeName)

	stream, err := js.Stream(ctx, streamName)
	if err != nil {
		fmt.Println("Could not find stream: "+streamName, err)
		return
	}

	c, err := stream.CreateOrUpdateConsumer(ctx, *configuration.JetstreamConfiguration)
	if err != nil {
		log.Error("Error creating consumer", zap.Error(err))
		panic("Cannot start consumer")
	}

	// TODO: This must be closed
	_, err = c.Consume(func(msg jetstream.Msg) {
		messageCtx, message, _ := convertJetstreamToListenerMessage(configuration, &msg)
		listener.Process(messageCtx, &message)
	}, jetstream.ConsumeErrHandler(func(_ jetstream.ConsumeContext, err error) {
		fmt.Println(err)
	}))
	if err != nil {
		log.Fatal("consume error", zap.Error(err))
	}

	fmt.Println("Listening for stream spec events on subject: " + streamName)

	//
}

// RespondToJetstreamEvent processes an inbound request, modifies the provided message, and sends a response through NATS.
func RespondToJetstreamEvent(_ context.Context, request *ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	jm := *request.Message

	err4 := jm.Ack()
	if err4 != nil {
		log.Error("Error acknowledging message", zap.Error(err4))
		return
	}
}

// convertNatsToListenerMessage transforms a NATS message into a ListenerMessage while setting up the required context.
// It unmarshals data from the NATS message into a Spec object and attaches it to the ListenerMessage.
// Returns a context, ListenerMessage populated with details from the input, and an error if unmarshalling fails.
//
//nolint:unparam
func convertNatsToListenerMessage(config *ListenerConfiguration, msg *nats.Msg) (context.Context, ListenerMessage, error) {
	ctx := context.Background()

	s := &specproto.Spec{}
	m := *msg
	err := protopb.Unmarshal(m.Data, s)
	if err != nil {
		fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec")))
	}

	// ctx = interceptor.DecorateContextWithSpec(ctx, *s)

	responseSubject := GetStreamResponseSubjectName(
		config.StreamType.StreamPrefix(),
		config.Topic,
		config.Procedure,
		s.MessageId,
	)

	return ctx, ListenerMessage{
		Spec:                  s,
		Subscription:          nil,
		NatsMessage:           &m,
		ListenerConfiguration: config,
		EventResponseChannel:  responseSubject,
	}, nil
}

func convertJetstreamToListenerMessage(config *ListenerConfiguration, msg *jetstream.Msg) (context.Context, ListenerMessage, error) {
	ctx := context.Background()

	s := &specproto.Spec{}
	m := *msg
	err := protopb.Unmarshal(m.Data(), s)
	if err != nil {
		fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec")))
		return ctx, ListenerMessage{}, err
	}

	// ctx = interceptor.DecorateContextWithSpec(ctx, *s)

	return ctx, ListenerMessage{
		Spec:                  s,
		Subscription:          nil,
		Message:               &m,
		ListenerConfiguration: config,
	}, nil
}
