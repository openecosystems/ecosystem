package natsnodev1

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
)

// MultiplexCommandSync sends a command synchronously by publishing it to a NATS stream and awaiting a reply.
// Uses Nats Publish and Subscribe Pattern
func (b *Binding) MultiplexCommandSync(ctx context.Context, s *specv2pb.Spec, command *SpecCommand) (*nats.Msg, error) {
	if command == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecCommand object is required"))
	}

	log := *zaploggerv1.Bound.Logger
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
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("internal error"))
	}

	s.Data = data

	specBytes, err := proto.Marshal(s)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	subject := GetMultiplexedRequestSubjectName(command.Stream.StreamPrefix(), command.CommandTopic, command.Procedure)

	log.Debug("Issuing a multiplex command: " + command.Procedure + ", on channel: " + subject)

	return publish(b.Nats, subject, s, specBytes)
}

// MultiplexEventSync sends an event to a multiplexed stream and waits for the response or error within the specified timeout.
// Uses Nats Publish and Subscribe Pattern
func (b *Binding) MultiplexEventSync(_ context.Context, s *specv2pb.Spec, event *SpecEvent) (*nats.Msg, error) {
	if event == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecEvent object is required"))
	}

	log := *zaploggerv1.Bound.Logger
	s.SpecEvent = event.EventName
	s.SpecType = event.EntityTypeName

	data, err := anypb.New(event.Request)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("internal error"))
	}

	s.Data = data

	specBytes, err := proto.Marshal(s)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	subject := GetMultiplexedRequestSubjectName(event.Stream.StreamPrefix(), event.EventTopic, event.Procedure)

	log.Debug("Issuing a multiplex event: " + event.Procedure + ", on channel: " + subject)

	return publish(b.Nats, subject, s, specBytes)
}

func publish(n *nats.Conn, subject string, s *specv2pb.Spec, specBytes []byte) (*nats.Msg, error) {
	reply, err := n.RequestMsg(&nats.Msg{
		Subject: subject,
		Data:    specBytes,
	}, 10*time.Second)
	if err != nil {
		switch {
		case errors.Is(err, nats.ErrTimeout):
			return nil, ErrTimeout.WithSpecDetail(s)
		case errors.Is(err, nats.ErrNoResponders):
			// no responders
			return nil, ErrNoResponders.WithSpecDetail(s)
		case errors.Is(err, nats.ErrConnectionClosed):
			// NATS connection was closed
			return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(err)
		case errors.Is(err, nats.ErrBadSubscription):
			// something wrong with the sub
			return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(err)
		default:
			// unknown or generic error
			return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("unhandled NATS error"), err)
		}
	}

	if reply == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("received nil reply from NATS responder"))
	}

	return reply, err
}

// MultiplexEventStreamSync sends an event to a multiplexed stream and waits for the response or error within the specified timeout.
// Uses Nats Sync Publish for streaming
func MultiplexEventStreamSync[T any](ctx context.Context, s *specv2pb.Spec, event *SpecStreamEvent, nats *nats.Conn, stream *connect.ServerStream[T], convert func(*nats.Msg) (*T, error)) error {
	if event == nil {
		return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("a SpecEvent object is required"))
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
		return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	// Encrypt here

	log.Debug("Publishing on " + subject)
	if err = nats.Publish(subject, specBytes); err != nil {
		return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("failed to publish")).WithInternalErrorDetail(err)
	}

	// Subscribe to streamed results
	log.Debug("Waiting on results from " + responseSubject)
	sub, err := nats.SubscribeSync(responseSubject)
	if err != nil {
		return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("could not subscribe stream sync to nats")).WithInternalErrorDetail(err)
	}

	defer sub.Unsubscribe() // nolint:errcheck

	for {
		msg, err1 := sub.NextMsgWithContext(ctx)
		if err1 != nil {
			if errors.Is(err1, context.Canceled) {
				return nil
			}
			return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("could not process additional events")).WithInternalErrorDetail(err1)
		}

		converted, err1 := convert(msg)
		if err1 != nil {
			return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("could not convert event")).WithInternalErrorDetail(err1)
		}

		err1 = stream.Send(converted)
		if err != nil {
			return sdkv2betalib.ErrServerInternal.WithSpecDetail(s).WithInternalErrorDetail(errors.New("could not stream event")).WithInternalErrorDetail(err1)
		}
	}
}
