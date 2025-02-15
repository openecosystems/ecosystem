package apiv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// APIServiceServiceCmd represents the root command for managing api-related operations.
var APIServiceServiceCmd = &cobra.Command{
	Use:   "api",
	Short: `Interact with Open Ecosystem APIs using multiple protocols`,
	Long:  `Interact with Open Ecosystem APIs using multiple protocols such as gRPC, gRPC-Web, REST, and Connect`,

	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}
