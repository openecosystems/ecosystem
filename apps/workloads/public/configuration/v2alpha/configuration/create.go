package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"google.golang.org/protobuf/types/known/timestamppb"
	"libs/partner/go/nats/v2"
	"libs/partner/go/zap/v1"
	"libs/private/go/configuration/v2alpha"
	"libs/private/go/configuration/v2alpha/defaults"
	"libs/private/go/ontology/v2alpha/defaults"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/protobuf/go/protobuf/gen/platform/type/v2"
	configurationv2alphapbmodel "libs/public/go/model/gen/platform/configuration/v2alpha"
	"libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"google.golang.org/protobuf/proto"
)

type CreateConfigurationListener struct{}

func (l *CreateConfigurationListener) GetConfiguration() *natsnodev2.ListenerConfiguration {

	entity := &configurationv2alphapbmodel.ConfigurationSpecEntity{}
	streamType := natsnodev2.InboundStream{}
	subject := natsnodev2.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
	queue := natsnodev2.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

	return &natsnodev2.ListenerConfiguration{
		Entity:     &configurationv2alphapbmodel.ConfigurationSpecEntity{},
		Subject:    subject,
		Queue:      queue,
		StreamType: &natsnodev2.InboundStream{},
		JetstreamConfiguration: &jetstream.ConsumerConfig{
			Durable:       "listener-configuration-createConfiguration",
			AckPolicy:     jetstream.AckExplicitPolicy,
			MemoryStorage: false,
			FilterSubject: "inbound-configuration.data.command",
			Metadata:      nil,
		},
	}

}

func (l *CreateConfigurationListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev2.ListenForMultiplexedSpecEventsSync(ctx, l)
}

func (l *CreateConfigurationListener) Process(ctx context.Context, request *natsnodev2.ListenerMessage) {

	log := *zaploggerv1.Bound.Logger
	acc := *configurationv2alphalib.Bound.AdaptiveConfigurationControl

	fmt.Println("CREATE")

	fmt.Println(request.Spec)

	if request.Spec == nil {
		return
	}

	if request.Spec.Context.OrganizationSlug == "" {
		log.Error("Organization Slug is required. Quietly dropping this message")
		return
	}

	now := timestamppb.Now()
	conf := configurationv2alphapb.Configuration{
		Id:                    "12345678",
		CreatedAt:             now,
		UpdatedAt:             now,
		Type:                  configurationv2alphapb.ConfigurationType_CONFIGURATION_TYPE_WORKSPACE,
		Status:                configurationv2alphapb.ConfigurationStatus_CONFIGURATION_STATUS_ACTIVE,
		StatusDetails:         "",
		ParentId:              "",
		DataCatalog:           &ontologydefaultsv2alphalib.Hippa,
		PlatformConfiguration: &configurationdefaultsv2alphalib.DefaultEnterpriseConfiguration,
	}

	b, err := proto.Marshal(&conf)
	if err != nil {
		log.Error(err.Error())
		return
	}

	err2 := acc.SavePlatformConfiguration(ctx, request.Spec.Context.WorkspaceSlug, b)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	response := configurationv2alphapb.CreateConfigurationResponse{
		SpecContext: &specv2pb.SpecResponseContext{
			ResponseValidation: &typev2pb.ResponseValidation{
				ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
			},
			OrganizationSlug: request.Spec.Context.OrganizationSlug,
			WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
			WorkspaceJan:     request.Spec.Context.WorkspaceJan,
		},
		Configuration: &conf,
	}

	natsnodev2.RespondToSyncCommand(ctx, request, &response)
}
