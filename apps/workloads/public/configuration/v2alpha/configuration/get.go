package main

import (
  "context"
  "go.uber.org/zap"
  "libs/partner/go/nats/v2"
  "libs/partner/go/zap/v1"
  "libs/private/go/configuration/v2alpha"
  "libs/protobuf/go/protobuf/gen/platform/spec/v2"
  "libs/protobuf/go/protobuf/gen/platform/type/v2"
  "libs/public/go/model/gen/platform/configuration/v2alpha"
  "libs/public/go/protobuf/gen/platform/configuration/v2alpha"
  "libs/public/go/sdk/v2alpha"

  "github.com/nats-io/nats.go/jetstream"
)

type GetConfigurationListener struct{}

func (l *GetConfigurationListener) GetConfiguration() *natsnodev2.ListenerConfiguration {
  entity := &configurationv2alphapbmodel.ConfigurationSpecEntity{}
  streamType := natsnodev2.InboundStream{}
  subject := natsnodev2.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.EventTopic())
  queue := natsnodev2.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

  return &natsnodev2.ListenerConfiguration{
    Entity:     &configurationv2alphapbmodel.ConfigurationSpecEntity{},
    Subject:    subject,
    Queue:      queue,
    StreamType: &natsnodev2.InboundStream{},
    JetstreamConfiguration: &jetstream.ConsumerConfig{
      Durable: "listener-configuration-getConfiguration",
      //Durable: natsnodev2.GetListenerGroup(
      //	&configurationv2alphapb.ConfigurationSpecEntity{},
      //	&configurationv2alphapb.ConfigurationSpecEntity{},
      //),
      AckPolicy:     jetstream.AckExplicitPolicy,
      MemoryStorage: false,
      FilterSubject: "inbound-configuration.data.event",
      Metadata:      nil,
    },
  }
}

func (l *GetConfigurationListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
  natsnodev2.ListenForMultiplexedSpecEventsSync(ctx, l)
}

func (l *GetConfigurationListener) Process(ctx context.Context, request *natsnodev2.ListenerMessage) {
  log := *zaploggerv1.Bound.Logger
  acc := *configurationv2alphalib.Bound.AdaptiveConfigurationControl

  if request.Spec == nil {
    return
  }

  if request.Spec.Context.OrganizationSlug == "" {
    log.Error("Organization Slug is required. Quietly dropping this message")
    return
  }

  configuration, err := acc.GetPlatformConfiguration(ctx, request.Spec.Context.WorkspaceSlug)
  if err != nil {
    return
  }

  response := &configurationv2alphapb.GetConfigurationResponse{
    SpecContext: &specv2pb.SpecResponseContext{
      ResponseValidation: &typev2pb.ResponseValidation{
        ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
      },
      OrganizationSlug: request.Spec.Context.OrganizationSlug,
      WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
      WorkspaceJan:     request.Spec.Context.WorkspaceJan,
    },
    Configuration: configuration,
  }

  log.Info("Get Configuration Response", zap.Any("response", response))

  natsnodev2.RespondToSyncCommand(ctx, request, response)
}
