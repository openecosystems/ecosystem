package sidebar

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	tea "github.com/charmbracelet/bubbletea"
)

// DashboardBaseModel is a specialized model that embeds BaseModel to manage dashboard-specific UI state and behaviors.
type DashboardBaseModel struct {
	*BaseModel
}

// NewDashboardBaseOptions represents configuration options specifically for initializing a DashboardBaseModel.
// It embeds NewBaseOptions for broader base model configuration functionality.
type NewDashboardBaseOptions struct {
	*NewBaseOptions
}

// NewDashboardBaseModel initializes and returns a new DashboardBaseModel using the provided ProgramContext and options.
func NewDashboardBaseModel(ctx *context.ProgramContext, options *NewDashboardBaseOptions) *DashboardBaseModel {
	m := &DashboardBaseModel{
		BaseModel: NewBaseModel(ctx, options.NewBaseOptions),
	}

	return m
}

// UpdateDashboardBase updates the DashboardBaseModel by delegating to UpdateBase and batching related commands.
func (m *DashboardBaseModel) UpdateDashboardBase(msg tea.Msg) (*DashboardBaseModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)
	cmds = append(
		cmds,
		cmd,
	)

	return m, tea.Batch(cmds...)
}

// ViewDashboardBase renders the dashboard base view by delegating to the ViewBase method of the embedded BaseModel.
func (m *DashboardBaseModel) ViewDashboardBase() string {
	return m.ViewBase()

	// return m.Viewport.View()
	// s := m.ViewDebug()
	// return s.String()
}
