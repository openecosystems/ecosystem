package contract

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/components/table"
	"apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"

	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
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

// Component represents a general interface for reusable and composable UI or program components.
type Component interface {
	tea.Model
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
	GetPageSettings() *PageSettings
	Configurable
	ContextAware
	Responsive
	Displayable
	tea.Model
}

// Dashboard represents a configurable and responsive visual Page
type Dashboard interface {
	Page
	GetItemSingularForm() string
	GetItemPluralForm() string
	GetTotalCount() *int
}

// MainContent is an interface that combines Configurable, ContextAware, Responsive, and Displayable behaviors.
type MainContent interface {
	Configurable
	ContextAware
	Responsive
	Displayable
	tea.Model
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
	Component
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
type Task interface {
	IsRunning() bool
}

// Table represents a tabular data structure interface for managing and navigating rows within a table.
// NumRows returns the total number of rows in the table.
// GetCurrRow retrieves the current row's data in the form of a RowData interface.
// CurrRow returns the index of the current row.
// NextRow moves to and returns the index of the next row.
// PrevRow moves to and returns the index of the previous row.
// FirstItem returns the index of the first row in the table.
// LastItem returns the index of the last row in the table.
// FetchNextPageSectionRows fetches additional rows for the next page/section using a slice of commands.
// BuildRows constructs and returns all rows in the table as a slice of table.Row.
// ResetRows clears or resets the existing rows in the table.
// IsLoading determines if the table is currently in a loading state.
// SetIsLoading sets the loading state of the table to the given boolean value.
type Table interface {
	// NumRows() int
	// GetCurrRow() data.RowData
	CurrRow() int
	NextRow() int
	PrevRow() int
	FirstItem() int
	LastItem() int
	// FetchNextPageSectionRows() []tea.Cmd
	BuildRows() []table.Row
	ResetRows()
	IsLoading() bool
	SetIsLoading(val bool)
	SetDimensions(dimensions constants.Dimensions)
	UpdateProgramContext(ctx *context.ProgramContext)
	SyncViewPortContent()
	UpdateTotalItemsCount(count int)
	GetRows() []table.Row
	tea.Model
}
