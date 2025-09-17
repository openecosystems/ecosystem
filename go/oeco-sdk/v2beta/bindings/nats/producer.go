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

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MultiplexCommandSync sends a command synchronously by publishing it to a NATS stream and awaiting a reply.
// Uses Nats Publish and Subscribe Pattern
func (b *Binding) MultiplexCommandSync(ctx context.Context, spec *specv2pb.Spec, command *SpecCommand) (*nats.Msg, error) {
	if command == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecCommand object is required"))
	}

	if spec == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a Spec object is required"))
	}

	log := *zaploggerv1.Bound.Logger
	spec.SpecEvent = command.CommandName
	spec.SpecType = command.EntityTypeName

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

	spec.Data = data

	specBytes, err := proto.Marshal(spec)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	subject := GetMultiplexedRequestSubjectName(command.Stream.StreamPrefix(), command.CommandTopic, command.Procedure)

	fields := receivedFields(spec, subject)
	log.Info("Issuing a multiplex command: "+command.Procedure, fields...)

	return publish(b.Nats, subject, spec, specBytes)
}

// MultiplexEventSync sends an event to a multiplexed stream and waits for the response or error within the specified timeout.
// Uses Nats Publish and Subscribe Pattern
func (b *Binding) MultiplexEventSync(_ context.Context, spec *specv2pb.Spec, event *SpecEvent) (*nats.Msg, error) {
	if event == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a SpecEvent object is required"))
	}

	if spec == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("a Spec object is required"))
	}

	log := *zaploggerv1.Bound.Logger
	spec.SpecEvent = event.EventName
	spec.SpecType = event.EntityTypeName

	data, err := anypb.New(event.Request)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("internal error"))
	}

	spec.Data = data

	specBytes, err := proto.Marshal(spec)
	if err != nil {
		return nil, sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("could not marshall spec"))
	}

	subject := GetMultiplexedRequestSubjectName(event.Stream.StreamPrefix(), event.EventTopic, event.Procedure)

	fields := receivedFields(spec, subject)
	log.Info("Issuing a multiplex event: "+event.Procedure, fields...)

	return publish(b.Nats, subject, spec, specBytes)
}

func publish(n *nats.Conn, subject string, spec *specv2pb.Spec, specBytes []byte) (*nats.Msg, error) {
	log := *zaploggerv1.Bound.Logger
	reply, err := n.RequestMsg(&nats.Msg{
		Subject: subject,
		Data:    specBytes,
	}, 10*time.Second)
	if err != nil {
		switch {
		case errors.Is(err, nats.ErrTimeout):
			return nil, ErrTimeout.WithSpecDetail(spec).WithInternalErrorDetail(err)
		case errors.Is(err, nats.ErrNoResponders):
			// no responders
			return nil, ErrNoResponders.WithSpecDetail(spec).WithInternalErrorDetail(err)
		case errors.Is(err, nats.ErrConnectionClosed):
			// NATS connection was closed
			return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(err)
		case errors.Is(err, nats.ErrBadSubscription):
			// something wrong with the sub
			return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(err)
		default:
			// unknown or generic error
			return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("unhandled NATS error"), err)
		}
	}

	if reply == nil {
		return nil, sdkv2betalib.ErrServerInternal.WithSpecDetail(spec).WithInternalErrorDetail(errors.New("received nil reply from NATS responder"))
	}

	spec.CompletedAt = timestamppb.Now()
	fields := completedFields(spec, subject)
	var milliseconds int64
	if spec.CompletedAt != nil && spec.ReceivedAt != nil {
		milliseconds = spec.CompletedAt.AsTime().Sub(spec.ReceivedAt.AsTime()).Milliseconds()
	}
	log.Info(fmt.Sprintf("Completed multiplexed request in %d ms\n", milliseconds), fields...)

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
