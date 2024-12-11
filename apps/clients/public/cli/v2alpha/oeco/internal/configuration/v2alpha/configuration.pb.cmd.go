package configurationv2alphapbint

import (
	configurationv2alphapbcmd "libs/public/go/cli/v2alpha/gen/platform/configuration/v2alpha"
	"os"

	"github.com/spf13/cobra"
)

var ConfigurationServiceServiceCmd = &cobra.Command{
	Use:   "configuration",
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
	ConfigurationServiceServiceCmd.AddCommand(CreateConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(configurationv2alphapbcmd.ListConfigurationsV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(GetConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(configurationv2alphapbcmd.UpdateConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(configurationv2alphapbcmd.LoadConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(DeleteConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(configurationv2alphapbcmd.PublishConfigurationV2AlphaCmd)

	ConfigurationServiceServiceCmd.AddCommand(configurationv2alphapbcmd.ArchiveConfigurationV2AlphaCmd)
}
