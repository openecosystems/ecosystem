package ecosystemcreateform

import (
	"errors"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/segmentio/ksuid"

	components "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	keys "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	tasks "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks"
	ecosystem "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks/ecosystem"
	theme "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	ecosystemv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
)

// maxWidth defines the maximum allowed width for the application elements, ensuring consistent layout and readability.
const maxWidth = theme.MaxContentWidth

// state represents an integer-based enumeration to define various states in a model or application logic.
type state int

// statusNormal represents the default or normal state.
// stateDone represents the completed or finished state.
const (
	statusNormal state = iota
	stateDone
)

// ModelConfig represents the configuration settings for a specific model, encapsulating its properties and preferences.
type ModelConfig struct{}

// CreateEcosystemData defines the structure for encapsulating information related to creating an ecosystem.
// It includes the domain, type of ecosystem, and CIDR notation for the network.
type CreateEcosystemData struct {
	Domain        string
	EcosystemType string
	CIDR          string

	ecosystemType ecosystemv2alphapb.EcosystemType
}

// Model represents the main application model containing state, renderer, styles, form, rendered form, and layout width.
type Model struct {
	components.BaseModel[ModelConfig]
	Data *CreateEcosystemData

	state        state
	lg           *lipgloss.Renderer
	styles       *theme.Styles
	form         *huh.Form
	renderedForm string
	width        int
}

// NewModel initializes and returns a new Model with predefined configurations and styles.
func NewModel(pctx *context.ProgramContext) contract.Component {
	m := &Model{
		width:  maxWidth,
		state:  statusNormal,
		lg:     lipgloss.DefaultRenderer(),
		styles: &pctx.Styles,
		BaseModel: components.BaseModel[ModelConfig]{
			Ctx:             pctx,
			Keys:            nil,
			KeyBindings:     nil,
			ComponentConfig: &ModelConfig{},
		},
	}

	var data CreateEcosystemData

	m.form = huh.NewForm(
		huh.NewGroup(

			huh.NewInput().
				Key("domain").
				Title("Pick a domain name").
				DescriptionFunc(func() string {
					if data.Domain == "" {
						return "For example if you choose 'oeco', your domain name will be 'oeco.mesh'"
					}
					return "Your domain name will be: " + data.Domain + ".mesh"
				}, &data).
				Suggestions([]string{"For example 'oeco'"}).
				Prompt("? ").
				// Validate(isFood).
				Value(&data.Domain).WithHeight(29),

			huh.NewSelect[string]().
				Key("type").
				Options(huh.NewOptions("Private", "Public")...).
				Title("Choose your ecosystem type").
				Description("This will impact how people access your ecosystem").
				Value(&data.EcosystemType),

			huh.NewSelect[string]().
				Key("cidr").
				Options(huh.NewOptions("192.168.0.0/16", "172.16.0.0/12", "10.0.0.0/8", "100.64.0.0/10")...).
				Title("Choose your CIDR block").
				Description("The range you choose must be part of the RFC 1918 Private Address Space or the RFC 6598 Shared Address Space.").
				Value(&data.CIDR),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return errors.New("we found issues")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("Wait, no"),
		),
	).
		WithWidth(80).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(theme.FormTheme()).
		WithKeyMap(keys.NewFormKeyMap())

	m.form.Init()
	m.Data = &data

	m.Ctx.Logger.Debug("Component - Ecosystem Create Form: Initial Configuration")

	return m
}

// minimum returns the smaller of two integers, x and y. If x is less than or equal to y, it returns x; otherwise, it returns y.
func minimum(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *Model) Init() tea.Cmd {
	return tea.Batch()
}

// Update processes the incoming message, updates the model's state, and returns the updated model along with commands.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.BaseModel, cmd = m.UpdateBase(msg)

	switch message := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = minimum(message.Width, maxWidth) - m.styles.MainContent.ContainerStyle.GetHorizontalFrameSize()
	}

	if m.state != stateDone {
		frm, _cmd := m.form.Update(msg)
		if f, ok := frm.(*huh.Form); ok {
			m.form = f
			if f.State == huh.StateCompleted {
				m.state = stateDone
				tasks.AddTask(tasks.Task{
					Ctx:          m.Ctx,
					ID:           ksuid.New().String(),
					StartText:    "Creating Ecosystem",
					FinishedText: "Created Ecosystem",
					State:        tasks.TaskStart,
					Error:        nil,
					StartTime:    time.Now(),
					TaskExecutor: &ecosystem.CreateEcosystemMsg{
						CreateEcosystemRequest: ecosystemv2alphapb.CreateEcosystemRequest{
							Slug: m.Data.Domain,
							Type: m.Data.ecosystemType,
							Name: m.Data.Domain,
							Cidr: m.Data.CIDR,
						},
					},
					Done: false,
				})
			}
			cmds = append(cmds, _cmd)
		}
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the current state of the model as a string for display in the terminal. It adapts based on the form state.
func (m *Model) View() string {
	switch m.form.State {
	case huh.StateCompleted:
		return m.CompletedView()
	default:
		return m.InProgressView()
	}
}

// InProgressView renders the view for the state when the form is in progress, combining the form and help view with styling.
func (m *Model) InProgressView() string {
	s := m.styles
	v := strings.TrimSuffix(m.form.View(), "\n\n")
	frm := m.lg.NewStyle().Margin(1, 0).Render(v)
	m.renderedForm = frm

	errs := m.form.Errors()
	body := lipgloss.JoinHorizontal(lipgloss.Top, frm)

	footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
	if len(errs) > 0 {
		footer = m.appErrorBoundaryView("")
	}

	return s.MainContent.ContainerStyle.Render(body + "\n\n" + footer)
}

// CompletedView generates a congratulatory message for the user upon form completion, including their new role and description.
func (m *Model) CompletedView() string {
	s := m.styles
	title, role := m.getEcosystemType()
	title = s.Header.H2Text.Render(title)
	var b strings.Builder
	_, _ = fmt.Fprintf(&b, "Congratulations, your Ecosystem configuration is ready.\n%s!\n\n", title)
	_, _ = fmt.Fprintf(&b, "Please follow the documentation \n\n%s\n\nto get the Ecosystem started.", role)
	return s.StatusBox.Box.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
}

// SidebarView generates the sidebar UI component for the form, displaying current build details and projected role information.
func (m *Model) SidebarView() string {
	switch m.form.State {
	case huh.StateCompleted:
		return m.SideBarCompletedView()
	default:
		return m.SideBarInProgressView()
	}
}

// SideBarInProgressView generates a styled sidebar view displaying current build status, projected role, and related information.
func (m *Model) SideBarInProgressView() string {
	s := m.styles

	var domain string
	if m.form.GetString("domain") != "" {
		domain = "Domain Name: " + m.form.GetString("domain") + ".mesh" +
			"\n\n" + "Edge Endpoint: edge." + m.form.GetString("domain") + ".mesh" +
			"\n\n" + "Mesh Endpoint: api." + m.form.GetString("domain") + ".mesh"
	}

	var status string
	{
		var (
			domainInfo    = "(None)"
			role          string
			ecosystemType string
			eType         string
		)

		if m.form.GetString("type") != "" {
			eType = "Type: " + m.form.GetString("type")
			role, ecosystemType = m.getEcosystemType()
			role = "\n\n" + s.Sidebar.StatusHeader.Render("Ecosystem Role") + "\n" + role
			ecosystemType = "\n\n" + s.Sidebar.StatusHeader.Render("Type") + "\n" + ecosystemType
		}
		if m.form.GetString("domain") != "" {
			domainInfo = fmt.Sprintf("%s\n%s", domain, eType)
		}

		// const statusWidth = 28
		// statusMarginLeft := m.width - statusWidth - lipgloss.Width(m.renderedForm) - s.StatusBox.Box.GetMarginRight()
		status = s.StatusBox.Box.
			Height(lipgloss.Height(m.renderedForm)).
			// Width(statusWidth).
			// MarginLeft(statusMarginLeft).
			Render(s.Header.H1Text.Render("Ecosystem Details") + "\n" +
				domainInfo +
				role +
				ecosystemType)
	}

	body := lipgloss.JoinHorizontal(lipgloss.Top, status)

	return s.Sidebar.ContainerStyle.Render(body)
}

// SideBarCompletedView generates a congratulatory view for the user upon form completion, displaying their new role and description.
func (m *Model) SideBarCompletedView() string {
	s := m.styles
	title, role := m.getEcosystemType()
	title = s.Header.H2Text.Render(title)
	var b strings.Builder
	_, _ = fmt.Fprintf(&b, "Congratulations, you’re Charm’s newest\n%s!\n\n", title)
	_, _ = fmt.Fprintf(&b, "Your job description is as follows:\n\n%s\n\nPlease proceed to HR immediately.", role)
	return s.StatusBox.Box.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
}

// errorView aggregates and returns all error messages from the model's form as a single concatenated string.
//func (m *Model) errorView() string {
//	var s string
//	for _, err := range m.form.Errors() {
//		s += err.Error()
//	}
//	return s
//}

// appBoundaryView formats and horizontally aligns the provided text based on the model's width and style configurations.
func (m *Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.Header.H1Text.Render(text),
		// lipgloss.WithWhitespaceForeground(indigo),
	)
}

// appErrorBoundaryView renders an error-styled boundary with the given text, formatted to fit within the defined width.
func (m *Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.Header.H1TextError.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(theme.DefaultTheme.SecondaryColor500),
	)
}

// getEcosystemType determines the role and its description based on the class and level values retrieved from the form.
func (m *Model) getEcosystemType() (string, string) {
	switch m.form.GetString("type") {
	case "Private":
		m.Data.ecosystemType = ecosystemv2alphapb.EcosystemType_ECOSYSTEM_TYPE_PRIVATE
		return "Private", "An Ecosystem that is private to your company and network. "
	case "Public":
		m.Data.ecosystemType = ecosystemv2alphapb.EcosystemType_ECOSYSTEM_TYPE_PUBLIC
		return "Public", "An Ecosystem that is public and allows access to people over the internet"
	default:
		return "", ""
	}
}
