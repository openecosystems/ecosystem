// Code generated by protoc-gen-platform go/cli-methods. DO NOT EDIT.
// source: platform/iam/v2alpha/account_authority.proto

package iamv2alphapbcmd

import (
	"connectrpc.com/connect"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/golang/protobuf/jsonpb"
	"libs/public/go/sdk/gen/iam/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"libs/public/go/protobuf/gen/platform/iam/v2alpha"
)

var (
	createAccountAuthorityRequest      string
	createAccountAuthorityFieldMask    string
	createAccountAuthorityValidateOnly bool
)

var CreateAccountAuthorityV2AlphaCmd = &cobra.Command{
	Use:   "createAccountAuthority",
	Short: ``,
	Long: ` Method to CreateAccountAuthority to events based on scopes
`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Debug("Calling createAccountAuthority accountAuthority")

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := iamv2alphapb.CreateAccountAuthorityRequest{}
		log.Debug(_r.String())
		err = jsonpb.Unmarshal(strings.NewReader(_request), &_r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createAccountAuthorityFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createAccountAuthorityValidateOnly

		request := connect.NewRequest[iamv2alphapb.CreateAccountAuthorityRequest](&_r)
		client := *iamv2alphapbsdk.NewAccountAuthorityServiceSpecClient(sdkv2alphalib.Config, sdkv2alphalib.Config.Platform.Endpoint, connect.WithSendGzip(), connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))
		response, err := client.CreateAccountAuthority(context.Background(), request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

func init() {
	CreateAccountAuthorityV2AlphaCmd.PersistentFlags().StringVarP(&createAccountAuthorityRequest, "request", "r", "{}", "Request for api call")
	CreateAccountAuthorityV2AlphaCmd.PersistentFlags().BoolVar(&createAccountAuthorityValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateAccountAuthorityV2AlphaCmd.PersistentFlags().StringVarP(&createAccountAuthorityFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
