// Code generated by protoc-gen-platform go/cli-methods. DO NOT EDIT.
// source: platform/communication/v1beta/preference_center.proto

package communicationv1betapbcmd

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"
	"libs/public/go/sdk/gen/communication/v1beta"
	"libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"libs/public/go/protobuf/gen/platform/communication/v1beta"
)

var (
	getPreferenceRequest      string
	getPreferenceFieldMask    string
	getPreferenceValidateOnly bool
)

var GetPreferenceV1BetaCmd = &cobra.Command{
	Use:   "getPreference",
	Short: ``,
	Long: ` Get Communication Preferences
`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Debug("Calling getPreference preferenceCenter")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := communicationv1betapb.GetPreferenceRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = getPreferenceFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = getPreferenceValidateOnly

		request := connect.NewRequest[communicationv1betapb.GetPreferenceRequest](&_r)
		client := *communicationv1betapbsdk.NewPreferenceCenterServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.GetPreference(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	GetPreferenceV1BetaCmd.PersistentFlags().StringVarP(&getPreferenceRequest, "request", "r", "{}", "Request for api call")
	GetPreferenceV1BetaCmd.PersistentFlags().BoolVar(&getPreferenceValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	GetPreferenceV1BetaCmd.PersistentFlags().StringVarP(&getPreferenceFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
