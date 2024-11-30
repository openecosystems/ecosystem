package contract

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	"github.com/charmbracelet/bubbles/key"
)

type PageSettings struct {
	Title         string
	IsDefault     bool
	KeyBindings   []key.Binding
	ContentHeight int
	Type          config.PageType
}

type ContextAware interface {
	UpdateProgramContext(ctx *context.ProgramContext)
}

type Configurable interface {
}

type Responsive interface {
	SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext
	OnWindowSizeChanged(ctx *context.ProgramContext)
}

type Displayable interface {
	View() string
}

// Section represents the largest grouping of functionality. A Section can have multiple [Pages]
type Section interface {
	GetPages() []Page
	Configurable
	ContextAware
	Displayable
}

// Tabs are a special type of [Component] used to toggle between [Page]s
type Tabs interface {
	GetContentHeight() int
	Configurable
	ContextAware
	Displayable
}

// Page represents a single view of functionality within a [Section]. A [Page] is made up of multiple [Component]s
type Page interface {
	GetPageSettings() PageSettings
	Configurable
	ContextAware
	Responsive
	Displayable
}

// MainContent is a special type of [Component] that contains the main details of a [Page]
type MainContent interface {
	Configurable
	ContextAware
	Responsive
	Displayable
}

// Sidebar is a special type of [Component] that highlights details of a [Page]'s [Component]s
type Sidebar interface {
	IsOpen() bool
	Open()
	Close()
	Configurable
	ContextAware
	Responsive
	Displayable
}

// Footer is a special type of [Component] that shows different capabilities
type Footer interface {
	IsOpen() bool
	Configurable
	ContextAware
	Displayable
}

type Task interface{}
