package configurationv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	configurationv2alphapb "libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	cryptographyv2alphapb "libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	configurationv2alphapbsdk "libs/public/go/sdk/gen/configuration/v2alpha"
	cryptographyv2alphapbsdk "libs/public/go/sdk/gen/cryptography/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"os"

	"connectrpc.com/connect"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// createConfigurationRequest represents the request payload for creating a configuration.
// createConfigurationFieldMask specifies the fields to be updated in the configuration.
// createConfigurationValidateOnly indicates if the operation should validate the request without applying changes.
var (
	createConfigurationRequest      string
	createConfigurationFieldMask    string
	createConfigurationValidateOnly bool
)

// CreateConfigurationV2AlphaCmd is a command to create a new configuration using the provided request parameters.
var CreateConfigurationV2AlphaCmd = &cobra.Command{
	Use:   "createConfiguration",
	Short: ``,
	Long: `
`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling createConfiguration configuration")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := configurationv2alphapb.CreateConfigurationRequest{}
		log.Debug(_r.String())
		err = protojson.Unmarshal([]byte(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createConfigurationFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createConfigurationValidateOnly

		// parse Field mask
		// get Policy from configuration for entity
		// get only fields from configuration for entity
		// return encrypted fields

		specClient := cryptographyv2alphapbsdk.NewEncryptionServiceSpecClient(sdkv2alphalib.Config, "http://localhost:6487")
		c := *specClient

		_er := cryptographyv2alphapb.EncryptRequest{
			PlainText:      []byte("hello world text needing encryption"),
			AssociatedData: []byte("metadata for hello world"),
		}
		er := connect.NewRequest[cryptographyv2alphapb.EncryptRequest](&_er)
		encrypt, err := c.Encrypt(cmd.Context(), er)
		if err != nil {
			fmt.Println("encrypt error: ", err)
			return
		}

		fmt.Println(encrypt.Msg.GetErr())

		request := connect.NewRequest[configurationv2alphapb.CreateConfigurationRequest](&_r)
		client := *configurationv2alphapbsdk.NewConfigurationServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.CreateConfiguration(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

// init initializes persistent flags for the CreateConfigurationV2AlphaCmd command.
// It defines CLI options for request payload, validation-only mode, and field mask for response fields.
func init() {
	CreateConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&createConfigurationRequest, "request", "r", "{}", "Request for api call")
	CreateConfigurationV2AlphaCmd.PersistentFlags().BoolVar(&createConfigurationValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&createConfigurationFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
