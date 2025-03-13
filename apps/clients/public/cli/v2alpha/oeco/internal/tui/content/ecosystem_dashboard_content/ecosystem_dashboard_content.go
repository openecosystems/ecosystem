package ecosystemdashboardcontent

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	content "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
)

// PageType defines the type of page as "dashboard" to differentiate it from other potential page types in the system.
const (
	PageType = "dashboard"
)

// Model represents a structure combining a dashboard base model with additional issue data for ecosystems.
type Model struct {
	*content.DashboardBaseModel
}

// NewModel initializes and returns a new Model with the provided context.
func NewModel(ctx *context.ProgramContext) *Model {
	viewportModel := viewport.New(ctx.MainContentBodyWidth, ctx.MainContentBodyHeight)
	model := &Model{}
	model.DashboardBaseModel = content.NewDashboardBaseModel(
		ctx,
		&content.NewDashboardBaseOptions{
			NewBaseOptions: &content.NewBaseOptions{
				Viewport: &viewportModel,
			},
			ID:            0,
			Singular:      "packet",
			Plural:        "packets",
			Type:          PageType,
			SearchFilters: "",
			LastUpdated:   time.Now(),
			CreatedAt:     time.Now(),
		},
	)

	return model
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds,
		m.DashboardBaseModel.InitBase(),
		m.DashboardBaseModel.SearchBar.Init(),
		m.DashboardBaseModel.Table.Init(),
		m.DashboardBaseModel.PromptConfirmationBox.Init(),
	)

	return tea.Batch(cmds...)
}

// Update processes a message, updates the model state, and returns the updated model along with a command to be executed.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		baseCmd     tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.DashboardBaseModel, baseCmd = m.DashboardBaseModel.UpdateDashboardBase(msg)

	switch message := msg.(type) {
	case tea.KeyMsg:

		if m.DashboardBaseModel.IsPacketCaptureFocused() {
			_ = message
			//switch {
			//case message.Type == tea.KeyLeft, message.Type == tea.KeyRight:
			//	m.SetIsPacketCapturing(false)
			//}
		}

		if m.DashboardBaseModel.IsSearchFocused() {
			switch {
			case message.Type == tea.KeyCtrlC, message.Type == tea.KeyEsc:
				m.DashboardBaseModel.SearchBar.SetValue(m.SearchValue)
				blinkCmd := m.DashboardBaseModel.SetIsSearching(false)
				return m, blinkCmd

				// case message.Type == tea.KeyEnter:
				//	m.SearchValue = m.SearchBar.Value()
				//	m.SetIsSearching(false)
				//	m.ResetRows()
				//	return m, tea.Batch(m.FetchNextPageSectionRows()...)
			}

			break
		}

		if m.DashboardBaseModel.IsPromptConfirmationFocused() {
			switch {
			case message.Type == tea.KeyCtrlC, message.Type == tea.KeyEsc:
				m.DashboardBaseModel.PromptConfirmationBox.Reset()
				baseCmd = m.DashboardBaseModel.SetIsPromptConfirmationShown(false)
				return m, baseCmd

			case message.Type == tea.KeyEnter:
				input := m.DashboardBaseModel.PromptConfirmationBox.Value()
				action := m.DashboardBaseModel.GetPromptConfirmationAction()
				if input == "Y" || input == "y" {
					switch action {
					case "details":
						baseCmd = m.details()
					}
				}

				m.DashboardBaseModel.PromptConfirmationBox.Reset()
				blinkCmd := m.DashboardBaseModel.SetIsPromptConfirmationShown(false)

				return m, tea.Batch(baseCmd, blinkCmd)
			}
			break
		}
	}

	m.DashboardBaseModel.Viewport.SetContent(m.DashboardBaseModel.DashboardBaseView())
	v, c := m.DashboardBaseModel.Viewport.Update(msg)
	m.DashboardBaseModel.Viewport = &v
	viewportCmd = c

	_, searchCmd := m.DashboardBaseModel.SearchBar.Update(msg)
	_, promptCmd := m.DashboardBaseModel.PromptConfirmationBox.Update(msg)
	_, tableCmd := m.DashboardBaseModel.Table.Update(msg)

	cmds = append(cmds,
		baseCmd,
		viewportCmd,
		searchCmd,
		promptCmd,
		tableCmd,
	)

	return m, tea.Batch(cmds...)
}

// View returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m *Model) View() string {
	return m.ViewBase()
}
