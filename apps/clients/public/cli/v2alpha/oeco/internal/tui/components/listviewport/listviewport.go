package listviewport

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"

	constants "apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	utils "apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
)

// Model represents the main model structure managing application context, viewport, and item-related information.
type Model struct {
	ctx             *context.ProgramContext
	viewport        viewport.Model
	topBoundID      int
	bottomBoundID   int
	currID          int
	ListItemHeight  int
	NumCurrentItems int
	NumTotalItems   int
	LastUpdated     time.Time
	CreatedAt       time.Time
	ItemTypeLabel   string
}

// NewModel initializes and returns a new Model instance with the provided context, dimensions, timestamps, and other parameters.
func NewModel(
	ctx *context.ProgramContext,
	dimensions constants.Dimensions,
	lastUpdated time.Time,
	createdAt time.Time,
	itemTypeLabel string,
	numItems, listItemHeight int,
) Model {
	model := Model{
		ctx:             ctx,
		NumCurrentItems: numItems,
		ListItemHeight:  listItemHeight,
		currID:          0,
		viewport: viewport.Model{
			Width:  dimensions.Width,
			Height: dimensions.Height,
		},
		topBoundID:    0,
		ItemTypeLabel: itemTypeLabel,
		LastUpdated:   lastUpdated,
		CreatedAt:     createdAt,
	}
	model.bottomBoundID = utils.Min(
		model.NumCurrentItems-1,
		model.getNumPrsPerPage()-1,
	)
	return model
}

// SetNumItems sets the current number of items and updates the bottom bound ID based on the viewport's capacity.
func (m *Model) SetNumItems(numItems int) {
	m.NumCurrentItems = numItems
	m.bottomBoundID = utils.Min(m.NumCurrentItems-1, m.getNumPrsPerPage()-1)
}

// SetTotalItems updates the total number of items in the model by setting the NumTotalItems field to the specified value.
func (m *Model) SetTotalItems(total int) {
	m.NumTotalItems = total
}

// SyncViewPort updates the content of the viewport with the provided string argument.
func (m *Model) SyncViewPort(content string) {
	m.viewport.SetContent(content)
}

// getNumPrsPerPage calculates and returns the maximum number of items that can be displayed on a single page.
func (m *Model) getNumPrsPerPage() int {
	if m.ListItemHeight == 0 {
		return 0
	}
	return m.viewport.Height / m.ListItemHeight
}

// ResetCurrItem resets the current item identifier to 0 and scrolls the viewport to the top.
func (m *Model) ResetCurrItem() {
	m.currID = 0
	m.viewport.GotoTop()
}

// GetCurrItem retrieves the current item's identifier from the model.
func (m *Model) GetCurrItem() int {
	return m.currID
}

// NextItem moves the current item selection to the next item in the list and adjusts the viewport if needed.
// It returns the updated current item ID after the operation.
func (m *Model) NextItem() int {
	atBottomOfViewport := m.currID >= m.bottomBoundID
	if atBottomOfViewport {
		m.topBoundID++
		m.bottomBoundID++
		m.viewport.LineDown(m.ListItemHeight)
	}

	newID := utils.Min(m.currID+1, m.NumCurrentItems-1)
	newID = utils.Max(newID, 0)
	m.currID = newID
	return m.currID
}

// PrevItem adjusts the current item index to the previous item, updating viewport bounds if at the top of the viewport.
func (m *Model) PrevItem() int {
	atTopOfViewport := m.currID < m.topBoundID
	if atTopOfViewport {
		m.topBoundID--
		m.bottomBoundID--
		m.viewport.LineUp(m.ListItemHeight)
	}

	m.currID = utils.Max(m.currID-1, 0)
	return m.currID
}

// FirstItem resets the current item to the first in the list and moves the viewport to the top. Returns the updated current item ID.
func (m *Model) FirstItem() int {
	m.currID = 0
	m.viewport.GotoTop()
	return m.currID
}

// LastItem sets the current item index to the last item and moves the viewport to the bottom. Returns the updated index.
func (m *Model) LastItem() int {
	m.currID = m.NumCurrentItems - 1
	m.viewport.GotoBottom()
	return m.currID
}

// SetDimensions updates the viewport's width and height based on the provided dimensions.
func (m *Model) SetDimensions(dimensions constants.Dimensions) {
	m.viewport.Height = dimensions.Height
	m.viewport.Width = dimensions.Width
}

// View renders the current state of the viewport as a styled string representation.
func (m *Model) View() string {
	vp := m.viewport.View()
	return lipgloss.NewStyle().
		Width(m.viewport.Width).
		MaxWidth(m.viewport.Width).
		Render(
			vp,
		)
}

// UpdateProgramContext updates the program's context by setting a new ProgramContext instance to the Model.
func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
