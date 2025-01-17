// Code generated by protoc-gen-platform go/cli-methods. DO NOT EDIT.
// source: platform/cryptography/v2alpha/certificate_authority.proto

package cryptographyv2alphapbcmd

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"
	"libs/public/go/sdk/gen/cryptography/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
)

var (
	createCertificateAuthorityRequest      string
	createCertificateAuthorityFieldMask    string
	createCertificateAuthorityValidateOnly bool
)

var CreateCertificateAuthorityV2AlphaCmd = &cobra.Command{
	Use:   "createCertificateAuthority",
	Short: ``,
	Long: ` Method to CreateCertificateAuthority to events based on scopes
`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Debug("Calling createCertificateAuthority certificateAuthority")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := cryptographyv2alphapb.CreateCertificateAuthorityRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createCertificateAuthorityFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createCertificateAuthorityValidateOnly

		request := connect.NewRequest[cryptographyv2alphapb.CreateCertificateAuthorityRequest](&_r)
		client := *cryptographyv2alphapbsdk.NewCertificateAuthorityServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.CreateCertificateAuthority(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	CreateCertificateAuthorityV2AlphaCmd.PersistentFlags().StringVarP(&createCertificateAuthorityRequest, "request", "r", "{}", "Request for api call")
	CreateCertificateAuthorityV2AlphaCmd.PersistentFlags().BoolVar(&createCertificateAuthorityValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateCertificateAuthorityV2AlphaCmd.PersistentFlags().StringVarP(&createCertificateAuthorityFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
