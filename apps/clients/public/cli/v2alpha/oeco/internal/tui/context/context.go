package context

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// State is an alias for the int type, commonly used to represent various states or stages in a process.
type State = int

// TaskStart represents the initial state of a task.
// TaskFinished represents the state when a task is successfully completed.
// TaskError represents the state when a task encounters an error.
const (
	TaskStart State = iota
	TaskFinished
	TaskError
)

// Task represents a unit of work with its associated metadata and lifecycle states.
type Task struct {
	Id           string
	StartText    string
	FinishedText string
	State        State
	Error        error
	StartTime    time.Time
	FinishedTime *time.Time
}

// ProgramContext encapsulates the application UI's state, configuration, and behavior for rendering and interaction.
type ProgramContext struct {
	// Identified
	User string

	// Responsive
	ScreenHeight             int
	ScreenWidth              int
	PageContentWidth         int
	PageContentHeight        int
	MainContentWidth         int
	MainContentHeight        int
	MainContentBodyWidth     int
	MainContentBodyHeight    int
	SidebarContentWidth      int
	SidebarContentHeight     int
	SidebarContentBodyWidth  int
	SidebarContentBodyHeight int

	// Configurable
	Section  config.SectionType
	Page     config.PageType
	Settings *specv2pb.SpecSettings
	Config   *config.Config
	Error    error

	// Executable
	StartTask func(task Task) tea.Cmd

	// Stylistic
	Theme  theme.Theme
	Styles theme.Styles
}
