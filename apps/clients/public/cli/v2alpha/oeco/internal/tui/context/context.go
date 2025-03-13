package context

import (
	"github.com/charmbracelet/log"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	theme "apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
	cliv2alphalib "libs/public/go/cli/v2alpha"
)

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
	Settings *cliv2alphalib.Configuration
	Config   *config.Config
	Error    error

	// Bindings
	Logger *log.Logger

	// Stylistic
	Theme  theme.Theme
	Styles theme.Styles
}
