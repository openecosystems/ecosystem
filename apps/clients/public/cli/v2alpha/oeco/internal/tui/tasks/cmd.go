package tasks

import (
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	tea "github.com/charmbracelet/bubbletea"
)

// CmdMsg represents a command message structure used for encapsulating command operations or instructions in the system.
type CmdMsg struct{}

// Execute processes the given ProgramContext and error, returning a command message encapsulated as tea.Msg.
func Execute(_ *context.ProgramContext, _ error) tea.Msg {
	// System calls create account authority <br/>internally to create a new Account Authority credential: <br/>api.{ecosystem-name}.mesh
	// System calls create account <br/>internally to create a new Edge Service Account credential: <br/>edge.{ecosystem-name}.mesh. <br/>system assigns a reserved ipaddress
	// System calls create account <br/>internally to create a new Local Machine Service Account credential: <br/>{sanitized.os.hostname}.{ecosystem-name}.mesh
	// System calls provision edge <br/>internally to configure Edge: <br/>configurations/edge.{ecosystem-name}.mesh
	// System calls provision ecosystem <br/>internally to configure ecosystem: <br/>configurations/api.{ecosystem-name}.mesh
	// User deploys the Ecosystem following installation guide

	return CmdMsg{}
}
