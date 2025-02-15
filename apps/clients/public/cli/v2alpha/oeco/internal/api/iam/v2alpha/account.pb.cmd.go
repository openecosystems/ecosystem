package iamv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// AccountServiceServiceCmd is a Cobra command for managing accounts within an ecosystem. It displays help by default.
var AccountServiceServiceCmd = &cobra.Command{
	Use:   "account",
	Short: `Manage your account within an ecosystem`,
	Long:  ``,
	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

// init registers the CreateAccountV2AlphaCmd as a subcommand of AccountServiceServiceCmd.
func init() {
	AccountServiceServiceCmd.AddCommand(CreateAccountV2AlphaCmd)
}
