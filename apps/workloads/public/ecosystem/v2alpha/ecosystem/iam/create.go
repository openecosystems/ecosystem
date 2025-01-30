package iam

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	natsnodev2 "libs/partner/go/nats/v2"
	zaploggerv1 "libs/partner/go/zap/v1"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	iamv2alphapbmodel "libs/public/go/model/gen/platform/iam/v2alpha"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// CreateAccountListener is a struct that listens for create configuration events and processes them.
type CreateAccountListener struct{}

// GetConfiguration returns the listener configuration for the CreateAccountListener, including entity, subject, and queue details.
func (l *CreateAccountListener) GetConfiguration() *natsnodev2.ListenerConfiguration {
	entity := &iamv2alphapbmodel.AccountSpecEntity{}
	streamType := natsnodev2.InboundStream{}
	subject := natsnodev2.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
	queue := natsnodev2.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

	return &natsnodev2.ListenerConfiguration{
		Entity:     &iamv2alphapbmodel.AccountSpecEntity{},
		Subject:    subject,
		Queue:      queue,
		StreamType: &natsnodev2.InboundStream{},
		JetstreamConfiguration: &jetstream.ConsumerConfig{
			Durable:       "iam-createAccount",
			AckPolicy:     jetstream.AckExplicitPolicy,
			MemoryStorage: false,
			FilterSubject: "inbound-iam.data.command",
			Metadata:      nil,
		},
	}
}

// Listen starts the listener to process multiplexed spec events synchronously based on the provided context and configuration.
func (l *CreateAccountListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev2.ListenForMultiplexedSpecEventsSync(ctx, l)
}

// Process handles incoming listener messages to create and store a configuration, ensuring required fields are validated.
func (l *CreateAccountListener) Process(ctx context.Context, request *natsnodev2.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger

	if request.Spec == nil {
		return
	}

	req := iamv2alphapb.CreateAccountRequest{}
	err := request.Spec.Data.UnmarshalTo(&req)
	if err != nil {
		return
	}

	// Sign Cert

	now := timestamppb.Now()
	conf := iamv2alphapb.Account{
		Id:         "12345678",
		CreatedAt:  now,
		UpdatedAt:  now,
		SignedCert: req.Cert,
	}

	response := iamv2alphapb.CreateAccountResponse{
		SpecContext: &specv2pb.SpecResponseContext{
			ResponseValidation: &typev2pb.ResponseValidation{
				ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
			},
			OrganizationSlug: request.Spec.Context.OrganizationSlug,
			WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
			WorkspaceJan:     request.Spec.Context.WorkspaceJan,
		},
		Account: &conf,
	}
	log.Info("Create Account Response", zap.Any("id", response.Account.Id))

	natsnodev2.RespondToSyncCommand(ctx, request, &response)
}
