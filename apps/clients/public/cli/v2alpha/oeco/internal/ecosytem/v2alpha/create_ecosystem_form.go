// Package ecosystemv2alphapbint
// nolint:unused,revive
package ecosystemv2alphapbint

import (
	"errors"
	"fmt"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 80

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

// Styles defines a collection of reusable lipgloss style definitions for various UI components such as Base, HeaderText, Status, etc.
type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

// NewStyles initializes and returns a pointer to a Styles struct containing customized lipgloss styles.
// It sets up styles for Base, HeaderText, Status, StatusHeader, Highlight, ErrorHeaderText, and Help.
func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

// state represents an integer-based enumeration to define various states in a model or application logic.
type state int

// statusNormal represents the default or normal state.
// stateDone represents the completed or finished state.
const (
	statusNormal state = iota
	stateDone
)

type CreateEcosystemData struct {
	Name          string
	EcosystemType string
	Cidr          string
}

// CreateEcosystemModel represents the main structure containing the state, styles, renderer, form, and width of the application.
type CreateEcosystemModel struct {
	Data   CreateEcosystemData
	state  state
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

// NewCreateEcosystemModel creates and initializes a new CreateEcosystemModel instance with default configurations and predefined form settings.
func NewCreateEcosystemModel(_ *cliv2alphalib.Configuration) *CreateEcosystemModel {
	m := CreateEcosystemModel{width: maxWidth, state: statusNormal}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	var data CreateEcosystemData

	m.form = huh.NewForm(
		huh.NewGroup(

			huh.NewInput().
				Title("Name your Economic System").
				Suggestions([]string{"For example 'Open Ecosystems'"}).
				Prompt("? ").
				// Validate(isFood).
				Value(&data.Name),

			huh.NewSelect[string]().
				Key("class").
				Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
				Title("Choose your class").
				Description("This will determine your department").
				Value(&data.Name),

			huh.NewSelect[string]().
				Key("level").
				Options(huh.NewOptions("1", "20", "9999")...).
				Title("Choose your level").
				Description("This will determine your benefits package").
				Value(&data.EcosystemType),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return errors.New("welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)

	m.Data = data
	return &m
}

// Init initializes the model and sets up the form. Returns an initial command to be executed.
func (m *CreateEcosystemModel) Init() tea.Cmd {
	return m.form.Init()
}

// _min returns the smaller of two integers, x and y. If x is greater than y, it returns y; otherwise, it returns x.
func _min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Update processes incoming messages and updates the model's state and commands accordingly.
// It handles window resize, keyboard input, form updates, and determines when to quit the application.
// Returns the updated model with a batch of commands to execute next.
func (m *CreateEcosystemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = _min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			// return m, tea.Interrupt
			return m, tea.Quit
		case "esc", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

// View generates and returns the string representation of the current UI based on the model's state and data.
func (m *CreateEcosystemModel) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		title, role := m.getRole()
		title = s.Highlight.Render(title)
		m.state = stateDone
		var b strings.Builder
		_, err := fmt.Fprintf(&b, "Congratulations, you’re Charm’s newest\n%s!\n\n", title)
		if err != nil {
			return ""
		}
		_, err = fmt.Fprintf(&b, "Your job description is as follows:\n\n%s\n\nPlease proceed to HR immediately.", role)
		if err != nil {
			return ""
		}
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:

		var class string
		if m.form.GetString("class") != "" {
			class = "Class: " + m.form.GetString("class")
		}

		// Form (left side)
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		// Status (right side)
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
			statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
			status = s.Status.
				Height(lipgloss.Height(form)).
				Width(statusWidth).
				MarginLeft(statusMarginLeft).
				Render(s.StatusHeader.Render("Current Build") + "\n" +
					buildInfo +
					role +
					jobDescription)
		}

		errs := m.form.Errors()
		header := m.appBoundaryView("Create a new Economic System")
		if len(errs) > 0 {
			header = m.appErrorBoundaryView(m.errorView())
		}
		body := lipgloss.JoinHorizontal(lipgloss.Top, form, status)

		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
		if len(errs) > 0 {
			footer = m.appErrorBoundaryView("")
		}

		return s.Base.Render(header + "\n" + body + "\n\n" + footer)
	}
}

// errorView aggregates and returns all error messages from the model's form as a single concatenated string.
func (m *CreateEcosystemModel) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

// appBoundaryView places horizontally aligned text within the app boundary with custom styles and whitespace characters.
func (m *CreateEcosystemModel) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars(""),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

// appErrorBoundaryView formats and styles an error header message within a bounded horizontal space using the provided text.
func (m *CreateEcosystemModel) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

// getRole determines the role and its description based on the class and level values retrieved from the form.
func (m *CreateEcosystemModel) getRole() (string, string) {
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
			return "Stealth Junior Designer", "Designs rougue-like activities. Reports to Stealth Lead."
		case "9999":
			return "Stealth Lead", "Lead designer for all things stealth. Some travel required."
		default:
			return "Sneaky Person", "Sneaks around and does sneaky things. Reports to Stealth Lead."
		}
	default:
		return "", ""
	}
}
