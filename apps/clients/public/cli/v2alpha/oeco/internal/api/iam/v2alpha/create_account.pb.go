package iamv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	nebulav1ca "libs/partner/go/nebula/v1/ca"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	iamv2alphapbconnect "libs/public/go/protobuf/gen/platform/iam/v2alpha/iamv2alphapbconnect"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// createAccountRequest stores the request data for creating an account.
// createAccountFieldMask defines the fields to be updated or included in the create account operation.
// createAccountValidateOnly determines if the operation should only validate the request without making changes.
var (
	createAccountRequest      string
	createAccountFieldMask    string
	createAccountValidateOnly bool
	// fs                        = sdkv2alphalib.NewFileSystem()
)

// CreateAccountV2AlphaCmd is a Cobra command for creating an account to connect to an ecosystem.
// It generates a PKI certificate and optionally requests signing from an Ecosystem Account Authority.
var CreateAccountV2AlphaCmd = &cobra.Command{
	Use:   "create",
	Short: `Create an Account to connect to an ecosystem`,
	Long: `[ Create an account to connect to an ecosystem.
Facilitates creating a PKI certificate and getting it signed by an Ecosystem Account Authority ]`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling createAccount account")
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*cliv2alphalib.Configuration)

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		_r := iamv2alphapb.CreateAccountRequest{}
		err2 := protojson.Unmarshal([]byte(_request), &_r)
		if err2 != nil {
			fmt.Println(err2)
			os.Exit(1)
		}

		cert, key, err3 := nebulav1ca.Bound.GetPKI(context.Background(), &_r)
		if err3 != nil {
			return
		}

		_r.Cert = cert

		sdkv2alphalib.Overrides.FieldMask = createAccountFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createAccountValidateOnly

		request := connect.NewRequest[iamv2alphapb.CreateAccountRequest](&_r)
		httpClient := http.DefaultClient
		url := "https://" + settings.Platform.Endpoint
		if settings.Platform.Insecure {
			url = "http://" + settings.Platform.Endpoint
		}

		client := iamv2alphapbconnect.NewAccountServiceClient(httpClient, url, connect.WithInterceptors(cliv2alphalib.NewCLIInterceptor(settings, sdkv2alphalib.Overrides)))

		response, err4 := client.CreateAccount(context.Background(), request)
		if err4 != nil {
			fmt.Println(err4)
			os.Exit(1)
		}

		if response.Msg == nil && response.Msg.Account == nil && response.Msg.Account.Credential == nil {
			fmt.Println("internal error parsing credential")
			return
		}

		if response.Msg.SpecContext.EcosystemSlug == "" {
			fmt.Println("ecosystem slug is not set within the response context; internal error")
			return
		}

		provider, err5 := sdkv2alphalib.NewCredentialProvider()
		if err5 != nil {
			return
		}

		response.Msg.Account.Credential.PrivateKey = string(key.Content)

		err6 := provider.SaveCredential(response.Msg.Account.Credential)
		if err6 != nil {
			return
		}

		val, _ := json.MarshalIndent(&response, "", "    ")
		fmt.Println(string(val))
	},
}

// init initializes persistent flags for the CreateAccountV2Alpha command, including request body, validation mode, and field mask.
func init() {
	CreateAccountV2AlphaCmd.PersistentFlags().StringVarP(&createAccountRequest, "request", "r", "{}", "Request for api call")
	CreateAccountV2AlphaCmd.PersistentFlags().BoolVar(&createAccountValidateOnly, "validate-only", false, "Only validate this request without modifying the resource")
	CreateAccountV2AlphaCmd.PersistentFlags().StringVarP(&createAccountFieldMask, "field-mask", "m", "", "Limit the returned response fields")
}
