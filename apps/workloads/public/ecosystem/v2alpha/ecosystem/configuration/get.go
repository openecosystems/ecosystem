package configuration

import (
	"context"

	"go.uber.org/zap"

	configurationv2alphalib "github.com/openecosystems/ecosystem/libs/partner/go/configuration/v2alpha"
	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	configurationv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/configuration/v2alpha"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// GetConfigurationListener represents a listener responsible for handling configuration-related events and requests.
type GetConfigurationListener struct{}

// GetConfiguration creates and returns a ListenerConfiguration for the GetConfigurationListener.
func (l *GetConfigurationListener) GetConfiguration() *natsnodev1.ListenerConfiguration {
	handler := configurationv2alphapb.ConfigurationServiceHandler{}
	return handler.GetGetConfigurationConfiguration()
}

// Listen subscribes the listener to a NATS subject to process multiplexed specification events synchronously.
func (l *GetConfigurationListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev1.ListenForMultiplexedRequests(ctx, l)
}

// Process handles incoming listener messages, validates the request, retrieves platform configurations, and sends a response.
func (l *GetConfigurationListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
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

	natsnodev1.RespondToMultiplexedRequest(ctx, request, response)
}
