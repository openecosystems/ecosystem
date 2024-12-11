package context

import (
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"time"

	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"

	tea "github.com/charmbracelet/bubbletea"
)

type State = int

const (
	TaskStart State = iota
	TaskFinished
	TaskError
)

type Task struct {
	Id           string
	StartText    string
	FinishedText string
	State        State
	Error        error
	StartTime    time.Time
	FinishedTime *time.Time
}

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
