package contract

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/context"

	"github.com/charmbracelet/bubbles/key"
)

// PageSettings defines the configuration and metadata for a page, including its title, default status, and other settings.
type PageSettings struct {
	Title         string
	IsDefault     bool
	KeyBindings   []key.Binding
	ContentHeight int
	Type          config.PageType
}

// ContextAware is an interface for components that can update or respond to changes in the ProgramContext.
// It provides the UpdateProgramContext method to receive the current program context instance.
type ContextAware interface {
	UpdateProgramContext(ctx *context.ProgramContext)
}

// Configurable represents an entity that allows customization or adjustment of its properties or behavior.
type Configurable interface{}

// Responsive defines an interface for components to handle UI responsiveness and dynamic resizing behavior.
type Responsive interface {
	SyncDimensions(ctx *context.ProgramContext) *context.ProgramContext
	OnWindowSizeChanged(ctx *context.ProgramContext)
}

// Displayable defines an interface for entities that can render or return a string representation for display purposes.
type Displayable interface {
	View() string
}

// Section defines an interface representing a categorized module containing pages, capable of being displayed and configured.
type Section interface {
	GetPages() []Page
	Configurable
	ContextAware
	Displayable
}

// Tabs represents an interface for managing and displaying tab-based UI components with context and configuration support.
// It combines Configurable, ContextAware, and Displayable interfaces to provide a flexible and dynamic tab system.
// The GetContentHeight method retrieves the height of the content within the tabs.
type Tabs interface {
	GetContentHeight() int
	Configurable
	ContextAware
	Displayable
}

// Page represents a configurable and responsive UI element with context awareness and display capabilities.
type Page interface {
	GetPageSettings() PageSettings
	Configurable
	ContextAware
	Responsive
	Displayable
}

// MainContent is an interface that combines Configurable, ContextAware, Responsive, and Displayable behaviors.
type MainContent interface {
	Configurable
	ContextAware
	Responsive
	Displayable
}

// Sidebar represents a UI component, ensuring open/close state handling, configurability, responsiveness, and display logic.
type Sidebar interface {
	IsOpen() bool
	Open()
	Close()
	Configurable
	ContextAware
	Responsive
	Displayable
}

// Footer represents an interface for footer components in the application.
// Footer provides functionalities for state checking, configuration, context updates, and display rendering.
type Footer interface {
	IsOpen() bool
	Configurable
	ContextAware
	Displayable
}

// Task represents a generic interface for tasks, used as a placeholder for various implementations.
type Task interface{}
