package form

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/accessibility"
	"github.com/charmbracelet/lipgloss"
)

// updateFieldMsg is a message to update the fields of a group that is currently
// displayed.
//
// This is used to update all TitleFunc, DescriptionFunc, and ...Func update
// methods to make all fields dynamically update based on user input.
type updateFieldMsg struct{}

// Confirm is a form confirm field.
type Confirm struct {
	accessor huh.Accessor[bool]
	key      string
	id       int

	// customization
	title       Eval[string]
	description Eval[string]
	affirmative string
	negative    string

	// error handling
	validate func(bool) error
	err      error

	// state
	focused bool

	// options
	width      int
	height     int
	inline     bool
	accessible bool
	theme      *huh.Theme
	keymap     huh.ConfirmKeyMap
}

// NewConfirm returns a new confirm field.
func NewConfirm() *Confirm {
	return &Confirm{
		accessor: &huh.EmbeddedAccessor[bool]{},
		// id:          nextID(),
		title:       Eval[string]{cache: make(map[uint64]string)},
		description: Eval[string]{cache: make(map[uint64]string)},
		affirmative: "Yes",
		negative:    "No",
		validate:    func(bool) error { return nil },
	}
}

// Validate sets the validation function of the confirm field.
func (c *Confirm) Validate(validate func(bool) error) *Confirm {
	c.validate = validate
	return c
}

// Error returns the error of the confirm field.
func (c *Confirm) Error() error {
	return c.err
}

// Skip returns whether the confirm should be skipped or should be blocking.
func (*Confirm) Skip() bool {
	return false
}

// Zoom returns whether the input should be zoomed.
func (*Confirm) Zoom() bool {
	return false
}

// Affirmative sets the affirmative value of the confirm field.
func (c *Confirm) Affirmative(affirmative string) *Confirm {
	c.affirmative = affirmative
	return c
}

// Negative sets the negative value of the confirm field.
func (c *Confirm) Negative(negative string) *Confirm {
	c.negative = negative
	return c
}

// Value sets the value of the confirm field.
func (c *Confirm) Value(value *bool) *Confirm {
	return c.Accessor(huh.NewPointerAccessor(value))
}

// Accessor sets the accessor of the confirm field.
func (c *Confirm) Accessor(accessor huh.Accessor[bool]) *Confirm {
	c.accessor = accessor
	return c
}

// Key sets the key of the confirm field.
func (c *Confirm) Key(key string) *Confirm {
	c.key = key
	return c
}

// Title sets the title of the confirm field.
func (c *Confirm) Title(title string) *Confirm {
	c.title.val = title
	c.title.fn = nil
	return c
}

// TitleFunc sets the title func of the confirm field.
func (c *Confirm) TitleFunc(f func() string, bindings any) *Confirm {
	c.title.fn = f
	c.title.bindings = bindings
	return c
}

// Description sets the description of the confirm field.
func (c *Confirm) Description(description string) *Confirm {
	c.description.val = description
	c.description.fn = nil
	return c
}

// DescriptionFunc sets the description function of the confirm field.
func (c *Confirm) DescriptionFunc(f func() string, bindings any) *Confirm {
	c.description.fn = f
	c.description.bindings = bindings
	return c
}

// Inline sets whether the field should be inline.
func (c *Confirm) Inline(inline bool) *Confirm {
	c.inline = inline
	return c
}

// Focus focuses the confirm field.
func (c *Confirm) Focus() tea.Cmd {
	c.focused = true
	return nil
}

// Blur blurs the confirm field.
func (c *Confirm) Blur() tea.Cmd {
	c.focused = false
	c.err = c.validate(c.accessor.Get())
	return nil
}

// KeyBinds returns the help message for the confirm field.
func (c *Confirm) KeyBinds() []key.Binding {
	return []key.Binding{c.keymap.Toggle, c.keymap.Prev, c.keymap.Submit, c.keymap.Next, c.keymap.Accept, c.keymap.Reject}
}

// Init initializes the confirm field.
func (c *Confirm) Init() tea.Cmd {
	return nil
}

// Update updates the confirm field.
func (c *Confirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case updateFieldMsg:
		if ok, hash := c.title.shouldUpdate(); ok {
			c.title.bindingsHash = hash
			if !c.title.loadFromCache() {
				c.title.loading = true
				cmds = append(cmds, func() tea.Msg {
					return updateTitleMsg{id: c.id, title: c.title.fn(), hash: hash}
				})
			}
		}
		if ok, hash := c.description.shouldUpdate(); ok {
			c.description.bindingsHash = hash
			if !c.description.loadFromCache() {
				c.description.loading = true
				cmds = append(cmds, func() tea.Msg {
					return updateDescriptionMsg{id: c.id, description: c.description.fn(), hash: hash}
				})
			}
		}

	case updateTitleMsg:
		if msg.id == c.id && msg.hash == c.title.bindingsHash {
			c.title.val = msg.title
			c.title.loading = false
		}
	case updateDescriptionMsg:
		if msg.id == c.id && msg.hash == c.description.bindingsHash {
			c.description.val = msg.description
			c.description.loading = false
		}
	case tea.KeyMsg:
		c.err = nil
		switch {
		case key.Matches(msg, c.keymap.Toggle):
			if c.negative == "" {
				break
			}
			c.accessor.Set(!c.accessor.Get())
		case key.Matches(msg, c.keymap.Prev):
			cmds = append(cmds, huh.PrevField)
		case key.Matches(msg, c.keymap.Next, c.keymap.Submit):
			cmds = append(cmds, huh.NextField)
		case key.Matches(msg, c.keymap.Accept):
			c.accessor.Set(true)
			cmds = append(cmds, huh.NextField)
		case key.Matches(msg, c.keymap.Reject):
			c.accessor.Set(false)
			cmds = append(cmds, huh.NextField)
		}
	}

	return c, tea.Batch(cmds...)
}

func (c *Confirm) activeStyles() *huh.FieldStyles {
	theme := c.theme
	if theme == nil {
		theme = huh.ThemeCharm()
	}
	if c.focused {
		return &theme.Focused
	}
	return &theme.Blurred
}

// View renders the confirm field.
func (c *Confirm) View() string {
	styles := c.activeStyles()

	var sb strings.Builder
	sb.WriteString(styles.Title.Render(c.title.val))
	if c.err != nil {
		sb.WriteString(styles.ErrorIndicator.String())
	}

	description := styles.Description.Render(c.description.val)

	if !c.inline && (c.description.val != "" || c.description.fn != nil) {
		sb.WriteString("\n")
	}
	sb.WriteString(description)

	if !c.inline {
		sb.WriteString("\n")
		sb.WriteString("\n")
	}

	var negative string
	var affirmative string
	if c.negative != "" {
		if c.accessor.Get() {
			affirmative = styles.FocusedButton.Render(c.affirmative)
			negative = styles.BlurredButton.Render(c.negative)
		} else {
			affirmative = styles.BlurredButton.Render(c.affirmative)
			negative = styles.FocusedButton.Render(c.negative)
		}
		c.keymap.Reject.SetHelp("n", c.negative)
	} else {
		affirmative = styles.FocusedButton.Render(c.affirmative)
		c.keymap.Reject.SetEnabled(false)
	}

	c.keymap.Accept.SetHelp("y", c.affirmative)

	buttonsRow := lipgloss.JoinVertical(lipgloss.Left, affirmative, " ", negative)

	promptWidth := lipgloss.Width(sb.String())
	buttonsWidth := lipgloss.Width(buttonsRow)

	renderWidth := promptWidth
	if buttonsWidth > renderWidth {
		renderWidth = buttonsWidth
	}

	style := lipgloss.NewStyle().Width(renderWidth).Align(lipgloss.Left)

	sb.WriteString(style.Render(buttonsRow))
	return styles.Base.Render(sb.String())
}

// Run runs the confirm field in accessible mode.
func (c *Confirm) Run() error {
	if c.accessible {
		return c.runAccessible()
	}
	return huh.Run(c)
}

// runAccessible runs the confirm field in accessible mode.
func (c *Confirm) runAccessible() error {
	styles := c.activeStyles()
	fmt.Println(styles.Title.Render(c.title.val))
	fmt.Println()
	c.accessor.Set(accessibility.PromptBool())
	fmt.Println(styles.SelectedOption.Render("Chose: "+c.String()) + "\n")
	return nil
}

func (c *Confirm) String() string {
	if c.accessor.Get() {
		return c.affirmative
	}
	return c.negative
}

// WithTheme sets the theme of the confirm field.
func (c *Confirm) WithTheme(theme *huh.Theme) huh.Field {
	if c.theme != nil {
		return c
	}
	c.theme = theme
	return c
}

// WithKeyMap sets the keymap of the confirm field.
func (c *Confirm) WithKeyMap(k *huh.KeyMap) huh.Field {
	c.keymap = k.Confirm
	return c
}

// WithAccessible sets the accessible mode of the confirm field.
func (c *Confirm) WithAccessible(accessible bool) huh.Field {
	c.accessible = accessible
	return c
}

// WithWidth sets the width of the confirm field.
func (c *Confirm) WithWidth(width int) huh.Field {
	c.width = width
	return c
}

// WithHeight sets the height of the confirm field.
func (c *Confirm) WithHeight(height int) huh.Field {
	c.height = height
	return c
}

// WithPosition sets the position of the confirm field.
func (c *Confirm) WithPosition(p huh.FieldPosition) huh.Field {
	c.keymap.Prev.SetEnabled(!p.IsFirst())
	c.keymap.Next.SetEnabled(!p.IsLast())
	c.keymap.Submit.SetEnabled(p.IsLast())
	return c
}

// GetKey returns the key of the field.
func (c *Confirm) GetKey() string {
	return c.key
}

// GetValue returns the value of the field.
func (c *Confirm) GetValue() any {
	return c.accessor.Get()
}
