package connector_form

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/form"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/keys"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// maxWidth defines the maximum allowed width for the application elements, ensuring consistent layout and readability.
const maxWidth = 80

// Styles is a container for styling various UI components using lipgloss styles.
type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

// NewStyles initializes and returns a Styles struct with predefined lipgloss styles for UI components.
func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(theme.DefaultTheme.PrimaryColor).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(theme.DefaultTheme.TertiaryColor).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(theme.DefaultTheme.SecondaryColor)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

// state represents the application states and is typically used with iota to define specific states.
type state int

// statusNormal represents the default or normal state in the state enum.
// stateDone indicates that the process or operation is complete in the state enum.
const (
	statusNormal state = iota
	stateDone
)

// burger represents the type of burger being prepared or served.
// toppings holds a list of toppings added to the burger.
// classifications categorizes the burger based on its type or style.
// sauceLevel indicates the amount of sauce applied to the burger.
// name is the name of the burger or dish.
// instructions contain preparation or serving instructions for the burger.
// discount determines whether a discount is applicable to the burger.
var (
	burger          string
	toppings        []string
	classifications []string
	sauceLevel      int
	name            string
	instructions    string
	discount        bool
)

// Model represents the main application model containing state, renderer, styles, form, rendered form, and layout width.
type Model struct {
	state        state
	lg           *lipgloss.Renderer
	styles       *Styles
	form         *huh.Form
	renderedForm string
	width        int
}

// NewModel initializes and returns a new Model with predefined configurations and styles.
func NewModel(_ *context.ProgramContext) Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("class").
				Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
				Title("Choose your class").
				Description("This will determine your department"),

			huh.NewSelect[string]().
				Key("level").
				Options(huh.NewOptions("1", "20", "9999")...).
				Title("Choose your level").
				Description("This will determine your benefits package"),

			huh.NewSelect[string]().
				Key("type").
				Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
				Title("What connector type?").
				Description("This will determine your department"),

			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("Lettuce", "Lettuce").Selected(true),
					huh.NewOption("Tomatoes", "Tomatoes").Selected(true),
					huh.NewOption("Charm Sauce", "Charm Sauce"),
					huh.NewOption("Jalapeños", "Jalapeños"),
					huh.NewOption("Cheese", "Cheese"),
					huh.NewOption("Vegan Cheese", "Vegan Cheese"),
					huh.NewOption("Nutella", "Nutella"),
				).
				Title("Data Classification").
				Limit(4).
				Value(&classifications),

			form.NewConfirm().
				Key("done").
				Title("Register and begin listening?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("let's finish up")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("Wait, No"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false).
		WithTheme(theme.FormTheme()).
		WithKeyMap(keys.NewFormKeyMap())

	m.form.Init()

	return m
}

// minimum returns the smaller of two integers, x and y. If x is less than or equal to y, it returns x; otherwise, it returns y.
func minimum(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Update processes the incoming message, updates the model's state, and returns the updated model along with commands.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = minimum(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	}

	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

// View renders the current state of the model as a string for display in the terminal. It adapts based on the form state.
func (m Model) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		title, role := m.getRole()
		title = s.Highlight.Render(title)
		var b strings.Builder
		fmt.Fprintf(&b, "Congratulations, you’re Charm’s newest\n%s!\n\n", title)
		fmt.Fprintf(&b, "Your job description is as follows:\n\n%s\n\nPlease proceed to HR immediately.", role)
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:

		// Form (left side)
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)
		m.renderedForm = form

		errors := m.form.Errors()
		//header := m.appBoundaryView("Charm Employment Application")
		//if len(errors) > 0 {
		//	header = m.appErrorBoundaryView(m.errorView())
		//}
		body := lipgloss.JoinHorizontal(lipgloss.Top, form)

		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
		if len(errors) > 0 {
			footer = m.appErrorBoundaryView("")
		}

		return s.Base.Render(body + "\n\n" + footer)
	}
}

// SidebarView generates the sidebar UI component for the form, displaying current build details and projected role information.
func (m Model) SidebarView() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		title, role := m.getRole()
		title = s.Highlight.Render(title)
		var b strings.Builder
		fmt.Fprintf(&b, "Congratulations, you’re Charm’s newest\n%s!\n\n", title)
		fmt.Fprintf(&b, "Your job description is as follows:\n\n%s\n\nPlease proceed to HR immediately.", role)
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:

		var class string
		if m.form.GetString("class") != "" {
			class = "Class: " + m.form.GetString("class")
		}

		var status string
		{
			var (
				buildInfo      = "(None)"
				role           string
				jobDescription string
				level          string
			)

			if m.form.GetString("level") != "" {
				level = "Level: " + m.form.GetString("level")
				role, jobDescription = m.getRole()
				role = "\n\n" + s.StatusHeader.Render("Projected Role") + "\n" + role
				jobDescription = "\n\n" + s.StatusHeader.Render("Duties") + "\n" + jobDescription
			}
			if m.form.GetString("class") != "" {
				buildInfo = fmt.Sprintf("%s\n%s", class, level)
			}

			const statusWidth = 28
			statusMarginLeft := m.width - statusWidth - lipgloss.Width(m.renderedForm) - s.Status.GetMarginRight()
			status = s.Status.
				Height(lipgloss.Height(m.renderedForm)).
				Width(statusWidth).
				MarginLeft(statusMarginLeft).
				Render(s.StatusHeader.Render("Current Build") + "\n" +
					buildInfo +
					role +
					jobDescription)
		}
		body := lipgloss.JoinHorizontal(lipgloss.Top, status)
		return s.Base.Render(body)
	}
}

// errorView generates a string containing all error messages from the form's Errors method and concatenates them.
func (m Model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

// appBoundaryView formats and horizontally aligns the provided text based on the model's width and style configurations.
func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		// lipgloss.WithWhitespaceForeground(indigo),
	)
}

// appErrorBoundaryView renders an error-styled boundary with the given text, formatted to fit within the defined width.
func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(theme.DefaultTheme.SecondaryColor),
	)
}

// getRole returns the role title and description based on the class and level obtained from the form data.
func (m Model) getRole() (string, string) {
	level := m.form.GetString("level")
	switch m.form.GetString("class") {
	case "Warrior":
		switch level {
		case "1":
			return "Tank Intern", "Assists with tank-related activities. Paid position."
		case "9999":
			return "Tank Manager", "Manages tanks and tank-related activities."
		default:
			return "Tank", "General tank. Does damage, takes damage. Responsible for tanking."
		}
	case "Mage":
		switch level {
		case "1":
			return "DPS Associate", "Finds DPS deals and passes them on to DPS Manager."
		case "9999":
			return "DPS Operating Officer", "Oversees all DPS activities."
		default:
			return "DPS", "Does damage and ideally does not take damage. Logs hours in JIRA."
		}
	case "Rogue":
		switch level {
		case "1":
			return "Stealth Junior Designer", "Designs rogue-like activities. Reports to Stealth Lead."
		case "9999":
			return "Stealth Lead", "Lead designer for all things stealth. Some travel required."
		default:
			return "Sneaky Person", "Sneaks around and does sneaky things. Reports to Stealth Lead."
		}
	default:
		return "", ""
	}
}
