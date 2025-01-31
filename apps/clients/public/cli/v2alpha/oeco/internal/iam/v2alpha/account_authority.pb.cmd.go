package iamv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// AccountAuthorityServiceServiceCmd represents the root command for managing account authority-related operations.
var AccountAuthorityServiceServiceCmd = &cobra.Command{
	Use:   "accountAuthority",
	Short: `When creating a new ecosystem, you must create an account authority.`,
	Long:  `You have complete control over your keys`,
	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

// init registers the CreateAccountAuthorityV2AlphaCmd as a sub-command of AccountAuthorityServiceServiceCmd.
func init() {
	AccountAuthorityServiceServiceCmd.AddCommand(CreateAccountAuthorityV2AlphaCmd)
}
