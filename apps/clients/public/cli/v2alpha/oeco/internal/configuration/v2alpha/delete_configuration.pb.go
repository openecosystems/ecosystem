package configurationv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	configurationv2alphapb "libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	configurationv2alphapbsdk "libs/public/go/sdk/gen/configuration/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"os"

	"connectrpc.com/connect"

	"github.com/apex/log"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/spf13/cobra"
)

// deleteConfigurationRequest holds the request identifier for the delete configuration action.
// deleteConfigurationFieldMask specifies the fields to be included in the delete configuration operation.
// deleteConfigurationValidateOnly indicates if the operation should only validate the request without executing it.
var (
	deleteConfigurationRequest      string
	deleteConfigurationFieldMask    string
	deleteConfigurationValidateOnly bool
)

// DeleteConfigurationV2AlphaCmd is a Cobra command for deleting a configuration in the V2 Alpha API service.
var DeleteConfigurationV2AlphaCmd = &cobra.Command{
	Use:   "deleteConfiguration",
	Short: ``,
	Long: `
`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling deleteConfiguration configuration")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := configurationv2alphapb.DeleteConfigurationRequest{}
		log.Debug(_r.String())
		err = protojson.Unmarshal([]byte(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = deleteConfigurationFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = deleteConfigurationValidateOnly

		request := connect.NewRequest[configurationv2alphapb.DeleteConfigurationRequest](&_r)
		client := *configurationv2alphapbsdk.NewConfigurationServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.DeleteConfiguration(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

// init initializes the persistent flags for the DeleteConfigurationV2AlphaCmd, including request, validate-only, and field-mask options.
func init() {
	DeleteConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&deleteConfigurationRequest, "request", "r", "{}", "Request for api call")
	DeleteConfigurationV2AlphaCmd.PersistentFlags().BoolVar(&deleteConfigurationValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	DeleteConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&deleteConfigurationFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
