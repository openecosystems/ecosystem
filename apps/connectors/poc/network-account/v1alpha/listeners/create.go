package listeners

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	configurationv2alphalib "github.com/openecosystems/ecosystem/libs/partner/go/configuration/v2alpha"
	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	ecosystemv2alphapbmodel "github.com/openecosystems/ecosystem/libs/public/go/model/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/ecosystem/v2alpha"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/ecosystem/v2alpha/ecosystemv2alphapbconnect"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// CreateEcosystemListener is a struct that listens for create configuration events and processes them.
type CreateEcosystemListener struct{}

// GetConfiguration returns the listener configuration for the CreateEcosystemListener, including entity, subject, and queue details.
func (l *CreateEcosystemListener) GetConfiguration() *natsnodev1.ListenerConfiguration {
	return &natsnodev1.ListenerConfiguration{
		Entity:     &ecosystemv2alphapbmodel.EcosystemSpecEntity{},
		StreamType: &natsnodev1.InboundStream{},
		Procedure:  ecosystemv2alphapbconnect.EcosystemServiceCreateEcosystemProcedure,
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
	natsnodev1.ListenForMultiplexedRequests(ctx, l)
}

// Process handles incoming listener messages to create and store a configuration, ensuring required fields are validated.
func (l *CreateEcosystemListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
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

	natsnodev1.RespondToMultiplexedRequest(ctx, request, &response)
}
