package cryptographyv2alphapbint

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"libs/partner/go/nebula/v1/ca"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/protobuf/go/protobuf/gen/platform/type/v2"
	"libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"os"
	"path/filepath"
	"strings"
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

		//log := *zaploggerv1.Bound.Logger
		nca := *nebulav1ca.Bound

		var req cryptographyv2alphapb.CreateCertificateAuthorityRequest
		err = protojson.Unmarshal([]byte(_request), &req)
		if err != nil {
			return
		}

		ca, err := nca.GetCertificateAuthority(cmd.Context(), &req)
		if err != nil {
			return
		}

		sdkv2alphalib.Overrides.FieldMask = createCertificateAuthorityFieldMask
		sdkv2alphalib.Overrides.ValidateOnly = createCertificateAuthorityValidateOnly

		fs := sdkv2alphalib.NewFileSystem()
		if err := fs.WriteFile(filepath.Join(fs.CredentialsDirectory, "ca.crt"), ca.CaCert.Content, os.ModePerm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := fs.WriteFile(filepath.Join(fs.CredentialsDirectory, "ca.key"), ca.CaKey.Content, os.ModePerm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := fs.WriteFile(filepath.Join(fs.CredentialsDirectory, "ca.png"), ca.CaQrCode.Content, os.ModePerm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		response := cryptographyv2alphapb.CreateCertificateAuthorityResponse{
			SpecContext: &specv2pb.SpecResponseContext{
				ResponseValidation: &typev2pb.ResponseValidation{
					ValidateOnly: createCertificateAuthorityValidateOnly,
				},
				ResponseMask: &typev2pb.ResponseMask{
					FieldMask: &fieldmaskpb.FieldMask{Paths: strings.Split(createCertificateAuthorityFieldMask, ",")},
				},
				//OrganizationSlug: request.Spec.Context.OrganizationSlug,
				//WorkspaceSlug:    request.Spec.Context.WorkspaceSlug,
				//WorkspaceJan:     request.Spec.Context.WorkspaceJan,
			},
			CertificateAuthority: ca,
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
