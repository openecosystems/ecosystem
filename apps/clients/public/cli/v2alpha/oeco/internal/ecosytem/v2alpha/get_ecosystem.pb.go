package ecosystemv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"connectrpc.com/connect"
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	cliv2alphalib "libs/public/go/cli/v2alpha"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapbsdk "libs/public/go/sdk/gen/ecosystem/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

var (
	getEcosystemRequest      string
	getEcosystemFieldMask    string
	getEcosystemValidateOnly bool
)

// GetEcosystemV2AlphaCmd represents the command to retrieve details of an ecosystem in the Ecosystem V2 Alpha service.
var GetEcosystemV2AlphaCmd = &cobra.Command{
	Use:   "get",
	Short: ``,
	Long:  `[]`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling getEcosystem ecosystem")
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*cliv2alphalib.Configuration)

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := ecosystemv2alphapb.GetEcosystemRequest{}
		err = protojson.Unmarshal([]byte(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = getEcosystemFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = getEcosystemValidateOnly

		request := connect.NewRequest[ecosystemv2alphapb.GetEcosystemRequest](&_r)
		// Add GZIP Support: connect.WithSendGzip(),
		url := "https://" + settings.Platform.Mesh.Endpoint
		if settings.Platform.Insecure {
			url = "http://" + settings.Platform.Mesh.Endpoint
		}
		client := *ecosystemv2alphapbsdk.NewEcosystemServiceSpecClient(&settings.Platform, url, connect.WithInterceptors(cliv2alphalib.NewCLIInterceptor(settings, sdkv2alphalib.Overrides)))

		response, err := client.GetEcosystem(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	GetEcosystemV2AlphaCmd.PersistentFlags().StringVarP(&getEcosystemRequest, "request", "r", "{}", "Request for api call")
	GetEcosystemV2AlphaCmd.PersistentFlags().BoolVar(&getEcosystemValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	GetEcosystemV2AlphaCmd.PersistentFlags().StringVarP(&getEcosystemFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
