package sdkv2alphalib

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/emptypb"
	v2alpha "libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	"net/http"
	"reflect"
	"strings"
)

const (
	ConfigurationServiceCreateConfigurationProcedure = "/platform.configuration.v2alpha.ConfigurationService/CreateConfiguration"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	configurationServiceServiceDescriptor                   = v2alpha.File_platform_configuration_v2alpha_configuration_proto.Services().ByName("ConfigurationService")
	configurationServiceCreateConfigurationMethodDescriptor = configurationServiceServiceDescriptor.Methods().ByName("CreateConfiguration")
)

type DynamicConnectorClient interface {
	DynamicUnary(context.Context, *connect.Request[v2alpha.CreateConfigurationRequest]) (*connect.Response[v2alpha.CreateConfigurationResponse], error)
}

func NewDynamicConnectorClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DynamicConnectorClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dynamicConnectorClient{
		createConfiguration: connect.NewClient[v2alpha.CreateConfigurationRequest, v2alpha.CreateConfigurationResponse](
			httpClient,
			baseURL+ConfigurationServiceCreateConfigurationProcedure,
			connect.WithSchema(configurationServiceCreateConfigurationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

type dynamicConnectorClient struct {
	createConfiguration *connect.Client[v2alpha.CreateConfigurationRequest, v2alpha.CreateConfigurationResponse]
}

func (c *dynamicConnectorClient) DynamicUnary(ctx context.Context, req *connect.Request[v2alpha.CreateConfigurationRequest]) (*connect.Response[v2alpha.CreateConfigurationResponse], error) {
	return c.createConfiguration.CallUnary(ctx, req)
}

type DynamicConnectorHandler interface {
	DynamicUnary(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v2alpha.CreateConfigurationResponse], error)
}

func NewDynamicConnectorHandler(c *Connector, opts ...connect.HandlerOption) http.Handler {

	_c := *c
	_c.MethodsByPath()
	mpb := _c.MethodsByPath()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if method, ok := mpb[r.URL.Path]; ok {

			//i := dynamicpb.NewMessage(method.Input).Type()
			//o := dynamicpb.NewMessage(method.Output).Type()
			g := func(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {

				//example := v2alpha.CreateConfigurationResponse{
				//	SpecContext: &specv2pb.SpecResponseContext{
				//		ResponseValidation: &typev2pb.ResponseValidation{
				//			ValidateOnly: true,
				//		},
				//		OrganizationSlug: "hello",
				//		WorkspaceSlug:    "world",
				//		WorkspaceJan:     1,
				//		RoutineId:        "123",
				//	},
				//	Configuration: &v2alpha.Configuration{
				//		Id:               "123",
				//		OrganizationSlug: "hello",
				//		WorkspaceSlug:    "world",
				//	},
				//}

				//a := typev2pb.SpecErrorDetail{
				//	CorrelationId: "123",
				//	UserMessage:   "hello world",
				//}

				//marshal, err := protopb.Marshal(&a)
				//if err != nil {
				//	return nil, err
				//}

				// Convert the struct to a dynamicpb.Message
				//message, err := ConvertStructToDynamicMessage(example, method.Output)
				//if err != nil {
				//	fmt.Println("Failed to convert struct to dynamic message: ", err)
				//}

				//val := reflect.ValueOf(example)
				//typ := val.Type()

				return nil, errors.New("error message from test")

				//return connect.NewResponse[bytes.Buffer](bytes.NewBuffer(marshal)), nil
			}

			_method := *method

			connectorDynamicHandler := connect.NewUnaryHandler(
				r.URL.Path,
				g, //c.DynamicUnary,
				connect.WithSchema(_method.Schema()),
				connect.WithHandlerOptions(opts...),
			)
			connectorDynamicHandler.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}

	})
}

// ConvertStructToDynamicMessage converts a struct to a dynamicpb.Message based on a MessageDescriptor
func ConvertStructToDynamicMessage(input interface{}, messageDescriptor protoreflect.MessageDescriptor) (*dynamicpb.Message, error) {
	// Create an empty dynamic message from the descriptor
	message := dynamicpb.NewMessage(messageDescriptor)

	// Get the value and type of the input struct
	val := reflect.ValueOf(input)
	typ := val.Type()

	//for i := 0; i < messageDescriptor.Fields().Len(); i++ {
	//	field := messageDescriptor.Fields().Get(i)
	//
	//	fmt.Println(val.FieldByNameFunc(func(name string) bool {
	//
	//		fmt.Println(name)
	//		return true
	//	}))
	//	fmt.Println("\n")
	//	fmt.Println(val.Field(i).String())
	//	fmt.Println(field.Name())
	//
	//	//n.Set(field, protoreflect.ValueOfMessage(val.))
	//
	//}

	// Loop through each field in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		// Get the corresponding field descriptor in the message by name
		messageField := messageDescriptor.Fields().ByName(protoreflect.Name(fieldName))
		if messageField == nil {
			fmt.Printf("No matching field for %s in the message descriptor\n", fieldName)
			continue
		}

		// Set the field in the dynamic message based on the kind of the field
		switch messageField.Kind() {
		case protoreflect.StringKind:
			if field.Kind() == reflect.String {
				message.Set(messageField, protoreflect.ValueOfString(field.String()))
			}

		//case protoreflect.Int32Kind:
		//	if field.Kind() == reflect.Int32 || field.Kind() == reflect.Int {
		//		message.Set(messageField, protoreflect.ValueOfInt32(int32(field.Int())))
		//	}
		//
		//case protoreflect.FloatKind:
		//	if field.Kind() == reflect.Float32 {
		//		message.Set(messageField, protoreflect.ValueOfFloat32(float32(field.Float())))
		//	}
		//
		//case protoreflect.BoolKind:
		//	if field.Kind() == reflect.Bool {
		//		message.Set(messageField, protoreflect.ValueOfBool(field.Bool()))
		//	}
		//
		//case protoreflect.MessageKind:
		//	// Handle nested struct by recursively converting it to a dynamic message
		//	nestedMessageDescriptor := messageField.Message()
		//	nestedMessage, err := ConvertStructToDynamicMessage(field.Interface(), nestedMessageDescriptor)
		//	if err != nil {
		//		return nil, fmt.Errorf("failed to convert nested message for field %s: %w", fieldName, err)
		//	}
		//	message.Set(messageField, protoreflect.ValueOfMessage(nestedMessage))
		//
		default:
			fmt.Printf("Unsupported field type for %s\n", fieldName)
		}
	}

	return message, nil
}
