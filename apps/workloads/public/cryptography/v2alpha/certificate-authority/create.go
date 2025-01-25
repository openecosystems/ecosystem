package main

import (
	"context"
	natsnodev2 "libs/partner/go/nats/v2"
	zaploggerv1 "libs/partner/go/zap/v1"

	"github.com/nats-io/nats.go/jetstream"

	nebulav1ca "libs/partner/go/nebula/v1/ca"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	cryptographyv2alphapbmodel "libs/public/go/model/gen/platform/cryptography/v2alpha"
	cryptographyv2alphapb "libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// CreateCertificateAuthorityListener represents a listener for handling requests to create a certificate authority.
type CreateCertificateAuthorityListener struct{}

// GetConfiguration provides the listener configuration for CreateCertificateAuthorityListener, including subject, queue, and jetstream settings.
func (l *CreateCertificateAuthorityListener) GetConfiguration() *natsnodev2.ListenerConfiguration {
	entity := &cryptographyv2alphapbmodel.CertificateAuthoritySpecEntity{}
	streamType := natsnodev2.InboundStream{}
	subject := natsnodev2.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
	queue := natsnodev2.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

	return &natsnodev2.ListenerConfiguration{
		Entity:     &cryptographyv2alphapbmodel.CertificateAuthoritySpecEntity{},
		Subject:    subject,
		Queue:      queue,
		StreamType: &natsnodev2.InboundStream{},
		JetstreamConfiguration: &jetstream.ConsumerConfig{
			Durable:       "listener-cryptography-createCertificateAuthority",
			AckPolicy:     jetstream.AckExplicitPolicy,
			MemoryStorage: false,
			FilterSubject: "inbound-certificate-authority.data.command",
			Metadata:      nil,
		},
	}
}

// Listen synchronously listens for multiplexed spec events and routes them to the associated handler.
func (l *CreateCertificateAuthorityListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev2.ListenForMultiplexedSpecEventsSync(ctx, l)
}

// Process handles the incoming ListenerMessage, processes the request, and sends an appropriate response back to the client.
// It validates the Spec field in the request, extracts the necessary data, retrieves or creates a Certificate Authority,
// and constructs a response to be sent. Logs errors and success for debugging and tracking purposes.
func (l *CreateCertificateAuthorityListener) Process(ctx context.Context, request *natsnodev2.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	nca := *nebulav1ca.Bound

	if request.Spec == nil {
		return
	}

	var req cryptographyv2alphapb.CreateCertificateAuthorityRequest
	if err := request.Spec.Data.UnmarshalTo(&req); err != nil {
		log.Error("Failed to unpack Any message: " + err.Error())
	}

	ca, err := nca.GetCertificateAuthority(ctx, &req)
	if err != nil {
		return
	}

	response := cryptographyv2alphapb.CreateCertificateAuthorityResponse{
		SpecContext: &specv2pb.SpecResponseContext{
			ResponseValidation: &typev2pb.ResponseValidation{
				ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
			},
			ResponseMask: &typev2pb.ResponseMask{
				FieldMask:  request.Spec.SpecData.FieldMask,
				PolicyMask: nil,
			},
			OrganizationSlug: request.Spec.Context.OrganizationSlug,
			WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
			WorkspaceJan:     request.Spec.Context.WorkspaceJan,
		},
		CertificateAuthority: ca,
	}

	log.Info("Create certificate authority successfully: " + response.CertificateAuthority.Id)

	natsnodev2.RespondToSyncCommand(ctx, request, &response)
}
