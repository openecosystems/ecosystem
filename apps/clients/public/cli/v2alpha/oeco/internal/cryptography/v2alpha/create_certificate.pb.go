package cryptographyv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	cryptographyv2alphapb "libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	cryptographyv2alphapbsdk "libs/public/go/sdk/gen/cryptography/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"connectrpc.com/connect"

	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"

	"github.com/spf13/cobra"
)

var (
	createCertificateRequest      string
	createCertificateFieldMask    string
	createCertificateValidateOnly bool
)

var CreateCertificateV2AlphaCmd = &cobra.Command{
	Use:   "createCertificate",
	Short: ``,
	Long: ` Method to CreateCertificate to events based on scopes
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Calling createCertificate certificate")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := cryptographyv2alphapb.CreateCertificateRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createCertificateFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createCertificateValidateOnly

		request := connect.NewRequest[cryptographyv2alphapb.CreateCertificateRequest](&_r)
		client := *cryptographyv2alphapbsdk.NewCertificateServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.CreateCertificate(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	CreateCertificateV2AlphaCmd.PersistentFlags().StringVarP(&createCertificateRequest, "request", "r", "{}", "Request for api call")
	CreateCertificateV2AlphaCmd.PersistentFlags().BoolVar(&createCertificateValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateCertificateV2AlphaCmd.PersistentFlags().StringVarP(&createCertificateFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
