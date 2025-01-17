// Code generated by protoc-gen-platform go/cli-methods. DO NOT EDIT.
// source: platform/configuration/v2alpha/configuration.proto

package configurationv2alphapbcmd

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
	updateConfigurationRequest      string
	updateConfigurationFieldMask    string
	updateConfigurationValidateOnly bool
)

var UpdateConfigurationV2AlphaCmd = &cobra.Command{
	Use:   "updateConfiguration",
	Short: ``,
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Debug("Calling updateConfiguration configuration")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := configurationv2alphapb.UpdateConfigurationRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = updateConfigurationFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = updateConfigurationValidateOnly

		request := connect.NewRequest[configurationv2alphapb.UpdateConfigurationRequest](&_r)
		client := *configurationv2alphapbsdk.NewConfigurationServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.UpdateConfiguration(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	UpdateConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&updateConfigurationRequest, "request", "r", "{}", "Request for api call")
	UpdateConfigurationV2AlphaCmd.PersistentFlags().BoolVar(&updateConfigurationValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	UpdateConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&updateConfigurationFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
