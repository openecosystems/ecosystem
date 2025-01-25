package keys

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap defines a set of key bindings and navigation mappings for user interaction in different application contexts.
type KeyMap struct {
	sectionType   config.SectionType
	pageType      config.PageType
	Up            key.Binding
	Down          key.Binding
	FirstLine     key.Binding
	LastLine      key.Binding
	TogglePreview key.Binding
	OpenGithub    key.Binding
	Refresh       key.Binding
	RefreshAll    key.Binding
	PageDown      key.Binding
	PageUp        key.Binding
	NextPage      key.Binding
	PrevPage      key.Binding
	Search        key.Binding
	CopyURL       key.Binding
	CopyNumber    key.Binding
	Help          key.Binding
	Quit          key.Binding
}

// CreateKeyMapForView creates and returns a KeyMap instance configured for the specified page type.
func CreateKeyMapForView(pageType config.PageType) help.KeyMap {
	Keys.pageType = pageType
	return Keys
}

// ShortHelp returns a slice of key bindings that provide concise help information.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

// FullHelp returns a two-dimensional slice of key bindings categorized by their context, such as navigation, app, section, and page.
func (k KeyMap) FullHelp() [][]key.Binding {
	var sectionKeys []key.Binding
	if k.sectionType == config.ConnectorSection {
		sectionKeys = ConnectorFullHelp()
	}

	var pageKeys []key.Binding
	if k.pageType == config.ConnectorDetailsPage {
		pageKeys = ConnectorDetailsPageFullHelp()
	}

	return [][]key.Binding{
		k.NavigationKeys(),
		k.AppKeys(),
		sectionKeys,
		pageKeys,
		k.QuitAndHelpKeys(),
	}
}

// NavigationKeys returns a slice of key bindings used for navigation actions such as moving up, down, paging, and line navigation.
func (k KeyMap) NavigationKeys() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
		k.PrevPage,
		k.NextPage,
		k.FirstLine,
		k.LastLine,
		k.PageDown,
		k.PageUp,
	}
}

// AppKeys returns a list of key bindings related to application-level actions, including refresh, search, and copy operations.
func (k KeyMap) AppKeys() []key.Binding {
	return []key.Binding{
		k.Refresh,
		k.RefreshAll,
		k.TogglePreview,
		k.OpenGithub,
		k.CopyNumber,
		k.CopyURL,
		k.Search,
	}
}

// QuitAndHelpKeys returns a slice of key bindings for the Help and Quit actions.
func (k KeyMap) QuitAndHelpKeys() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// Keys represents a predefined set of key bindings used for navigation, actions, and commands within the application.
var Keys = &KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	FirstLine: key.NewBinding(
		key.WithKeys("g", "home"),
		key.WithHelp("g/home", "first item"),
	),
	LastLine: key.NewBinding(
		key.WithKeys("G", "end"),
		key.WithHelp("G/end", "last item"),
	),
	TogglePreview: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "open in Preview"),
	),
	OpenGithub: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open in GitHub"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
	RefreshAll: key.NewBinding(
		key.WithKeys("R"),
		key.WithHelp("R", "refresh all"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("Ctrl+d", "preview page down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("ctrl+u"),
		key.WithHelp("Ctrl+u", "preview page up"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("󰁔/l", "next page"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("󰁍/h", "previous page"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
	CopyNumber: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "copy number"),
	),
	CopyURL: key.NewBinding(
		key.WithKeys("Y"),
		key.WithHelp("Y", "copy url"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
