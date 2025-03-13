package enclavev2alphapbint

import (
	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	clog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// SignEnclaveV2AlphaCmd represents the command to retrieve details of an ecosystem in the Enclave V2 Alpha service.
var SignEnclaveV2AlphaCmd = &cobra.Command{
	Use:   "sign",
	Short: `Sign a certificate in the Enclave`,
	Long:  `[Sign a certificate in the Enclave]`,
	Run: func(cmd *cobra.Command, _ []string) {
		log := cmd.Context().Value(sdkv2alphalib.LoggerContextKey).(*clog.Logger)
		log.Info("Sign Enclave")
	},
}

func init() {
}
