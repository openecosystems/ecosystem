package configurationv2alphapbint

import (
	"github.com/spf13/cobra"
	"os"
)

var SystemCmd = &cobra.Command{
	Use:   "configuration",
	Short: `configuration system`,
	Long:  `Interact with the configuration system`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	},
}

func init() {
	SystemCmd.AddCommand(ConfigurationServiceServiceCmd)
}
