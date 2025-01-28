package ecosystem

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	natsnodev2 "libs/partner/go/nats/v2"
	zaploggerv1 "libs/partner/go/zap/v1"
	configurationv2alphalib "libs/private/go/configuration/v2alpha"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	ecosystemv2alphapbmodel "libs/public/go/model/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// CreateEcosystemListener is a struct that listens for create configuration events and processes them.
type CreateEcosystemListener struct{}

// GetConfiguration returns the listener configuration for the CreateEcosystemListener, including entity, subject, and queue details.
func (l *CreateEcosystemListener) GetConfiguration() *natsnodev2.ListenerConfiguration {
	entity := &ecosystemv2alphapbmodel.EcosystemSpecEntity{}
	streamType := natsnodev2.InboundStream{}
	subject := natsnodev2.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
	queue := natsnodev2.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

	return &natsnodev2.ListenerConfiguration{
		Entity:     &ecosystemv2alphapbmodel.EcosystemSpecEntity{},
		Subject:    subject,
		Queue:      queue,
		StreamType: &natsnodev2.InboundStream{},
		JetstreamConfiguration: &jetstream.ConsumerConfig{
			Durable:       "ecosystem-createEcosystem",
			AckPolicy:     jetstream.AckExplicitPolicy,
			MemoryStorage: false,
			FilterSubject: "inbound-ecosystem.data.command",
			Metadata:      nil,
		},
	}
}

// Listen starts the listener to process multiplexed spec events synchronously based on the provided context and configuration.
func (l *CreateEcosystemListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev2.ListenForMultiplexedSpecEventsSync(ctx, l)
}

// Process handles incoming listener messages to create and store a configuration, ensuring required fields are validated.
func (l *CreateEcosystemListener) Process(ctx context.Context, request *natsnodev2.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	acc := *configurationv2alphalib.Bound.AdaptiveConfigurationControl

	if request.Spec == nil {
		return
	}

	if request.Spec.Context.OrganizationSlug == "" {
		log.Error("Organization Slug is required. Quietly dropping this message")
		return
	}

	now := timestamppb.Now()
	conf := ecosystemv2alphapb.Ecosystem{
		Id:            "12345678",
		CreatedAt:     now,
		UpdatedAt:     now,
		Type:          ecosystemv2alphapb.EcosystemType_ECOSYSTEM_TYPE_PUBLIC,
		Status:        ecosystemv2alphapb.EcosystemStatus_ECOSYSTEM_STATUS_ACTIVE,
		StatusDetails: "",
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

	response := ecosystemv2alphapb.CreateEcosystemResponse{
		SpecContext: &specv2pb.SpecResponseContext{
			ResponseValidation: &typev2pb.ResponseValidation{
				ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
			},
			OrganizationSlug: request.Spec.Context.OrganizationSlug,
			WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
			WorkspaceJan:     request.Spec.Context.WorkspaceJan,
		},
		Ecosystem: &conf,
	}
	log.Info("Create Ecosystem Response", zap.Any("id", response.Ecosystem.Id))

	natsnodev2.RespondToSyncCommand(ctx, request, &response)
}
