package context

import (
	"github.com/charmbracelet/log"

	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	theme "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"
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
	Settings *sdkv2betalib.CLIConfiguration
	Config   *config.Config
	Error    error

	// Bindings
	Logger *log.Logger

	// Stylistic
	Theme  theme.Theme
	Styles theme.Styles
}
