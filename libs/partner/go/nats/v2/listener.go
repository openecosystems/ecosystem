package natsnodev2

import (
	"context"
	"errors"
	"fmt"

	zaploggerv1 "libs/partner/go/zap/v1"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"github.com/mennanov/fmutils"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	protopb "google.golang.org/protobuf/proto"

	specproto "libs/protobuf/go/protobuf/gen/platform/spec/v2"
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
	Subject                string
	Queue                  string
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
}

// ListenerErr represents an error encountered by a listener, including the related subscription for context.
type ListenerErr struct {
	Error        error
	Subscription *nats.Subscription
}

// ListenForMultiplexedSpecEventsSync subscribes to a NATS subject to process multiplexed spec events synchronously.
func ListenForMultiplexedSpecEventsSync(_ context.Context, listener SpecEventListener) {
	configuration := listener.GetConfiguration()

	if configuration == nil || configuration.Subject == "" || configuration.Queue == "" {
		fmt.Println("Configuration is missing. Subject and Queue are required when configuring a SpecListener")
		panic("Configuration is missing")
	}

	n := Bound.Nats
	_, err := n.QueueSubscribe(configuration.Subject, configuration.Queue, func(msg *nats.Msg) {
		messageCtx, message, _ := convertNatsToListenerMessage(configuration, msg)
		listener.Process(messageCtx, &message)
	})
	if err != nil {
		fmt.Println("Found error in queue subscribing to nats subject: " + err.Error())
		panic("Cannot start queue subscriber")
	}
}

// RespondToSyncCommand processes an inbound request, modifies the provided message, and sends a response through NATS.
func RespondToSyncCommand(_ context.Context, request *ListenerMessage, m protopb.Message) {
	log := *zaploggerv1.Bound.Logger
	js := *Bound.JetStream
	nm := *request.NatsMessage

	fmutils.Filter(m, request.Spec.SpecData.FieldMask.GetPaths())

	specBytes, err := protopb.Marshal(request.Spec)
	if err != nil {
		return
	}

	configuration := request.ListenerConfiguration

	subject := GetSubjectName(configuration.StreamType.StreamPrefix(), configuration.Entity.CommandTopic())

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

// ContinuouslyListenForEvents initializes and continuously listens for events from a JetStream consumer in a concurrent manner.
// It processes incoming messages by converting them to a listener-compatible format and invokes the listener's process method.
// The function creates a specified number of workers to handle messages concurrently, using a semaphore to control concurrency.
// Requires a valid listener configuration with defined StreamType and Entity for initialization.
// On configuration errors or message processing errors, the function will panic or log the error respectively.
func ContinuouslyListenForEvents(ctx context.Context, rootConfig *sdkv2alphalib.Configuration, listener SpecEventListener, _ chan sdkv2alphalib.SpecListenableErr) {
	configuration := listener.GetConfiguration()

	if configuration == nil || configuration.StreamType == nil || configuration.Entity == nil {
		fmt.Println("Configuration is missing. StreamType and Entity are required when configuring a SpecListener")
		panic("Configuration is missing")
	}

	js := *Bound.JetStream

	stream := GetStreamName(rootConfig.App.EnvironmentName, configuration.StreamType.StreamPrefix(), configuration.Entity.TypeName())

	cons, _ := js.CreateOrUpdateConsumer(ctx, stream, *configuration.JetstreamConfiguration)

	// PullMaxMessages determines how many messages will be sent to the client in a single pull request
	iter, _ := cons.Messages(jetstream.PullMaxMessages(1))
	numWorkers := 5
	sem := make(chan struct{}, numWorkers)
	for {
		sem <- struct{}{}
		go func() {
			defer func() {
				<-sem
			}()
			msg, err := iter.Next()
			if err != nil {
				fmt.Println(err)
				// handle err
			}

			messageCtx, message, err2 := convertToListenerMessage(configuration, &msg)
			if err2 != nil {
				fmt.Println(err2)
				// handle err
			}

			listener.Process(messageCtx, &message)
		}()
	}
}

// convertToListenerMessage transforms a JetStream message into a ListenerMessage, along with its associated context.
// It unmarshals the message into a Spec and decorates the context with metadata.
// Returns the context, the constructed ListenerMessage, and any error encountered during processing.
//
//nolint:unparam
func convertToListenerMessage(config *ListenerConfiguration, msg *jetstream.Msg) (context.Context, ListenerMessage, error) {
	ctx := context.Background()

	s := &specproto.Spec{}
	m := *msg
	err := protopb.Unmarshal(m.Data(), s)
	if err != nil {
		fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec")))
	}

	// ctx = interceptor.DecorateContextWithSpec(ctx, *s)

	return ctx, ListenerMessage{
		Spec:                  s,
		Subscription:          nil,
		Message:               &m,
		ListenerConfiguration: config,
	}, nil
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

	return ctx, ListenerMessage{
		Spec:                  s,
		Subscription:          nil,
		NatsMessage:           &m,
		ListenerConfiguration: config,
	}, nil
}
