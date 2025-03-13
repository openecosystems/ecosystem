package enclavev2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// EnclaveServiceServiceCmd represents the root command for managing enclave-related operations.
var EnclaveServiceServiceCmd = &cobra.Command{
	Use:   "enclave",
	Short: `Securely manage enclave data and certificates`,
	Long:  `Securely manage enclave data and certificates`,

	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

func init() {
	EnclaveServiceServiceCmd.AddCommand(SignEnclaveV2AlphaCmd)

	//EnclaveServiceServiceCmd.AddCommand(FindEnclavesV2AlphaCmd)
	//
	//EnclaveServiceServiceCmd.AddCommand(RemoveEnclaveV2AlphaCmd)
	//
	//EnclaveServiceServiceCmd.AddCommand(AttestEnclaveV2AlphaCmd)
}
