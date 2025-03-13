package content

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	packet "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/packet"
	prompt "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/prompt"
	search "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/search"
	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	constants "apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	theme "apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	utils "apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
)

// DashboardBaseModel represents a data structure used for managing dashboard-related contexts and configurations.
// SingularForm specifies the singular naming form for the dashboard model.
// PluralForm specifies the plural naming form for the dashboard model.
// Table provides metadata about the table associated with the dashboard.
// Columns defines the collection of column data for the dashboard table.
// TotalCount represents the total number of entries in the dashboard's dataset.
type DashboardBaseModel struct {
	*BaseModel

	ID           int
	SingularForm string
	PluralForm   string
	Type         string

	SearchBar         *search.Model
	IsSearching       bool
	SearchValue       string
	IsSearchSupported bool
	SearchFilters     string

	Table      contract.Table
	TotalCount int

	PromptConfirmationBox     *prompt.Model
	IsPromptConfirmationShown bool
	PromptConfirmationAction  string

	IsPacketCapturing bool
}

// NewDashboardBaseOptions defines options for initializing a dashboard base model, including columns and timestamps.
type NewDashboardBaseOptions struct {
	*NewBaseOptions

	ID       int
	Singular string
	Plural   string
	Type     string

	SearchFilters string

	LastUpdated time.Time
	CreatedAt   time.Time
}

// NewDashboardBaseModel initializes and returns a new DashboardBaseModel with provided context and configuration options.
func NewDashboardBaseModel(ctx *context.ProgramContext, options *NewDashboardBaseOptions) *DashboardBaseModel {
	s := search.NewModel(ctx, search.Options{
		Prefix:       "is:" + options.Type,
		InitialValue: options.SearchFilters,
	})

	m := &DashboardBaseModel{
		BaseModel: NewBaseModel(ctx, options.NewBaseOptions),

		ID: options.ID,

		SearchBar:   &s,
		SearchValue: options.SearchFilters,
		IsSearching: false,

		SingularForm: options.Singular,
		PluralForm:   options.Plural,
		TotalCount:   0,

		PromptConfirmationBox: prompt.NewModel(ctx),
	}

	t := packet.NewModel(
		ctx,
		&packet.NewModelOptions{
			Dimensions:    m.GetDimensions(),
			LastUpdated:   options.LastUpdated,
			CreatedAt:     options.CreatedAt,
			Rows:          nil,
			ItemTypeLabel: m.SingularForm,
			EmptyState: utils.StringPtr(m.Ctx.Styles.Section.EmptyStateStyle.Render(
				fmt.Sprintf(
					"No %s were found that match the given filters",
					m.PluralForm,
				),
			)),
			LoadingMessage: "Loading...",
			IsLoading:      false,
		},
	)

	m.Table = t

	return m
}

// UpdateDashboardBase updates the BaseModel's program context and dimensions, returning the updated BaseModel and a batch of commands.
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

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Logger.Debug("Key pressed", "key", message.String())
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.Search):

			if m.Table != nil {
				cmd = m.SetIsSearching(true)
				return m, cmd
			}
		}
	}

	return m, tea.Batch(cmds...)
}

// GetDimensions calculates and returns the dimensions of the main content area, considering padding and search height.
func (m *DashboardBaseModel) GetDimensions() constants.Dimensions {
	return constants.Dimensions{
		Width:  m.Ctx.MainContentWidth - m.Ctx.Styles.Section.ContainerStyle.GetHorizontalPadding(),
		Height: m.Ctx.MainContentHeight - theme.SearchHeight,
	}
}

// UpdateProgramContext updates the program context for the BaseModel and its associated components: main content and sidebar.
func (m *DashboardBaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	if ctx == nil {
		return
	}

	m.BaseModel.UpdateProgramContext(ctx)

	oldDimensions := m.GetDimensions()
	m.Ctx = ctx
	newDimensions := m.GetDimensions()
	tableDimensions := constants.Dimensions{
		Height: newDimensions.Height - 2,
		Width:  newDimensions.Width,
	}
	m.Table.SetDimensions(tableDimensions)
	m.Table.UpdateProgramContext(ctx)

	if oldDimensions.Height != newDimensions.Height ||
		oldDimensions.Width != newDimensions.Width {
		m.Table.SyncViewPortContent()
		m.SearchBar.UpdateProgramContext(ctx)
	}
}

// GetID returns the ID of the DashboardBaseModel.
func (m *DashboardBaseModel) GetID() int {
	return m.ID
}

// GetType returns the type of the DashboardBaseModel instance as a string.
func (m *DashboardBaseModel) GetType() string {
	return m.Type
}

// GetItemSingularForm retrieves the singular naming form of the dashboard model.
func (m *DashboardBaseModel) GetItemSingularForm() string {
	return m.SingularForm
}

// GetItemPluralForm returns the plural form of the item associated with the DashboardBaseModel instance.
func (m *DashboardBaseModel) GetItemPluralForm() string {
	return m.PluralForm
}

// GetTotalCount retrieves the total count of entries if they are not loading, returning nil during a loading state.
func (m *DashboardBaseModel) GetTotalCount() *int {
	if m.IsLoading() {
		return nil
	}
	c := m.TotalCount
	return &c
}

// IsSearchFocused returns true if the search bar is currently focused; otherwise, it returns false.
func (m *DashboardBaseModel) IsSearchFocused() bool {
	return m.IsSearching
}

// SetIsSearching sets the searching state and updates the SearchBar's focus state. Returns a command if needed.
func (m *DashboardBaseModel) SetIsSearching(val bool) tea.Cmd {
	m.IsSearching = val
	if val {
		m.SearchBar.Focus()
		return m.SearchBar.Init()
	}

	m.SearchBar.Blur()
	return nil
}

// ResetFilters clears and resets the search bar's text input value to the predefined search filters.
func (m *DashboardBaseModel) ResetFilters() {
	m.SearchBar.SetValue(m.SearchFilters)
}

// ResetPageInfo clears the page information, including pagination and other related data, resetting it to an initial state.
func (m *DashboardBaseModel) ResetPageInfo() {
	// m.PageInfo = nil
}

// IsPromptConfirmationFocused checks if the prompt confirmation box is currently in focus.
func (m *DashboardBaseModel) IsPromptConfirmationFocused() bool {
	return m.IsPromptConfirmationShown
}

// SetIsPromptConfirmationShown toggles the visibility of the prompt confirmation and adjusts its focus state accordingly.
func (m *DashboardBaseModel) SetIsPromptConfirmationShown(val bool) tea.Cmd {
	m.IsPromptConfirmationShown = val
	if val {
		m.PromptConfirmationBox.Focus()
		return m.PromptConfirmationBox.Init()
	}

	m.PromptConfirmationBox.Blur()
	return nil
}

// SetPromptConfirmationAction sets the action identifier for prompt confirmation interactions within the dashboard model.
func (m *DashboardBaseModel) SetPromptConfirmationAction(action string) {
	m.PromptConfirmationAction = action
}

// GetPromptConfirmationAction retrieves the current action set for the prompt confirmation in the dashboard model.
func (m *DashboardBaseModel) GetPromptConfirmationAction() string {
	return m.PromptConfirmationAction
}

// SectionMsg represents a message consisting of a section identifier, type, and an internal tea.Msg.
type SectionMsg struct {
	ID          int
	Type        string
	InternalMsg tea.Msg
}

// MakeSectionCmd wraps a given command with additional context, including the section's ID and type, returning a new command.
func (m *DashboardBaseModel) MakeSectionCmd(cmd tea.Cmd) tea.Cmd {
	if cmd == nil {
		return nil
	}

	return func() tea.Msg {
		internalMsg := cmd()
		return SectionMsg{
			ID:          m.ID,
			Type:        m.Type,
			InternalMsg: internalMsg,
		}
	}
}

// GetFilters retrieves the current value of the search bar from the dashboard model.
func (m *DashboardBaseModel) GetFilters() string {
	return m.SearchBar.Value()
}

// GetMainContent generates and returns the main content view based on the table state and context dimensions.
func (m *DashboardBaseModel) GetMainContent() string {
	if m.Table.GetRows() == nil {
		d := m.GetDimensions()
		return lipgloss.Place(
			d.Width,
			d.Height,
			lipgloss.Center,
			lipgloss.Center,

			fmt.Sprintf(
				"%s you can change the search query by pressing %s and submitting it with %s",
				lipgloss.NewStyle().Bold(true).Render(" Tip:"),
				m.Ctx.Styles.Section.KeyStyle.Render("/"),
				m.Ctx.Styles.Section.KeyStyle.Render("Enter"),
			),
		)
	}

	return m.Table.View()
}

// DashboardBaseView renders the dashboard model's main view, combining the search bar and main content into a vertically joined layout.
func (m *DashboardBaseModel) DashboardBaseView() string {
	s := m.SearchBar.View()
	return m.Ctx.Styles.Section.ContainerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			s,
			m.GetMainContent(),
		),
	)
}

// UpdateTotalItemsCount updates the total items count in the associated table model with the provided count value.
func (m *DashboardBaseModel) UpdateTotalItemsCount(count int) {
	m.Table.UpdateTotalItemsCount(count)
}

// IsLoading checks if the associated table in the dashboard model is currently in a loading state.
func (m *DashboardBaseModel) IsLoading() bool {
	return m.Table.IsLoading()
}

// GetPromptConfirmation generates and returns the appropriate confirmation prompt string based on the current context and action.
func (m *DashboardBaseModel) GetPromptConfirmation() string {
	if m.IsPromptConfirmationShown {
		var p string
		switch {
		case m.PromptConfirmationAction == "close" && m.Ctx.Page == config.EcosystemCreatePage:
			p = "Are you sure you want to close this PR? (Y/n) "

		case m.PromptConfirmationAction == "reopen" && m.Ctx.Page == config.EcosystemCreatePage:
			p = "Are you sure you want to reopen this PR? (Y/n) "

		case m.PromptConfirmationAction == "ready" && m.Ctx.Page == config.EcosystemCreatePage:
			p = "Are you sure you want to mark this PR as ready? (Y/n) "

		case m.PromptConfirmationAction == "merge" && m.Ctx.Page == config.EcosystemCreatePage:
			p = "Are you sure you want to merge this PR? (Y/n) "

		case m.PromptConfirmationAction == "update" && m.Ctx.Page == config.EcosystemCreatePage:
			p = "Are you sure you want to update this PR? (Y/n) "

		case m.PromptConfirmationAction == "close" && m.Ctx.Page == config.ConnectorDetailsPage:
			p = "Are you sure you want to close this issue? (Y/n) "

		case m.PromptConfirmationAction == "reopen" && m.Ctx.Page == config.ConnectorDetailsPage:
			p = "Are you sure you want to reopen this issue? (Y/n) "
		case m.PromptConfirmationAction == "delete" && m.Ctx.Page == config.ConnectorDetailsPage:
			p = "Are you sure you want to delete this branch? (Y/n) "
		case m.PromptConfirmationAction == "new" && m.Ctx.Page == config.ConnectorDetailsPage:
			p = "Enter branch name: "
		case m.PromptConfirmationAction == "create_pr" && m.Ctx.Page == config.ConnectorDetailsPage:
			p = "Enter PR title: "
		}

		m.PromptConfirmationBox.SetPrompt(p)

		return m.Ctx.Styles.ListViewPort.PagerStyle.Render(m.PromptConfirmationBox.View())
	}

	return ""
}

// IsPacketCaptureFocused checks if the packet capture feature is currently in focus, returning true if it is.
func (m *DashboardBaseModel) IsPacketCaptureFocused() bool {
	return m.IsPacketCapturing
}

// SetIsPacketCapturing sets the packet capturing state and adjusts the focus state of the SearchBar accordingly.
func (m *DashboardBaseModel) SetIsPacketCapturing(val bool) tea.Cmd {
	m.IsPacketCapturing = val
	if val {
		return m.Table.Init()
	}

	return nil
}
