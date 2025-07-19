package accountauthority

//
//import (
//	"context"
//
//	"github.com/nats-io/nats.go/jetstream"
//
//	natsnodev1 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/nats"
//	nebulav1ca "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/nebula/ca"
//	zaploggerv1 "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/zap"
//	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/sdk/v2beta/gen/platform/spec/v2"
//	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/sdk/v2beta/gen/platform/type/v2"
//	iamv2alphapbmodel "github.com/openecosystems/ecosystem/libs/public/go/model/gen/platform/iam/v2alpha"
//	iamv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/gen/platform/iam/v2alpha"
//	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"
//)
//
//// CreateAccountAuthorityListener represents a listener for handling requests to create a account authority.
//type CreateAccountAuthorityListener struct{}
//
//// GetConfiguration provides the listener configuration for CreateAccountAuthorityListener, including subject, queue, and jetstream settings.
//func (l *CreateAccountAuthorityListener) GetConfiguration() *natsnodev1.ListenerConfiguration {
//	entity := &iamv2alphapbmodel.AccountAuthoritySpecEntity{}
//	streamType := natsnodev1.InboundStream{}
//	subject := natsnodev1.GetMultiplexedRequestSubjectName(streamType.StreamPrefix(), entity.CommandTopic())
//	queue := natsnodev1.GetQueueGroupName(streamType.StreamPrefix(), entity.TypeName())
//
//	return &natsnodev1.ListenerConfiguration{
//		Entity:     &iamv2alphapbmodel.AccountAuthoritySpecEntity{},
//		Subject:    subject,
//		Queue:      queue,
//		StreamType: &natsnodev1.InboundStream{},
//		JetstreamConfiguration: &jetstream.ConsumerConfig{
//			Durable:       "iam-createAccountAuthority",
//			AckPolicy:     jetstream.AckExplicitPolicy,
//			MemoryStorage: false,
//			FilterSubject: "inbound-account-authority.data.command",
//			Metadata:      nil,
//		},
//	}
//}
//
//// Listen synchronously listens for multiplexed spec events and routes them to the associated handler.
//func (l *CreateAccountAuthorityListener) Listen(ctx context.Context, _ chan sdkv2betalib.SpecListenableErr) {
//	natsnodev1.ListenForMultiplexedSpecEventsSync(ctx, l)
//}
//
//// Process handles the incoming ListenerMessage, processes the request, and sends an appropriate response back to the client.
//// It validates the Spec field in the request, extracts the necessary data, retrieves or creates a Account Authority,
//// and constructs a response to be sent. Logs errors and success for debugging and tracking purposes.
//func (l *CreateAccountAuthorityListener) Process(ctx context.Context, request *natsnodev1.ListenerMessage) {
//	log := *zaploggerv1.Bound.Logger
//	nca := *nebulav1ca.Bound
//
//	if request.Spec == nil {
//		return
//	}
//
//	var req iamv2alphapb.CreateAccountAuthorityRequest
//	if err := request.Spec.Data.UnmarshalTo(&req); err != nil {
//		log.Error("Failed to unpack Any message: " + err.Error())
//	}
//
//	ca, err := nca.GetAccountAuthority(ctx, &req)
//	if err != nil {
//		return
//	}
//
//	response := iamv2alphapb.CreateAccountAuthorityResponse{
//		SpecContext: &specv2pb.SpecResponseContext{
//			ResponseValidation: &typev2pb.ResponseValidation{
//				ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
//			},
//			ResponseMask: &typev2pb.ResponseMask{
//				FieldMask:  request.Spec.SpecData.FieldMask,
//				PolicyMask: nil,
//			},
//			OrganizationSlug: request.Spec.Context.OrganizationSlug,
//			WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
//			WorkspaceJan:     request.Spec.Context.WorkspaceJan,
//		},
//		AccountAuthority: ca,
//	}
//
//	log.Info("Created account authority successfully: " + response.AccountAuthority.Id)
//
//	natsnodev1.RespondToSyncCommand(ctx, request, &response)
//}
