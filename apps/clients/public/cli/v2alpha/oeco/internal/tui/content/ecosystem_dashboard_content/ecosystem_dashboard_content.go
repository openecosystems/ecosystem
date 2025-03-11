package ecosystemdashboardcontent

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	content "apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
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
		m.InitBase(),
		m.Table.Init(),
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

	m.BaseModel, baseCmd = m.DashboardBaseModel.UpdateBase(msg)

	switch message := msg.(type) {
	case tea.KeyMsg:

		if m.DashboardBaseModel.IsPacketCaptureFocused() {
			_ = message
		}

		if m.DashboardBaseModel.IsSearchFocused() {
			switch {
			case message.Type == tea.KeyCtrlC, message.Type == tea.KeyEsc:
				m.SearchBar.SetValue(m.SearchValue)
				blinkCmd := m.SetIsSearching(false)
				return m, blinkCmd

				// case message.Type == tea.KeyEnter:
				//	m.SearchValue = m.SearchBar.Value()
				//	m.SetIsSearching(false)
				//	m.ResetRows()
				//	return m, tea.Batch(m.FetchNextPageSectionRows()...)
			}

			break
		}

		if m.IsPromptConfirmationFocused() {
			switch {
			case message.Type == tea.KeyCtrlC, message.Type == tea.KeyEsc:
				m.PromptConfirmationBox.Reset()
				baseCmd = m.SetIsPromptConfirmationShown(false)
				return m, baseCmd

			case message.Type == tea.KeyEnter:
				input := m.PromptConfirmationBox.Value()
				action := m.GetPromptConfirmationAction()
				if input == "Y" || input == "y" {
					switch action {
					case "details":
						baseCmd = m.details()
					}
				}

				m.PromptConfirmationBox.Reset()
				blinkCmd := m.SetIsPromptConfirmationShown(false)

				return m, tea.Batch(baseCmd, blinkCmd)
			}
			break
		}
	}

	m.DashboardBaseModel.Viewport.SetContent(m.Table.View())
	v, c := m.DashboardBaseModel.Viewport.Update(msg)
	m.DashboardBaseModel.Viewport = &v
	viewportCmd = c

	search, searchCmd := m.SearchBar.Update(msg)
	m.SearchBar = search

	prompt, promptCmd := m.PromptConfirmationBox.Update(msg)
	m.PromptConfirmationBox = prompt

	_, tableCmd := m.Table.Update(msg)

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
	//s := "\nReal-Time Mesh Packet Analyzer\n" + "---------------------------------\n"
	//
	//// Show captured packets
	//for i, p := range m.Packets {
	//	s += fmt.Sprintf("%2d: %s -> %s | %s | %d bytes\n",
	//		i+1, p.SrcIP, p.DstIP, p.Protocol, p.Length)
	//}
	//
	//return s
	return m.ViewBase()
}
