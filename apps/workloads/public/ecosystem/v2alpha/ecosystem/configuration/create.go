package configuration

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats/v1"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap/v1"
	configurationv2alphalib "github.com/openecosystems/ecosystem/libs/private/go/configuration/v2alpha"
	configurationdefaultsv2alphalib "github.com/openecosystems/ecosystem/libs/private/go/configuration/v2alpha/defaults"
	ontologydefaultsv2alphalib "github.com/openecosystems/ecosystem/libs/private/go/ontology/v2alpha/defaults"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	configurationv2alphapbmodel "github.com/openecosystems/ecosystem/libs/public/go/model/gen/platform/configuration/v2alpha"
	configurationv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// CreateConfigurationListener is a struct that listens for create configuration events and processes them.
type CreateConfigurationListener struct{}

// GetConfiguration returns the ListenerConfiguration for CreateConfigurationListener, defining subject, queue, entity, and stream settings.
func (l *CreateConfigurationListener) GetConfiguration() *natsnodev1.ListenerConfiguration {
	entity := &configurationv2alphapbmodel.ConfigurationSpecEntity{}
	streamType := natsnodev1.InboundStream{}
	subject := natsnodev1.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
	queue := natsnodev1.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

	return &natsnodev1.ListenerConfiguration{
		Entity:     &configurationv2alphapbmodel.ConfigurationSpecEntity{},
		Subject:    subject,
		Queue:      queue,
		StreamType: &natsnodev1.InboundStream{},
		JetstreamConfiguration: &jetstream.ConsumerConfig{
			Durable:       "configuration-createConfiguration",
			AckPolicy:     jetstream.AckExplicitPolicy,
			MemoryStorage: false,
			FilterSubject: "inbound-configuration.data.command",
			Metadata:      nil,
		},
	}
}

// Listen starts the listener to process multiplexed spec events synchronously based on the provided context and configuration.
func (l *CreateConfigurationListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev1.ListenForMultiplexedSpecEventsSync(ctx, l)
}

// Process handles incoming listener messages to create and store a configuration, ensuring required fields are validated.
func (l *CreateConfigurationListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	acc := *configurationv2alphalib.Bound.AdaptiveConfigurationControl

	if request.Spec == nil {
		return
	}

	//if request.Spec.Context.OrganizationSlug == "" {
	//	log.Error("Organization Slug is required. Quietly dropping this message")
	//	return
	//}

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
	log.Info("Create Configuration Response", zap.Any("id", response.Configuration.Id))

	natsnodev1.RespondToSyncCommand(ctx, request, &response)
}
