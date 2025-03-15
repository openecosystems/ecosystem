package main

import (
	"testing"

	"github.com/openecosystems/ecosystem/libs/partner/go/sendgrid/v3/contacts"
	"github.com/openecosystems/ecosystem/libs/partner/go/sendgrid/v3/lists"
	"github.com/openecosystems/ecosystem/libs/partner/go/zap"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

func TestServer(t *testing.T) {
	t.Parallel()

	_ = []sdkv2alphalib.Binding{
		&zaploggerv1.Binding{},
		&sendgridcontactsv3.Binding{},
		&sendgridlistv3.Binding{},
	}

	//serviceHandler := &communicationv1alphapb.PreferenceCenterService{
	//	//QueryHandler:    &query.Handler{},
	//	//MutationHandler: &mutation.Handler{},
	//}
	//
	//interceptors := connect.WithInterceptors(sdkv2alphalib.NewSpecInterceptor())
	//path, handler := communicationv1alphapbconnect.NewPreferenceCenterServiceHandler(serviceHandler, interceptors)
	//server := communicationv1alphapb.RegisterPreferenceCenterServiceSpecServer(bounds, path, &handler)
	//
	//testServer := httptest.NewUnstartedServer(server.HttpServerHandler)
	//testServer.EnableHTTP2 = true
	//testServer.StartTLS()
	//defer testServer.Close()
	//
	//connectClient := communicationv1alphapbconnect.NewPreferenceCenterServiceClient(
	//	testServer.Client(),
	//	testServer.URL,
	//)
	//grpcClient := communicationv1alphapbconnect.NewPreferenceCenterServiceClient(
	//	testServer.Client(),
	//	testServer.URL,
	//	connect.WithGRPC(),
	//)
	//
	//clients := []communicationv1alphapbconnect.PreferenceCenterServiceClient{connectClient, grpcClient}
	//
	//t.Run("get preference options", func(t *testing.T) {
	//	for _, client := range clients {
	//
	//		fmt.Println(client)
	//		_, err := client.GetPreferenceOptions(context.Background(), connect.NewRequest(&communicationv1alphapb.GetPreferenceOptionsRequest{}))
	//
	//		fmt.Println(err)
	//		//require.NoError(t, err)
	//		//assert.NotEmpty(t, result.Msg.GetIndustries())
	//	}
	//})

	//
	//t.Run("converse", func(t *testing.T) {
	//	for _, client := range clients {
	//		sendValues := []string{"Hello!", "How are you doing?", "I have an issue with my bike", "bye"}
	//		var receivedValues []string
	//		grp, ctx := errgroup.WithContext(context.Background())
	//		stream := client.Converse(ctx)
	//		grp.Go(func() error {
	//			for _, sentence := range sendValues {
	//				err := stream.Send(&elizav1.ConverseRequest{Sentence: sentence})
	//				if err != nil {
	//					return err
	//				}
	//			}
	//			return stream.CloseRequest()
	//		})
	//		grp.Go(func() error {
	//			for {
	//				msg, err := stream.Receive()
	//				if errors.Is(err, io.EOF) {
	//					break
	//				}
	//				assert.NotEmpty(t, msg.GetSentence())
	//				receivedValues = append(receivedValues, msg.GetSentence())
	//			}
	//			return stream.CloseResponse()
	//		})
	//		require.NoError(t, grp.Wait())
	//		assert.Equal(t, len(receivedValues), len(sendValues))
	//	}
	//})
	//
	//t.Run("introduce", func(t *testing.T) {
	//	total := 0
	//	for _, client := range clients {
	//		request := connect.NewRequest(&elizav1.IntroduceRequest{
	//			Name: "Ringo",
	//		})
	//		stream, err := client.Introduce(context.Background(), request)
	//		require.NoError(t, err)
	//		for stream.Receive() {
	//			total++
	//		}
	//		assert.NoError(t, stream.Err())
	//		assert.NoError(t, stream.Close())
	//		assert.Positive(t, total)
	//	}
	//})
}
