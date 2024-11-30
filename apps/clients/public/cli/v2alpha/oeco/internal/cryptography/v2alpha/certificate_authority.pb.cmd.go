package cryptographyv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

var CertificateAuthorityServiceServiceCmd = &cobra.Command{
	Use:   "certificateAuthority",
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

func init() {
	CertificateAuthorityServiceServiceCmd.AddCommand(CreateCertificateAuthorityV2AlphaCmd)
}
