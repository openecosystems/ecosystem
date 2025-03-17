package sections

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	footer "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/footer"
	tabs "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/tabs"
	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	constants "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	tasks "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks"
	theme "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	utils "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"

	cliv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
)

var once sync.Once

// BaseModel encapsulates the primary state and components for managing UI layout, navigation, and application context.
type BaseModel struct {
	Ctx           *context.ProgramContext
	Pages         []contract.Page
	CurrentPage   contract.Page
	CurrentPageID int
	Spinner       *spinner.Model
	Tabs          *tabs.Model
	Footer        *footer.Model
	Keys          *keys.KeyMap
	SingularForm  string
	PluralForm    string
}

// NewBaseOptions holds configuration for initializing a base model with singular/plural names, pages, and settings.
type NewBaseOptions struct {
	Singular      string
	Plural        string
	Pages         []contract.Page
	CurrentPageID int
	Settings      *cliv2alphalib.CLIConfiguration
}

// NewBaseModel creates and initializes a new BaseModel instance with the provided context and options.
func NewBaseModel(pctx *context.ProgramContext, options *NewBaseOptions) *BaseModel {
	if len(options.Pages) == 0 {
		panic("No pages provided")
	}

	taskSpinner := spinner.New(spinner.WithSpinner(spinner.Dot))
	tabsModel := tabs.NewModel(pctx, options.Pages)
	footerModel := footer.NewModel(pctx)

	m := &BaseModel{
		Ctx:          pctx,
		Spinner:      &taskSpinner,
		SingularForm: options.Singular,
		PluralForm:   options.Plural,
		Pages:        options.Pages,
		Tabs:         tabsModel,
		Footer:       footerModel,
	}

	m.Spinner.Style = lipgloss.NewStyle().Background(m.Ctx.Theme.SelectedBackground)

	m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(options.CurrentPageID)

	m.Ctx.Page = m.CurrentPage.GetPageSettings().Type
	m.Ctx.Settings = options.Settings
	m.Tabs = m.Tabs.SetCurrentPageID(m.CurrentPageID)

	return m
}

// initialize represents a configuration initialization containing the application's main configuration.
type initialize struct {
	Config config.Config
}

// init initializes the application by parsing the configuration file and handling potential parsing errors.
// Returns an `initialize` message containing the parsed configuration.
func (m *BaseModel) init() tea.Msg {
	cfg, err := config.ParseConfig()
	if err != nil {
		utils.ShowError(err)
		return initialize{Config: cfg}
	}

	return initialize{Config: cfg}
}

// InitBase initializes the BaseModel by batching the execution of the base `init` method and returning a command.
func (m *BaseModel) InitBase() tea.Cmd {
	tasks.ProcessTasks()
	go m.HandleCompletedTasks()
	return tea.Batch(m.init)
}

// UpdateBase processes incoming messages to update the state of the BaseModel and returns the updated model and command.
func (m *BaseModel) UpdateBase(msg tea.Msg) (*BaseModel, tea.Cmd) {
	var (
		cmd            tea.Cmd
		cmds           []tea.Cmd
		footerCmd      tea.Cmd
		tabsCmd        tea.Cmd
		currentPageCmd tea.Cmd
	)

	switch message := msg.(type) {
	case tea.KeyMsg:
		m.Ctx.Error = nil

		switch {
		case key.Matches(message, keys.Keys.PrevPage):
			m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(m.GetPrevPageID())
			m.Tabs = m.Tabs.SetCurrentPageID(m.CurrentPageID)
			m.Ctx.Page = m.CurrentPage.GetPageSettings().Type

		case key.Matches(message, keys.Keys.NextPage):
			m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(m.GetNextPageID())
			m.Tabs = m.Tabs.SetCurrentPageID(m.CurrentPageID)
			m.Ctx.Page = m.CurrentPage.GetPageSettings().Type

		case key.Matches(message, keys.Keys.Help):
			if !m.Footer.ShowAll {
				m.Ctx.PageContentHeight = m.Ctx.PageContentHeight + theme.FooterHeight - theme.ExpandedHelpHeight
			} else {
				m.Ctx.PageContentHeight = m.Ctx.PageContentHeight + theme.ExpandedHelpHeight - theme.FooterHeight
			}

		case key.Matches(message, keys.Keys.Quit):
			if !m.Ctx.Config.ConfirmQuit {
				_, cmd = m.Footer.Update(msg)
				return m, cmd
			}

			// cmds = append(cmds, tea.Quit)
			// return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		m.OnWindowSizeChanged(message)
	case initialize:
		m.Ctx.Config = &message.Config
		m.CurrentPage, m.CurrentPageID = m.SetCurrentPage(m.GetDefaultPageID())
	case tasks.TaskFinishedMsg:

		// taskSpinner, internalTickCmd := m.Spinner.Update(message)
		// m.Spinner = taskSpinner
		// cmds = append(cmds, internalTickCmd)

		// m.Footer.SetRightSection(m.RenderRunningTask(message))
		//_, footerCmd = m.Footer.Update(message)
		//m.CurrentPage.Update(message)

		m.Ctx.Logger.Debug("Section: Task finished", "id", message.Task.ID)
		//if message.Task.Error != nil {
		//	m.Ctx.Logger.Error("Task finished with error", "id", message.Task.ID, "err", message.Task.Error)
		//}
		//clr := tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
		//	return tasks.ClearTaskMsg{TaskID: message.Task.ID, Task: message.Task}
		//})
		//cmds = append(cmds, footerCmd, clr)
		cmds = append(cmds, footerCmd)

	case tasks.ClearTaskMsg:
		m.Ctx.Logger.Debug("Clear Task finished", "id", message.Task.ID)
		m.Footer.SetRightSection("")
		m.UpdateBase(message)
		//	delete(m.Tasks, message.TaskID)
		// case spinner.TickMsg:
		//	if len(m.Tasks) > 0 {
		//		taskSpinner, internalTickCmd := m.Spinner.Update(msg)
		//		m.Spinner = taskSpinner
		//		rTask := m.RenderRunningTask()
		//		m.Footer.SetRightSection(rTask)
		//		cmds = append(cmds, internalTickCmd)
		//	}
	}

	m.UpdateProgramContext(m.Ctx)
	_, tabsCmd = m.Tabs.Update(msg)
	_, footerCmd = m.Footer.Update(msg)
	_, currentPageCmd = m.CurrentPage.Update(msg)

	cmds = append(
		cmds,
		tabsCmd,
		footerCmd,
		currentPageCmd,
	)

	return m, tea.Batch(cmds...)
}

// HandleCompletedTasks processes completed and cleared tasks, updates their states in the model, and generates related commands.
func (m *BaseModel) HandleCompletedTasks() {
	once.Do(func() {
		for msg := range tasks.CompletedTaskCmdsChan {
			switch message := msg.(type) {
			case tasks.TaskFinishedMsg:
				_ = message
				m.UpdateBase(msg)
			}
		}
	})
}

// ViewBase generates and returns a string representation of the current UI, including tabs, content, error, and footer.
func (m *BaseModel) ViewBase(content string) string {
	s := strings.Builder{}
	s.WriteString(m.Tabs.View())
	s.WriteString("\n")
	if content == "" {
		content = "No page defined"
	}
	c := lipgloss.JoinHorizontal(
		lipgloss.Top,
		content,
	)
	s.WriteString(c)
	s.WriteString("\n")
	if m.Ctx.Error != nil {
		s.WriteString(
			m.Ctx.Styles.Common.ErrorStyle.
				Width(m.Ctx.ScreenWidth).
				Render(fmt.Sprintf("%s %s",
					m.Ctx.Styles.Common.FailureGlyph,
					lipgloss.NewStyle().
						Foreground(m.Ctx.Theme.ErrorText).
						Render(m.Ctx.Error.Error()),
				)),
		)
	} else {
		s.WriteString(m.Footer.View())
	}

	return s.String()
}

// ViewDebug generates and returns a structured debug representation of the BaseModel's current state as a strings.Builder.
func (m *BaseModel) ViewDebug() *strings.Builder {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString("Section: " + string(m.Ctx.Section) + "\n")
	s.WriteString("   Screen Width: " + strconv.Itoa(m.Ctx.ScreenWidth) + "\n")
	s.WriteString("   Screen Height: " + strconv.Itoa(m.Ctx.ScreenHeight) + "\n")
	s.WriteString("   Default Page ID: " + strconv.Itoa(m.GetDefaultPageID()) + "\n")
	s.WriteString("   Number of Section Pages: " + strconv.Itoa(len(m.Pages)) + "\n")
	for _, p := range m.Pages {
		s.WriteString("      Section Page: " + p.GetPageSettings().Title + "\n")
	}
	s.WriteString("\n")
	s.WriteString("   Current Page: " + string(m.Ctx.Page) + "\n")
	s.WriteString("      ID: " + strconv.Itoa(m.CurrentPageID) + "\n")
	s.WriteString("      Title: " + m.CurrentPage.GetPageSettings().Title + "\n")
	s.WriteString("      Width: " + strconv.Itoa(m.Ctx.PageContentWidth) + "\n")
	s.WriteString("      Height: " + strconv.Itoa(m.Ctx.PageContentHeight) + "\n")
	s.WriteString("      Main Content\n")
	s.WriteString("         Width: " + strconv.Itoa(m.Ctx.MainContentWidth) + "\n")
	s.WriteString("         Height: " + strconv.Itoa(m.Ctx.MainContentHeight) + "\n")
	s.WriteString("         Body Width: " + strconv.Itoa(m.Ctx.MainContentBodyWidth) + "\n")
	s.WriteString("         Body Height: " + strconv.Itoa(m.Ctx.MainContentBodyHeight) + "\n")
	s.WriteString("      Sidebar Content\n")
	s.WriteString("         Width: " + strconv.Itoa(m.Ctx.SidebarContentWidth) + "\n")
	s.WriteString("         Height: " + strconv.Itoa(m.Ctx.SidebarContentHeight) + "\n")
	s.WriteString("         Body Width: " + strconv.Itoa(m.Ctx.SidebarContentBodyWidth) + "\n")
	s.WriteString("         Body Height: " + strconv.Itoa(m.Ctx.SidebarContentBodyHeight) + "\n")
	nextPage := m.GetPageAt(m.GetNextPageID())
	s.WriteString("    Next Page: " + nextPage.GetPageSettings().Title + "\n")
	s.WriteString("       ID: " + strconv.Itoa(m.GetNextPageID()) + "\n")
	prevPage := m.GetPageAt(m.GetPrevPageID())
	s.WriteString("    Previous Page: " + prevPage.GetPageSettings().Title + "\n")
	s.WriteString("       ID: " + strconv.Itoa(m.GetPrevPageID()) + "\n")
	tabPage := m.GetPageAt(m.Tabs.GetCurrentPageID())
	s.WriteString("    Current Tab: " + tabPage.GetPageSettings().Title + "\n")
	s.WriteString("       ID: " + strconv.Itoa(m.Tabs.GetCurrentPageID()) + "\n")
	s.WriteString("    Available Tab Page Count: " + strconv.Itoa(len(m.Tabs.GetPages())) + "\n")
	for _, p := range m.Tabs.GetPages() {
		s.WriteString("       Tab Page: " + p.GetPageSettings().Title + "\n")
	}
	s.WriteString("\n")

	return &s
}

// UpdateProgramContext updates the BaseModel's context and propagates it to tabs, the current page, and the footer.
func (m *BaseModel) UpdateProgramContext(ctx *context.ProgramContext) {
	// m.Ctx = ctx
	m.Tabs.UpdateProgramContext(ctx)
	if m.CurrentPage != nil {
		m.CurrentPage.UpdateProgramContext(ctx)
	}
	m.Footer.UpdateProgramContext(ctx)
}

// OnWindowSizeChanged updates context dimensions and page content size based on the new window size from the message.
func (m *BaseModel) OnWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.Ctx.ScreenWidth = msg.Width
	m.Ctx.ScreenHeight = msg.Height
	m.Ctx.PageContentWidth = m.Ctx.ScreenWidth
	if m.Footer.ShowAll {
		m.Ctx.PageContentHeight = msg.Height - theme.TabsHeight - theme.ExpandedHelpHeight
	} else {
		m.Ctx.PageContentHeight = msg.Height - theme.TabsHeight - theme.FooterHeight
	}

	m.CurrentPage.OnWindowSizeChanged(m.Ctx)
}

// SetCurrentPage sets the current page of the BaseModel to the page at the given ID, updates the context and tabs, and returns the page and its ID.
func (m *BaseModel) SetCurrentPage(id int) (contract.Page, int) {
	page := m.GetPageAt(id)

	m.Ctx.Page = page.GetPageSettings().Type
	m.CurrentPageID = id
	m.CurrentPage = page
	m.Tabs = m.Tabs.SetCurrentPageID(id)

	return page, id
}

// GetCurrentPage returns the currently active page from the Pages slice based on the CurrentPageID. Returns nil if no valid page exists.
func (m *BaseModel) GetCurrentPage() contract.Page {
	p := m.Pages
	if len(p) == 0 || m.CurrentPageID >= len(p) {
		return nil
	}

	return p[m.CurrentPageID]
}

// GetPageAt retrieves the page at the specified index from the Pages slice. If the index is out of range, it returns nil.
func (m *BaseModel) GetPageAt(id int) contract.Page {
	p := m.Pages
	if len(p) <= id {
		return nil
	}

	return p[id]
}

// GetPrevPageID calculates and returns the ID of the previous page in the Pages slice, wrapping around if necessary.
func (m *BaseModel) GetPrevPageID() int {
	return (m.CurrentPageID - 1 + len(m.Pages)) % len(m.Pages)
}

// GetNextPageID calculates and returns the ID of the next page, cycling back to the start if the end is reached.
func (m *BaseModel) GetNextPageID() int {
	return (m.CurrentPageID + 1) % len(m.Pages)
}

// GetDefaultPageID identifies and returns the index of the default page in the Pages slice. Defaults to 0 if none is found.
func (m *BaseModel) GetDefaultPageID() int {
	for i, page := range m.Pages {
		if page.GetPageSettings().IsDefault {
			return i
		}
	}

	return 0
}

// RenderRunningTask returns a styled string representation of the currently running task and its status, including task count stats.
func (m *BaseModel) RenderRunningTask(msg tasks.TaskFinishedMsg) string {
	var currTaskStatus string

	switch msg.State {
	case tasks.TaskStart:
		currTaskStatus = lipgloss.NewStyle().
			Background(m.Ctx.Theme.SelectedBackground).
			Render(
				fmt.Sprintf(
					"%s%s",
					m.Spinner.View(),
					msg.Task.StartText,
				))
	case tasks.TaskError:
		currTaskStatus = lipgloss.NewStyle().
			Foreground(m.Ctx.Theme.ErrorText).
			Background(m.Ctx.Theme.SelectedBackground).
			Render(fmt.Sprintf("%s %s", constants.FailureIcon, msg.Task.Error.Error()))
	case tasks.TaskFinished:
		currTaskStatus = lipgloss.NewStyle().
			Foreground(m.Ctx.Theme.SuccessText).
			Background(m.Ctx.Theme.SelectedBackground).
			Render(fmt.Sprintf("%s %s", constants.SuccessIcon, msg.Task.FinishedText))
	}

	return lipgloss.NewStyle().
		Padding(0, 1).
		MaxHeight(1).
		Background(m.Ctx.Theme.SelectedBackground).
		Render(strings.TrimSpace(lipgloss.JoinHorizontal(lipgloss.Top, currTaskStatus)))
}
