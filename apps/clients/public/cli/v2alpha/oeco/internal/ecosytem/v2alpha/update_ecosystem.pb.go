package ecosystemv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapbsdk "libs/public/go/sdk/gen/ecosystem/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"os"

	"connectrpc.com/connect"

	cliv2alphalib "libs/public/go/cli/v2alpha"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	updateEcosystemRequest      string
	updateEcosystemFieldMask    string
	updateEcosystemValidateOnly bool
)

// UpdateEcosystemV2AlphaCmd represents a command to update ecosystem details via a CLI interface.
var UpdateEcosystemV2AlphaCmd = &cobra.Command{
	Use:   "update",
	Short: ``,
	Long:  `[]`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling updateEcosystem ecosystem")
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*cliv2alphalib.Configuration)

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := ecosystemv2alphapb.UpdateEcosystemRequest{}
		err = protojson.Unmarshal([]byte(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = updateEcosystemFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = updateEcosystemValidateOnly

		request := connect.NewRequest[ecosystemv2alphapb.UpdateEcosystemRequest](&_r)
		// Add GZIP Support: connect.WithSendGzip(),
		url := "https://" + settings.Platform.Mesh.Endpoint
		if settings.Platform.Insecure {
			url = "http://" + settings.Platform.Mesh.Endpoint
		}
		client := *ecosystemv2alphapbsdk.NewEcosystemServiceSpecClient(&settings.Platform, url, connect.WithInterceptors(cliv2alphalib.NewCLIInterceptor(settings, sdkv2alphalib.Overrides)))

		response, err := client.UpdateEcosystem(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	UpdateEcosystemV2AlphaCmd.PersistentFlags().StringVarP(&updateEcosystemRequest, "request", "r", "{}", "Request for api call")
	UpdateEcosystemV2AlphaCmd.PersistentFlags().BoolVar(&updateEcosystemValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	UpdateEcosystemV2AlphaCmd.PersistentFlags().StringVarP(&updateEcosystemFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
