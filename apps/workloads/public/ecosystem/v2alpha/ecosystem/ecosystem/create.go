package ecosystem

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	natsnodev1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/nats"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
	ecosystemv2alphapb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/ecosystem/v2alpha"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"
)

// CreateEcosystemListener is a struct that listens for create configuration events and processes them.
type CreateEcosystemListener struct{}

// GetConfiguration returns the listener configuration for the CreateEcosystemListener, including entity, subject, and queue details.
func (l *CreateEcosystemListener) Configure() *natsnodev1.ListenerConfiguration {
	handler := ecosystemv2alphapb.EcosystemServiceHandler{}
	return handler.GetCreateEcosystemConfiguration()
}

// Listen starts the listener to process multiplexed spec events synchronously based on the provided context and configuration.
func (l *CreateEcosystemListener) Listen(ctx context.Context, _ chan sdkv2betalib.SpecListenableErr) {
	natsnodev1.ListenForMultiplexedRequests(ctx, l)
}

// Process handles incoming listener messages to create and store a configuration, ensuring required fields are validated.
func (l *CreateEcosystemListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger

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
