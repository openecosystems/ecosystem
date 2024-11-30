package keys

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

// NewFormKeyMap returns a new form keymap.
func NewFormKeyMap() *huh.KeyMap {
	return &huh.KeyMap{
		Quit: key.NewBinding(key.WithKeys("ctrl+c")),
		Input: huh.InputKeyMap{
			AcceptSuggestion: key.NewBinding(key.WithKeys("ctrl+e"), key.WithHelp("ctrl+e", "complete")),
			Prev:             key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:             key.NewBinding(key.WithKeys("enter", "tab"), key.WithHelp("enter", "next")),
			Submit:           key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
		},
		FilePicker: huh.FilePickerKeyMap{
			GoToTop:  key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "first"), key.WithDisabled()),
			GoToLast: key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "last"), key.WithDisabled()),
			PageUp:   key.NewBinding(key.WithKeys("K", "pgup"), key.WithHelp("pgup", "page up"), key.WithDisabled()),
			PageDown: key.NewBinding(key.WithKeys("J", "pgdown"), key.WithHelp("pgdown", "page down"), key.WithDisabled()),
			Back:     key.NewBinding(key.WithKeys("h", "backspace", "left", "esc"), key.WithHelp("h", "back"), key.WithDisabled()),
			Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select"), key.WithDisabled()),
			Up:       key.NewBinding(key.WithKeys("up", "k", "ctrl+k", "ctrl+p"), key.WithHelp("↑", "up"), key.WithDisabled()),
			Down:     key.NewBinding(key.WithKeys("down", "j", "ctrl+j", "ctrl+n"), key.WithHelp("↓", "down"), key.WithDisabled()),

			Open:   key.NewBinding(key.WithKeys("l", "right", "enter"), key.WithHelp("enter", "open")),
			Close:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "close"), key.WithDisabled()),
			Prev:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:   key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next")),
			Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
		},
		Text: huh.TextKeyMap{
			Prev:    key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:    key.NewBinding(key.WithKeys("tab", "enter"), key.WithHelp("enter", "next")),
			Submit:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
			NewLine: key.NewBinding(key.WithKeys("alt+enter", "ctrl+j"), key.WithHelp("alt+enter / ctrl+j", "new line")),
			Editor:  key.NewBinding(key.WithKeys("ctrl+e"), key.WithHelp("ctrl+e", "open editor")),
		},
		Select: huh.SelectKeyMap{
			Prev:         key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:         key.NewBinding(key.WithKeys("enter", "tab"), key.WithHelp("enter", "select")),
			Submit:       key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
			Up:           key.NewBinding(key.WithKeys("up", "k", "ctrl+k", "ctrl+p"), key.WithHelp("↑", "up")),
			Down:         key.NewBinding(key.WithKeys("down", "j", "ctrl+j", "ctrl+n"), key.WithHelp("↓", "down")),
			Left:         key.NewBinding(key.WithKeys("h", "left"), key.WithHelp("←", "left"), key.WithDisabled()),
			Right:        key.NewBinding(key.WithKeys("l", "right"), key.WithHelp("→", "right"), key.WithDisabled()),
			Filter:       key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "filter")),
			SetFilter:    key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "set filter"), key.WithDisabled()),
			ClearFilter:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "clear filter"), key.WithDisabled()),
			HalfPageUp:   key.NewBinding(key.WithKeys("ctrl+u"), key.WithHelp("ctrl+u", "½ page up")),
			HalfPageDown: key.NewBinding(key.WithKeys("ctrl+d"), key.WithHelp("ctrl+d", "½ page down")),
			GotoTop:      key.NewBinding(key.WithKeys("home", "g"), key.WithHelp("g/home", "go to start")),
			GotoBottom:   key.NewBinding(key.WithKeys("end", "G"), key.WithHelp("G/end", "go to end")),
		},
		MultiSelect: huh.MultiSelectKeyMap{
			Prev:         key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:         key.NewBinding(key.WithKeys("enter", "tab"), key.WithHelp("enter", "confirm")),
			Submit:       key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
			Toggle:       key.NewBinding(key.WithKeys(" ", "x"), key.WithHelp("x", "toggle")),
			Up:           key.NewBinding(key.WithKeys("up", "k", "ctrl+p"), key.WithHelp("↑", "up")),
			Down:         key.NewBinding(key.WithKeys("down", "j", "ctrl+n"), key.WithHelp("↓", "down")),
			Filter:       key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "filter")),
			SetFilter:    key.NewBinding(key.WithKeys("enter", "esc"), key.WithHelp("esc", "set filter"), key.WithDisabled()),
			ClearFilter:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "clear filter"), key.WithDisabled()),
			HalfPageUp:   key.NewBinding(key.WithKeys("ctrl+u"), key.WithHelp("ctrl+u", "½ page up")),
			HalfPageDown: key.NewBinding(key.WithKeys("ctrl+d"), key.WithHelp("ctrl+d", "½ page down")),
			GotoTop:      key.NewBinding(key.WithKeys("home", "g"), key.WithHelp("g/home", "go to start")),
			GotoBottom:   key.NewBinding(key.WithKeys("end", "G"), key.WithHelp("G/end", "go to end")),
			SelectAll:    key.NewBinding(key.WithKeys("ctrl+a"), key.WithHelp("ctrl+a", "select all")),
			SelectNone:   key.NewBinding(key.WithKeys("ctrl+a"), key.WithHelp("ctrl+a", "select none"), key.WithDisabled()),
		},
		Note: huh.NoteKeyMap{
			Prev:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:   key.NewBinding(key.WithKeys("enter", "tab"), key.WithHelp("enter", "next")),
			Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
		},
		Confirm: huh.ConfirmKeyMap{
			Prev:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:   key.NewBinding(key.WithKeys("enter", "tab"), key.WithHelp("enter", "next")),
			Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
			Toggle: key.NewBinding(key.WithKeys("k", "j", "up", "down"), key.WithHelp("↑/↓", "toggle")),
			Accept: key.NewBinding(key.WithKeys("y", "Y"), key.WithHelp("y", "Yes")),
			Reject: key.NewBinding(key.WithKeys("n", "N"), key.WithHelp("n", "No")),
		},
	}
}
