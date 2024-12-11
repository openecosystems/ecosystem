package cryptographyv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

var SystemCmd = &cobra.Command{
	Use:   "cryptography",
	Short: `cryptography system`,
	Long:  `Interact with the cryptography system`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

func init() {
	SystemCmd.AddCommand(CertificateAuthorityServiceServiceCmd)
	SystemCmd.AddCommand(CertificateServiceServiceCmd)
}
