package natsnodev2

import (
	"context"
	"errors"
	"fmt"

	"libs/partner/go/zap/v1"
	"libs/public/go/sdk/v2alpha"

	zaploggerv1 "libs/partner/go/zap/v1"
	specproto "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"github.com/mennanov/fmutils"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	protopb "google.golang.org/protobuf/proto"

	specproto "libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

type SpecEventListener interface {
	GetConfiguration() *ListenerConfiguration
	Listen(ctx context.Context, listenerErr chan sdkv2alphalib.SpecListenableErr)
	Process(ctx context.Context, request *ListenerMessage)
}

type ListenerConfiguration struct {
	StreamType             Stream
	Entity                 sdkv2alphalib.Entity
	Subject                string
	Queue                  string
	JetstreamConfiguration *jetstream.ConsumerConfig
}

type ListenerMessage struct {
	SpecKey               *specproto.SpecKey
	Spec                  *specproto.Spec
	Subscription          *jetstream.Consumer
	Message               *jetstream.Msg
	NatsMessage           *nats.Msg
	ListenerConfiguration *ListenerConfiguration
}

type ListenerErr struct {
	Error        error
	Subscription *nats.Subscription
}

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

func ContinuouslyListenForEvents(ctx context.Context, rootConfig *sdkv2alphalib.Configuration, listener SpecEventListener, listenerErr chan sdkv2alphalib.SpecListenableErr) {
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

			messageCtx, message, err := convertToListenerMessage(configuration, &msg)

			listener.Process(messageCtx, &message)
		}()
	}
}

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
