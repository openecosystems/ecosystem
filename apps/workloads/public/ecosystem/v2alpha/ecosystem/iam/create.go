package iam

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	nebulav1ca "github.com/openecosystems/ecosystem/libs/partner/go/nebula/ca"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	iamv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/iam/v2alpha"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// CreateAccountListener is a struct that listens for create configuration events and processes them.
type CreateAccountListener struct{}

// GetConfiguration returns the listener configuration for the CreateAccountListener, including entity, subject, and queue details.
func (l *CreateAccountListener) GetConfiguration() *natsnodev1.ListenerConfiguration {
	handler := iamv2alphapb.AccountAuthorityServiceHandler{}
	return handler.GetCreateAccountAuthorityConfiguration()
}

// Listen starts the listener to process multiplexed spec events synchronously based on the provided context and configuration.
func (l *CreateAccountListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev1.ListenForMultiplexedRequests(ctx, l)
}

// Process handles incoming listener messages to create and store a configuration, ensuring required fields are validated.
func (l *CreateAccountListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger

	if request.Spec == nil {
		return
	}

	req := iamv2alphapb.CreateAccountRequest{}
	err := request.Spec.Data.UnmarshalTo(&req)
	if err != nil {
		return
	}

	if req.Name == "" {
		fmt.Println("Name is required")
		return
	}

	signreq := iamv2alphapb.SignAccountRequest{
		Name:       req.Name,
		PublicCert: req.Cert,
	}

	// Sign Cert
	credential, err := nebulav1ca.Bound.SignCert(ctx, &signreq)
	if err != nil {
		return
	}

	now := timestamppb.Now()
	conf := iamv2alphapb.Account{
		Id:         "12345678",
		CreatedAt:  now,
		UpdatedAt:  now,
		Credential: credential,
	}

	response := iamv2alphapb.CreateAccountResponse{
		SpecContext: &specv2pb.SpecResponseContext{
			ResponseValidation: &typev2pb.ResponseValidation{
				ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
			},
			EcosystemSlug:    request.Spec.Context.EcosystemSlug,
			OrganizationSlug: request.Spec.Context.OrganizationSlug,
			WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
			WorkspaceJan:     request.Spec.Context.WorkspaceJan,
		},
		Account: &conf,
	}
	log.Info("Create Account Response", zap.Any("id", response.Account.Id))

	natsnodev1.RespondToMultiplexedRequest(ctx, request, &response)
}
