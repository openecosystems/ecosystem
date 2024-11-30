package configurationv2alphapbint

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"
	"libs/public/go/sdk/gen/configuration/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"libs/public/go/protobuf/gen/platform/configuration/v2alpha"
)

var (
	deleteConfigurationRequest      string
	deleteConfigurationFieldMask    string
	deleteConfigurationValidateOnly bool
)

var DeleteConfigurationV2AlphaCmd = &cobra.Command{
	Use:   "deleteConfiguration",
	Short: ``,
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {

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
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
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

func init() {
	DeleteConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&deleteConfigurationRequest, "request", "r", "{}", "Request for api call")
	DeleteConfigurationV2AlphaCmd.PersistentFlags().BoolVar(&deleteConfigurationValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	DeleteConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&deleteConfigurationFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
