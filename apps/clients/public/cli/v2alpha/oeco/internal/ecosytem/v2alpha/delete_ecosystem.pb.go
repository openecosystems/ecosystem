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
	deleteEcosystemRequest      string
	deleteEcosystemFieldMask    string
	deleteEcosystemValidateOnly bool
)

// DeleteEcosystemV2AlphaCmd defines a Cobra command for deleting an ecosystem in the V2 Alpha API.
var DeleteEcosystemV2AlphaCmd = &cobra.Command{
	Use:   "delete",
	Short: ``,
	Long:  `[]`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling deleteEcosystem ecosystem")
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*cliv2alphalib.Configuration)

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := ecosystemv2alphapb.DeleteEcosystemRequest{}
		err = protojson.Unmarshal([]byte(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = deleteEcosystemFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = deleteEcosystemValidateOnly

		request := connect.NewRequest[ecosystemv2alphapb.DeleteEcosystemRequest](&_r)
		// Add GZIP Support: connect.WithSendGzip(),
		url := "https://" + settings.Platform.Mesh.Endpoint
		if settings.Platform.Insecure {
			url = "http://" + settings.Platform.Mesh.Endpoint
		}
		client := *ecosystemv2alphapbsdk.NewEcosystemServiceSpecClient(&settings.Platform, url, connect.WithInterceptors(cliv2alphalib.NewCLIInterceptor(settings, sdkv2alphalib.Overrides)))

		response, err := client.DeleteEcosystem(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	DeleteEcosystemV2AlphaCmd.PersistentFlags().StringVarP(&deleteEcosystemRequest, "request", "r", "{}", "Request for api call")
	DeleteEcosystemV2AlphaCmd.PersistentFlags().BoolVar(&deleteEcosystemValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	DeleteEcosystemV2AlphaCmd.PersistentFlags().StringVarP(&deleteEcosystemFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
