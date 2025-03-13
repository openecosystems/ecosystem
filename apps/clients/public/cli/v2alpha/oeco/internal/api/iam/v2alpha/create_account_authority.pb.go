package iamv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	"os"
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	nebulav1ca "libs/partner/go/nebula/v1/ca"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	createAccountAuthorityRequest      string
	createAccountAuthorityFieldMask    string
	createAccountAuthorityValidateOnly bool
)

// CreateAccountAuthorityV2AlphaCmd defines a Cobra command to create an Account Authority for managing ecosystem partners.
var CreateAccountAuthorityV2AlphaCmd = &cobra.Command{
	Use:   "create",
	Short: `Method to create an Account Authority to manage the ecosystem partners`,
	Long:  `[ Create an Account Authority ]`,
	Run: func(cmd *cobra.Command, _ []string) {
		log.Debug("Calling createAccountAuthority ")

		nca := *nebulav1ca.Bound

		_request, err := cmd.Flags().GetString("request")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _request == "" {
			_request = "{}"
		}

		request := iamv2alphapb.CreateAccountAuthorityRequest{}
		err = protojson.Unmarshal([]byte(_request), &request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sdkv2alphalib.Overrides.FieldMask = createAccountAuthorityFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createAccountAuthorityValidateOnly

		ca, err := nca.GetAccountAuthority(context.Background(), &request)
		if err != nil {
			return
		}

		response := iamv2alphapb.CreateAccountAuthorityResponse{
			//SpecContext: &specv2pb.SpecResponseContext{
			//	ResponseValidation: &typev2pb.ResponseValidation{
			//		ValidateOnly: request.Spec.Context.Validation.ValidateOnly,
			//	},
			//	ResponseMask: &typev2pb.ResponseMask{
			//		FieldMask:  request.Spec.SpecData.FieldMask,
			//		PolicyMask: nil,
			//	},
			//	OrganizationSlug: request.Spec.Context.OrganizationSlug,
			//	WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
			//	WorkspaceJan:     request.Spec.Context.WorkspaceJan,
			//},
			AccountAuthority: ca,
		}

		provider, err5 := sdkv2alphalib.NewCredentialProvider()
		if err5 != nil {
			return
		}

		err6 := provider.SaveCredential(ca.Credential)
		if err6 != nil {
			return
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
