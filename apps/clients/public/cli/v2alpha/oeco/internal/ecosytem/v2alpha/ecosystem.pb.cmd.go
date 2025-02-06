package ecosystemv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// EcosystemServiceServiceCmd represents the root command for managing ecosystem-related operations.
var EcosystemServiceServiceCmd = &cobra.Command{
	Use:   "ecosystem",
	Short: `Interact with ecosystems`,
	Long:  ``,

	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

func init() {
	EcosystemServiceServiceCmd.AddCommand(CreateEcosystemV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(ListEcosystemsV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(GetEcosystemV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(UpdateEcosystemV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(DeleteEcosystemV2AlphaCmd)
}
