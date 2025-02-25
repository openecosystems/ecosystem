package ecosystemv2alphapbint

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	tasks "apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks"
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
	// EcosystemServiceServiceCmd.AddCommand(CreateEcosystemV2AlphaCmd)
	EcosystemServiceServiceCmd.AddCommand(Cmd)

	EcosystemServiceServiceCmd.AddCommand(ListEcosystemsV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(GetEcosystemV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(UpdateEcosystemV2AlphaCmd)

	EcosystemServiceServiceCmd.AddCommand(DeleteEcosystemV2AlphaCmd)
}

// gracefulShutdown recovers from any panic that occurred and logs the recovery message before quitting the tea program.
func gracefulShutdown() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}

	tasks.Close()
	_ = tea.Quit()
}
