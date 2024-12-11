package cryptographyv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

var CertificateServiceServiceCmd = &cobra.Command{
	Use:   "certificate",
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
	CertificateServiceServiceCmd.AddCommand(CreateCertificateV2AlphaCmd)

	CertificateServiceServiceCmd.AddCommand(CreateAndSignCertificateV2AlphaCmd)
}
