package natsnodev1

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
	optionv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	protopb "google.golang.org/protobuf/proto"
)

type SpecEventBatchListener interface {
	Configure() *ListenerBatchConfiguration
	Listen(ctx context.Context, listenerErr chan sdkv2betalib.SpecListenableErr)
	BatchProcess(ctx context.Context, messages []*ListenerBatchMessage)
}

// ListenerBatchConfiguration defines the configuration for a listener, including stream type, entity, subject, queue, and Jetstream settings.
type ListenerBatchConfiguration struct {
	StreamType             Stream
	Entity                 sdkv2betalib.Entity
	BatchSize              int
	TypeName               string
	Procedure              string
	Topic                  string
	CQRS                   optionv2pb.CQRSType
	JetstreamConfiguration *jetstream.ConsumerConfig
}

type ListenerBatchMessage struct {
	Context                    context.Context
	SpecKey                    *specv2pb.SpecKey
	Spec                       *specv2pb.Spec
	Subscription               *jetstream.Consumer
	Message                    *jetstream.Msg
	NatsMessage                *nats.Msg
	ListenerBatchConfiguration *ListenerBatchConfiguration
	EventResponseChannel       string
	Request                    protopb.Message
	Response                   protopb.Message
	SpecError                  sdkv2betalib.SpecErrorable
}

// ListenForJetStreamEvents subscribes to a Jetstream subject.
func ListenForJetStreamEvents(ctx context.Context, env string, listener SpecEventBatchListener) {
	log := *zaploggerv1.Bound.Logger
	configuration := listener.Configure()

	if configuration == nil || configuration.Entity == nil || configuration.StreamType == nil {
		log.Error("Configuration is missing. Entity, StreamType are required when configuring a SpecListener")
		panic("Configuration is missing")
	}

	js := *Bound.JetStream
	streamName := GetStreamName(env, configuration.StreamType.StreamPrefix(), configuration.Entity.TypeName())

	stream, err := js.Stream(ctx, streamName)
	if err != nil {
		log.Error("Could not find stream: " + streamName + "; " + err.Error())
		return
	}

	c, err := stream.CreateOrUpdateConsumer(ctx, *configuration.JetstreamConfiguration)
	if err != nil {
		log.Error("SpecError creating consumer", zap.Error(err))
		panic("Cannot start consumer")
	}

	log.Info("Listening for stream spec events on subject: " + streamName)

	for {
		messages, err := c.Fetch(configuration.BatchSize)
		if err != nil && !errors.Is(err, nats.ErrTimeout) {
			log.Error("Fetch error:", zap.Error(err))
			continue
		}

		var batch []*ListenerBatchMessage
		for msg := range messages.Messages() {
			message, _ := convertJetstreamToListenerMessage(configuration, &msg)
			batch = append(batch, message)
		}

		listener.BatchProcess(ctx, batch)
	}

	//
	//_, err = c.Fetch(configuration.BatchSize, func(msg jetstream.Msg) {
	//	messageCtx, message, _ := convertJetstreamToListenerMessage(configuration, &msg)
	//
	//	// The Processor is responsible for replying to the Reply subject and responding with any errors
	//	listener.BatchProcess(messageCtx, message)
	//	// RespondToJetstreamEvent(messageCtx, message)
	//}, jetstream.ConsumeErrHandler(func(_ jetstream.ConsumeContext, err error) {
	//	fmt.Println(err)
	//}))
	//if err != nil {
	//	log.Fatal("consume error", zap.Error(err))
	//}

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

func convertJetstreamToListenerMessage(config *ListenerBatchConfiguration, msg *jetstream.Msg) (*ListenerBatchMessage, sdkv2betalib.SpecErrorable) {
	// Start with a new ctx here because it must remain transaction safe
	ctx := context.Background()
	s := &specv2pb.Spec{}
	m := *msg
	err := protopb.Unmarshal(m.Data(), s)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec"), err)
	}

	parentSpanCtx := convertSpecSpanContextToContext(s)
	ctx = trace.ContextWithRemoteSpanContext(ctx, parentSpanCtx)

	return &ListenerBatchMessage{
		Context:                    ctx,
		Spec:                       s,
		Subscription:               nil,
		Message:                    &m,
		ListenerBatchConfiguration: config,
	}, nil
}
