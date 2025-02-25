package tasks

import (
	tea "github.com/charmbracelet/bubbletea"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// CmdMsg represents a command message structure used for encapsulating command operations or instructions in the system.
type CmdMsg struct{}

// ExecuteCMD processes the given ProgramContext and error, returning a command message encapsulated as tea.Msg.
func ExecuteCMD(_ *context.ProgramContext, _ error) tea.Msg {
	return CmdMsg{}
}
