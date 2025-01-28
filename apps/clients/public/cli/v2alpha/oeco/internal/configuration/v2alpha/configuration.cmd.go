package configurationv2alphapbint

import (
	"os"

	"github.com/spf13/cobra"
)

// SystemCmd is a Cobra command for interacting with the configuration system, providing help and exiting on execution.
var SystemCmd = &cobra.Command{
	Use:   "configuration",
	Short: `configuration system`,
	Long:  `Interact with the configuration system`,
	Run: func(cmd *cobra.Command, _ []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

// init registers the ConfigurationServiceServiceCmd with the SystemCmd as a sub-command during the initialization phase.
func init() {
	SystemCmd.AddCommand(ConfigurationServiceServiceCmd)
}
