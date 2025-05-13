package temporalv1

//
//import (
//	"context"
//	"errors"
//	"fmt"
//
//	"github.com/mennanov/fmutils"
//	"go.uber.org/zap"
//	protopb "google.golang.org/protobuf/proto"
//
//	"github.com/nats-io/nats.go"
//	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
//	specproto "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
//	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
//	"go.temporal.io/sdk/worker"
//)
//
//// SpecEventListener is an interface for handling event streaming, listening, and processing for specific configurations.
//// It defines methods for retrieving listener configuration, listening to events, and processing event messages.
//type SpecEventListener interface {
//	GetConfiguration() *ListenerConfiguration
//	Listen(ctx context.Context, listenerErr chan sdkv2alphalib.SpecListenableErr)
//	Process(ctx context.Context, request *ListenerMessage)
//}
//
//// ListenerConfiguration defines the configuration for a listener, including stream type, entity, subject, queue, and Jetstream settings.
//type ListenerConfiguration struct {
//	Entity                sdkv2alphalib.Entity
//	Subject               string
//	Queue                 string
//	TemporalWorkerOptions *worker.Options
//}
//
//// ListenerMessage represents a message delivered to a consumer along with its associated metadata and configuration.
//type ListenerMessage struct {
//	SpecKey *specproto.SpecKey
//	Spec    *specproto.Spec
//
//	ListenerConfiguration *ListenerConfiguration
//}
//
//// ListenerErr represents an error encountered by a listener, including the related subscription for context.
//type ListenerErr struct {
//	Error error
//}
//
//// ListenForMultiplexedSpecEventsSync subscribes to a NATS subject to process multiplexed spec events synchronously.
//func ListenForMultiplexedSpecEventsSync(_ context.Context, listener SpecEventListener) {
//	configuration := listener.GetConfiguration()
//
//	if configuration == nil || configuration.Subject == "" || configuration.Queue == "" {
//		fmt.Println("Configuration is missing. Subject and Queue are required when configuring a SpecListener")
//		panic("Configuration is missing")
//	}
//}
//
//// convertNatsToListenerMessage transforms a NATS message into a ListenerMessage while setting up the required context.
//// It unmarshals data from the NATS message into a Spec object and attaches it to the ListenerMessage.
//// Returns a context, ListenerMessage populated with details from the input, and an error if unmarshalling fails.
////
////nolint:unparam
//func convertNatsToListenerMessage(config *ListenerConfiguration, msg *nats.Msg) (context.Context, ListenerMessage, error) {
//	ctx := context.Background()
//
//	s := &specproto.Spec{}
//	m := *msg
//	err := protopb.Unmarshal(m.Data, s)
//	if err != nil {
//		fmt.Println(sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not unmarshall spec")))
//	}
//
//	// ctx = interceptor.DecorateContextWithSpec(ctx, *s)
//
//	return ctx, ListenerMessage{
//		Spec:                  s,
//		Subscription:          nil,
//		NatsMessage:           &m,
//		ListenerConfiguration: config,
//	}, nil
//}
