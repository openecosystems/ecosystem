// Code generated by protoc-gen-platform go/cli-methods. DO NOT EDIT.
// source: platform/communication/v1beta/preference_center.proto

package communicationv1betapbcli

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"os"

	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/communication/v1beta"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/communication/v1beta/communicationv1betapbconnect"
)

var (
	createOrUpdatePreferenceRequest      string
	createOrUpdatePreferenceFieldMask    string
	createOrUpdatePreferenceValidateOnly bool
)

var CreateOrUpdatePreferenceV1BetaCmd = &cobra.Command{
	Use:   "createOrUpdatePreference",
	Short: `CreateOrUpdatePreference to our communications`,
	Long:  `[]`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Debug("Calling createOrUpdatePreference preferenceCenter")
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*sdkv2alphalib.CLIConfiguration)

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := communicationv1betapb.CreateOrUpdatePreferenceRequest{}
		err = protojson.Unmarshal([]byte(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createOrUpdatePreferenceFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createOrUpdatePreferenceValidateOnly

		request := connect.NewRequest[communicationv1betapb.CreateOrUpdatePreferenceRequest](&_r)
		// Add GZIP Support: connect.WithSendGzip(),
		url := "https://" + settings.Platform.Mesh.Endpoint
		if settings.Platform.Insecure {
			url = "http://" + settings.Platform.Mesh.Endpoint
		}
		client := *communicationv1betapbconnect.NewPreferenceCenterServiceSpecClient(&settings.Platform, url, connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(settings, sdkv2alphalib.Overrides)))

		response, err := client.CreateOrUpdatePreference(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	CreateOrUpdatePreferenceV1BetaCmd.PersistentFlags().StringVarP(&createOrUpdatePreferenceRequest, "request", "r", "{}", "Request for api call")
	CreateOrUpdatePreferenceV1BetaCmd.PersistentFlags().BoolVar(&createOrUpdatePreferenceValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateOrUpdatePreferenceV1BetaCmd.PersistentFlags().StringVarP(&createOrUpdatePreferenceFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
