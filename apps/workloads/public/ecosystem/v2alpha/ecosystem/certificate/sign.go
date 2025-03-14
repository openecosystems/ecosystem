package certificate

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"

	natsnodev1 "github.com/openecosystems/ecosystem/libs/partner/go/nats"
	zaploggerv1 "github.com/openecosystems/ecosystem/libs/partner/go/zap"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	cryptographyv2alphapbmodel "github.com/openecosystems/ecosystem/libs/public/go/model/gen/platform/cryptography/v2alpha"
	cryptographyv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

// SignCertificateListener represents a listener for handling requests to create a account authority.
type SignCertificateListener struct{}

// GetConfiguration provides the listener configuration for SignCertificateListener, including subject, queue, and jetstream settings.
func (l *SignCertificateListener) GetConfiguration() *natsnodev1.ListenerConfiguration {
	entity := &cryptographyv2alphapbmodel.CertificateSpecEntity{}
	streamType := natsnodev1.InboundStream{}
	subject := natsnodev1.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
	queue := natsnodev1.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())

	return &natsnodev1.ListenerConfiguration{
		Entity:     &cryptographyv2alphapbmodel.CertificateSpecEntity{},
		Subject:    subject,
		Queue:      queue,
		StreamType: &natsnodev1.InboundStream{},
		JetstreamConfiguration: &jetstream.ConsumerConfig{
			Durable:       "cryptography-signCertificate",
			AckPolicy:     jetstream.AckExplicitPolicy,
			MemoryStorage: false,
			FilterSubject: "inbound-certificate.data.command",
			Metadata:      nil,
		},
	}
}

// Listen synchronously listens for multiplexed spec events and routes them to the associated handler.
func (l *SignCertificateListener) Listen(ctx context.Context, _ chan sdkv2alphalib.SpecListenableErr) {
	natsnodev1.ListenForMultiplexedSpecEventsSync(ctx, l)
}

// Process handles the incoming ListenerMessage, processes the request, and sends an appropriate response back to the client.
// It validates the Spec field in the request, extracts the necessary data, retrieves or creates a Account Authority,
// and constructs a response to be sent. Logs errors and success for debugging and tracking purposes.
func (l *SignCertificateListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
	log := *zaploggerv1.Bound.Logger
	// nca := *nebulav1ca.Bound

	if request.Spec == nil {
		return
	}

	var req cryptographyv2alphapb.SignCertificateRequest
	if err := request.Spec.Data.UnmarshalTo(&req); err != nil {
		log.Error("Failed to unpack Any message: " + err.Error())
	}

	// fmt.Println(req.Certificate.Name)

	// nca.Sign()

	response := cryptographyv2alphapb.SignCertificateResponse{
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
		// Certificate: ,
	}

	// log.Info("Signed certificate successfully: " + response.Certificate.Id)

	natsnodev1.RespondToSyncCommand(ctx, request, &response)
}
