package natsnodev2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"libs/partner/go/zap/v1"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/public/go/sdk/v2alpha"

	"connectrpc.com/connect"

	zaploggerv1 "libs/partner/go/zap/v1"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func (b *Binding) MultiplexCommandSync(ctx context.Context, s *specv2pb.Spec, command *SpecCommand) (*nats.Msg, error) {
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

	subject := GetMultiplexedRequestSubjectName(command.Stream.StreamPrefix(), command.CommandTopic)

	log.Debug("Publishing on " + subject)

	n := b.Nats

	reply, err := n.RequestMsg(&nats.Msg{
		Subject: subject,
		Data:    specBytes,
	}, 4*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	return reply, err
}

func (b *Binding) MultiplexEventSync(ctx context.Context, s *specv2pb.Spec, event *SpecEvent) (*nats.Msg, error) {
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

	subject := GetMultiplexedRequestSubjectName(event.Stream.StreamPrefix(), event.EventTopic)

	log.Debug("Publishing on " + subject)

	n := b.Nats

	reply, err := n.RequestMsg(&nats.Msg{
		Subject: subject,
		Data:    specBytes,
	}, 4*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	return reply, err
}
