package constants

import (
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap defines a set of key bindings for navigation and actions within an application or interface.
type KeyMap struct {
	Up            key.Binding
	Down          key.Binding
	FirstItem     key.Binding
	LastItem      key.Binding
	TogglePreview key.Binding
	OpenGithub    key.Binding
	Refresh       key.Binding
	PageDown      key.Binding
	PageUp        key.Binding
	NextSection   key.Binding
	PrevSection   key.Binding
	Help          key.Binding
	Quit          key.Binding
}

// WaitingIcon represents the icon for a waiting state.
// FailureIcon represents the icon for a failure state.
// SuccessIcon represents the icon for a success state.
// DraftIcon represents the icon for a draft state.
// BehindIcon represents the icon for a behind state.
// BlockedIcon represents the icon for a blocked state.
// MergedIcon represents the icon for a merged state.
// OpenIcon represents the icon for an open state.
// ClosedIcon represents the icon for a closed state.
const (
	WaitingIcon = ""
	FailureIcon = "󰅙"
	SuccessIcon = ""

	DraftIcon   = ""
	BehindIcon  = "󰇮"
	BlockedIcon = ""
	MergedIcon  = ""
	OpenIcon    = ""
	ClosedIcon  = ""
)
