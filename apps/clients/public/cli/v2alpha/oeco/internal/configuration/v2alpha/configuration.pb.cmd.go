package configurationv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// ConfigurationServiceServiceCmd is a root command for managing configuration-related operations in the CLI.
var ConfigurationServiceServiceCmd = &cobra.Command{
	Use:   "configuration",
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

// init registers all configuration-related commands to the ConfigurationServiceServiceCmd.
func init() {
	ConfigurationServiceServiceCmd.AddCommand(CreateConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(GetConfigurationV2AlphaCmd)
}
