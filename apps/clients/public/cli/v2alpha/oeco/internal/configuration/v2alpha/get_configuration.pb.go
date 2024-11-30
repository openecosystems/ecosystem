package configurationv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	"libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	"libs/public/go/sdk/gen/configuration/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"path/filepath"
	"strings"

	"connectrpc.com/connect"

	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"

	"github.com/spf13/cobra"
)

var (
	getConfigurationRequest      string
	getConfigurationFieldMask    string
	getConfigurationValidateOnly bool
)

var GetConfigurationV2AlphaCmd = &cobra.Command{
	Use:   "getConfiguration",
	Short: ``,
	Long: `
 Get workspace location
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Calling getConfiguration configuration")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := configurationv2alphapb.GetConfigurationRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = getConfigurationFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = getConfigurationValidateOnly

		request := connect.NewRequest[configurationv2alphapb.GetConfigurationRequest](&_r)
		client := *configurationv2alphapbsdk.NewConfigurationServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.GetConfiguration(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		j, _ := json.MarshalIndent(&response, "", "    ")

		fs := sdkv2alphalib.NewFileSystem()
		err = fs.WriteFile(filepath.Join(fs.ConfigurationDirectory, "configuration.json"), j, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(j))
	},
}

func init() {
	GetConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&getConfigurationRequest, "request", "r", "{}", "Request for api call")
	GetConfigurationV2AlphaCmd.PersistentFlags().BoolVar(&getConfigurationValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	GetConfigurationV2AlphaCmd.PersistentFlags().StringVarP(&getConfigurationFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
