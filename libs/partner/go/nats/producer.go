package natsnodev1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// MultiplexCommandSync sends a command synchronously by publishing it to a NATS stream and awaiting a reply.
// Uses Nats Publish and Subscribe Pattern
func (b *Binding) MultiplexCommandSync(_ context.Context, s *specv2pb.Spec, command *SpecCommand) (*nats.Msg, error) {
	if command == nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecCommand object is required"))
	}

	log := *zaploggerv1.Bound.Logger
	// acc := *configurationv2alphalib.Bound.AdaptiveConfigurationControl

	s.SpecEvent = command.CommandName
	s.SpecType = command.EntityTypeName

	// Encrypt here
	//fmt.Println(command.Request.ProtoReflect())
	//configuration, err := acc.GetPlatformConfiguration(ctx, s.Context.WorkspaceSlug)
	//if err != nil {
	//	return nil, err
	//}

	// configuration.DataCatalog.Configuration.ConfigurationV2Alpha

	data, err := anypb.New(command.Request)
	if err != nil {
		log.Error(err.Error())
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("internal error"))
	}

	s.Data = data

	specBytes, err := proto.Marshal(s)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	subject := GetMultiplexedRequestSubjectName(command.Stream.StreamPrefix(), command.CommandTopic, command.Procedure)

	log.Debug("Publishing on " + subject)

	n := b.Nats

	reply, err := n.RequestMsg(&nats.Msg{
		Subject: subject,
		Data:    specBytes,
	}, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	return reply, err
}

// MultiplexEventSync sends an event to a multiplexed stream and waits for the response or error within the specified timeout.
// Uses Nats Publish and Subscribe Pattern
func (b *Binding) MultiplexEventSync(_ context.Context, s *specv2pb.Spec, event *SpecEvent) (*nats.Msg, error) {
	if event == nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecEvent object is required"))
	}

	log := *zaploggerv1.Bound.Logger

	s.SpecEvent = event.EventName
	s.SpecType = event.EntityTypeName

	data, err := anypb.New(event.Request)
	if err != nil {
		log.Error(err.Error())
		return nil, connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	s.Data = data

	specBytes, err := proto.Marshal(s)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	// Encrypt here

	subject := GetMultiplexedRequestSubjectName(event.Stream.StreamPrefix(), event.EventTopic, event.Procedure)

	log.Debug("Publishing on " + subject)

	n := b.Nats

	reply, err := n.RequestMsg(&nats.Msg{
		Subject: subject,
		Data:    specBytes,
	}, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	return reply, err
}

// MultiplexEventStreamSync sends an event to a multiplexed stream and waits for the response or error within the specified timeout.
// Uses Nats Sync Publish for streaming
func MultiplexEventStreamSync[T any](ctx context.Context, s *specv2pb.Spec, event *SpecStreamEvent, nats *nats.Conn, stream *connect.ServerStream[T], convert func(*nats.Msg) (*T, error)) error {
	if event == nil {
		return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecEvent object is required"))
	}

	log := *zaploggerv1.Bound.Logger

	s.SpecEvent = event.EventName
	s.SpecType = event.EntityTypeName

	data, err := anypb.New(event.Request)
	if err != nil {
		log.Error(err.Error())
		return connect.NewError(connect.CodeInternal, errors.New("internal error"))
	}

	s.Data = data

	subject := GetMultiplexedRequestSubjectName(event.Stream.StreamPrefix(), event.EventTopic, event.Procedure)
	responseSubject := GetStreamResponseSubjectName(event.Stream.StreamPrefix(), event.EventTopic, event.Procedure, s.MessageId)

	specBytes, err := proto.Marshal(s)
	if err != nil {
		return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	// Encrypt here

	log.Debug("Publishing on " + subject)
	if err = nats.Publish(subject, specBytes); err != nil {
		return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to publish")).WithInternalErrorDetail(err)
	}

	// Subscribe to streamed results
	log.Debug("Waiting on results from " + responseSubject)
	sub, err := nats.SubscribeSync(responseSubject)
	if err != nil {
		return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not subscribe stream sync to nats")).WithInternalErrorDetail(err)
	}

	defer sub.Unsubscribe()

	for {
		msg, err1 := sub.NextMsgWithContext(ctx)
		if err1 != nil {
			if errors.Is(err1, context.Canceled) {
				return nil
			}
			return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not process additional events")).WithInternalErrorDetail(err1)
		}

		converted, err1 := convert(msg)
		if err1 != nil {
			return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not convert event")).WithInternalErrorDetail(err1)
		}

		err1 = stream.Send(converted)
		if err != nil {
			return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not stream event")).WithInternalErrorDetail(err1)
		}
	}
}
