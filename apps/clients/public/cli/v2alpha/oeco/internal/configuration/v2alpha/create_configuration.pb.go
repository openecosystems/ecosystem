package configurationv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	"libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	cryptographyv2alphapb "libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	"libs/public/go/sdk/gen/configuration/v2alpha"
	cryptographyv2alphapbsdk "libs/public/go/sdk/gen/cryptography/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"connectrpc.com/connect"

	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"

	"github.com/spf13/cobra"
)

var (
	createConfigurationRequest      string
	createConfigurationFieldMask    string
	createConfigurationValidateOnly bool
)

var CreateConfigurationV2AlphaCmd = &cobra.Command{
	Use:   "createConfiguration",
	Short: ``,
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
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
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
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

func init() {
	CreateConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&createConfigurationRequest, "request", "r", "{}", "Request for api call")
	CreateConfigurationV2AlphaCmd.PersistentFlags().BoolVar(&createConfigurationValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&createConfigurationFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
