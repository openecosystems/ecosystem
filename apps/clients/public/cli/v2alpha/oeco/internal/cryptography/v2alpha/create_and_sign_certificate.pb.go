package cryptographyv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	"libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	"libs/public/go/sdk/gen/cryptography/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"connectrpc.com/connect"

	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"

	"github.com/spf13/cobra"
)

var (
	createAndSignCertificateRequest      string
	createAndSignCertificateFieldMask    string
	createAndSignCertificateValidateOnly bool
)

var CreateAndSignCertificateV2AlphaCmd = &cobra.Command{
	Use:   "createAndSignCertificate",
	Short: ``,
	Long: ` Method to CreateAndSignCertificate to events based on scopes
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Calling createAndSignCertificate certificate")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := cryptographyv2alphapb.CreateAndSignCertificateRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createAndSignCertificateFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createAndSignCertificateValidateOnly

		request := connect.NewRequest[cryptographyv2alphapb.CreateAndSignCertificateRequest](&_r)
		client := *cryptographyv2alphapbsdk.NewCertificateServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.CreateAndSignCertificate(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	CreateAndSignCertificateV2AlphaCmd.PersistentFlags().StringVarP(&createAndSignCertificateRequest, "request", "r", "{}", "Request for api call")
	CreateAndSignCertificateV2AlphaCmd.PersistentFlags().BoolVar(&createAndSignCertificateValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateAndSignCertificateV2AlphaCmd.PersistentFlags().StringVarP(&createAndSignCertificateFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
