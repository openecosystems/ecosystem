package table

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	listviewport "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/listviewport"
	constants "apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	theme "apps/clients/public/cli/v2alpha/oeco/internal/tui/theme"
)

// Model represents a table-like structure, managing columns, rows, viewport, and loading states for rendering and interaction.
type Model struct {
	ctx            context.ProgramContext
	Columns        []Column
	Rows           []Row
	EmptyState     *string
	loadingMessage string
	isLoading      bool
	loadingSpinner spinner.Model
	dimensions     constants.Dimensions
	rowsViewport   listviewport.Model
}

// Column represents the configuration for a single column in a table-like structure, including its visibility and dimensions.
type Column struct {
	Title         string
	Hidden        *bool
	Width         *int
	ComputedWidth int
	Grow          *bool
}

// Row represents a single row in a tabular data structure, defined as a slice of strings.
type Row []string

// NewModel initializes and returns a new Model instance with the provided context, dimensions, and data configuration.
func NewModel(
	ctx context.ProgramContext,
	dimensions constants.Dimensions,
	lastUpdated time.Time,
	createdAt time.Time,
	columns []Column,
	rows []Row,
	itemTypeLabel string,
	emptyState *string,
	loadingMessage string,
	isLoading bool,
) Model {
	itemHeight := 1
	if !ctx.Config.Theme.Tui.Table.Compact {
		itemHeight++
	}
	if ctx.Config.Theme.Tui.Table.ShowSeparator {
		itemHeight++
	}

	loadingSpinner := spinner.New()
	loadingSpinner.Spinner = spinner.Dot
	loadingSpinner.Style = lipgloss.NewStyle().Foreground(ctx.Theme.SecondaryText)

	return Model{
		ctx:            ctx,
		Columns:        columns,
		Rows:           rows,
		EmptyState:     emptyState,
		loadingMessage: loadingMessage,
		isLoading:      isLoading,
		loadingSpinner: loadingSpinner,
		dimensions:     dimensions,
		rowsViewport: listviewport.NewModel(
			ctx,
			dimensions,
			lastUpdated,
			createdAt,
			itemTypeLabel,
			len(rows),
			itemHeight,
		),
	}
}

// Update processes an incoming message, updates the model's state, and returns the updated model and a command to execute.
func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.isLoading {
		m.loadingSpinner, cmd = m.loadingSpinner.Update(msg)
	}
	return m, cmd
}

// StartLoadingSpinner initializes the loading spinner's tick action and returns a command to drive spinner updates.
func (m *Model) StartLoadingSpinner() tea.Cmd {
	return m.loadingSpinner.Tick
}

// View renders the table's complete view, combining the header and body, and returns it as a formatted string.
func (m *Model) View() string {
	header := m.renderHeader()
	body := m.renderBody()

	return lipgloss.JoinVertical(lipgloss.Left, header, body)
}

// SetIsLoading updates the isLoading state of the Model, indicating whether the loading process is active or not.
func (m *Model) SetIsLoading(isLoading bool) {
	m.isLoading = isLoading
}

// SetDimensions updates the dimensions of the model and synchronizes the rows viewport with the new dimensions.
func (m *Model) SetDimensions(dimensions constants.Dimensions) {
	m.dimensions = dimensions
	m.rowsViewport.SetDimensions(constants.Dimensions{
		Width:  m.dimensions.Width,
		Height: m.dimensions.Height,
	})
}

// ResetCurrItem resets the current item in the rows viewport to its default state.
func (m *Model) ResetCurrItem() {
	m.rowsViewport.ResetCurrItem()
}

// GetCurrItem retrieves the index of the currently selected item in the rows viewport.
func (m *Model) GetCurrItem() int {
	return m.rowsViewport.GetCurrItem()
}

// PrevItem moves the selection to the previous item in the viewport and updates the viewport content. Returns the current item index.
func (m *Model) PrevItem() int {
	currItem := m.rowsViewport.PrevItem()
	m.SyncViewPortContent()

	return currItem
}

// NextItem advances the current item position in the viewport and refreshes the viewport content. Returns the new item index.
func (m *Model) NextItem() int {
	currItem := m.rowsViewport.NextItem()
	m.SyncViewPortContent()

	return currItem
}

// FirstItem retrieves the index of the first item in the viewport and syncs the viewport content. It returns the current item index.
func (m *Model) FirstItem() int {
	currItem := m.rowsViewport.FirstItem()
	m.SyncViewPortContent()

	return currItem
}

// LastItem retrieves the index of the last visible item in the rows viewport and synchronizes the viewport content.
func (m *Model) LastItem() int {
	currItem := m.rowsViewport.LastItem()
	m.SyncViewPortContent()

	return currItem
}

// cacheColumnWidths calculates and stores the computed widths for visible table columns based on their rendered content.
func (m *Model) cacheColumnWidths() {
	columns := m.renderHeaderColumns()
	for i, col := range columns {
		if m.Columns[i].Hidden != nil && *m.Columns[i].Hidden {
			continue
		}
		m.Columns[i].ComputedWidth = lipgloss.Width(col)
	}
}

// SyncViewPortContent synchronizes the content of the viewport by rendering header columns and rows for display.
func (m *Model) SyncViewPortContent() {
	headerColumns := m.renderHeaderColumns()
	m.cacheColumnWidths()
	renderedRows := make([]string, 0, len(m.Rows))
	for i := range m.Rows {
		renderedRows = append(renderedRows, m.renderRow(i, headerColumns))
	}

	m.rowsViewport.SyncViewPort(
		lipgloss.JoinVertical(lipgloss.Left, renderedRows...),
	)
}

// SetRows updates the model's rows and synchronizes the viewport content for rendering.
func (m *Model) SetRows(rows []Row) {
	m.Rows = rows
	m.rowsViewport.SetNumItems(len(m.Rows))
	m.SyncViewPortContent()
}

// OnLineDown moves the current selection in the rows viewport to the next item.
func (m *Model) OnLineDown() {
	m.rowsViewport.NextItem()
}

// OnLineUp moves the viewport to the previous item in the rows list.
func (m *Model) OnLineUp() {
	m.rowsViewport.PrevItem()
}

// getShownColumns filters and returns all visible columns in the model by excluding columns marked as hidden.
func (m *Model) getShownColumns() []Column {
	shownColumns := make([]Column, 0, len(m.Columns))
	for _, col := range m.Columns {
		if col.Hidden != nil && *col.Hidden {
			continue
		}

		shownColumns = append(shownColumns, col)
	}
	return shownColumns
}

// renderHeaderColumns generates and returns a slice of strings representing the rendered header columns of a table.
func (m *Model) renderHeaderColumns() []string {
	shownColumns := m.getShownColumns()
	renderedColumns := make([]string, len(shownColumns))
	takenWidth := 0
	numGrowingColumns := 0
	for i, column := range shownColumns {
		if column.Grow != nil && *column.Grow {
			numGrowingColumns++
			continue
		}

		if column.Width != nil {
			renderedColumns[i] = m.ctx.Styles.Table.TitleCellStyle.
				Width(*column.Width).
				MaxWidth(*column.Width).
				Render(column.Title)
			takenWidth += *column.Width
			continue
		}

		cell := m.ctx.Styles.Table.TitleCellStyle.Render(column.Title)
		renderedColumns[i] = cell
		takenWidth += lipgloss.Width(cell)
	}

	if numGrowingColumns == 0 {
		return renderedColumns
	}

	leftoverWidth := m.dimensions.Width - takenWidth
	growCellWidth := leftoverWidth / numGrowingColumns
	for i, column := range shownColumns {
		if column.Grow == nil || !*column.Grow {
			continue
		}

		renderedColumns[i] = m.ctx.Styles.Table.TitleCellStyle.
			Width(growCellWidth).
			MaxWidth(growCellWidth).
			Render(column.Title)
	}

	return renderedColumns
}

// renderHeader generates and renders the header row of a table with styled and dimensioned header columns.
func (m *Model) renderHeader() string {
	headerColumns := m.renderHeaderColumns()
	header := lipgloss.JoinHorizontal(lipgloss.Top, headerColumns...)
	return m.ctx.Styles.Table.HeaderStyle.
		Width(m.dimensions.Width).
		MaxWidth(m.dimensions.Width).
		Height(theme.TableHeaderHeight).
		MaxHeight(theme.TableHeaderHeight).
		Render(header)
}

// renderBody generates a string representing the body of the model, considering loading states, empty state, or row content.
func (m *Model) renderBody() string {
	bodyStyle := lipgloss.NewStyle().
		Height(m.dimensions.Height).
		MaxWidth(m.dimensions.Width)

	if m.isLoading {
		return lipgloss.Place(
			m.dimensions.Width,
			m.dimensions.Height,
			lipgloss.Center,
			lipgloss.Center,
			fmt.Sprintf("%s%s", m.loadingSpinner.View(), m.loadingMessage),
		)
	}

	if len(m.Rows) == 0 && m.EmptyState != nil {
		return bodyStyle.Render(*m.EmptyState)
	}

	return m.rowsViewport.View()
}

// renderRow generates a styled string representing a single row in the table, considering column visibility and selection.
func (m *Model) renderRow(rowID int, headerColumns []string) string {
	var style lipgloss.Style

	if m.rowsViewport.GetCurrItem() == rowID {
		style = m.ctx.Styles.Table.SelectedCellStyle
	} else {
		style = m.ctx.Styles.Table.CellStyle
	}

	renderedColumns := make([]string, 0, len(m.Columns))
	headerColID := 0

	for i, column := range m.Columns {
		if column.Hidden != nil && *column.Hidden {
			continue
		}

		colWidth := lipgloss.Width(headerColumns[headerColID])
		colHeight := 1
		if !m.ctx.Config.Theme.Tui.Table.Compact {
			colHeight = 2
		}
		col := m.Rows[rowID][i]
		renderedCol := style.
			Width(colWidth).
			MaxWidth(colWidth).
			Height(colHeight).
			MaxHeight(colHeight).
			Render(col)

		renderedColumns = append(renderedColumns, renderedCol)
		headerColID++
	}

	return m.ctx.Styles.Table.RowStyle.
		BorderBottom(m.ctx.Config.Theme.Tui.Table.ShowSeparator).
		MaxWidth(m.dimensions.Width).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedColumns...))
}

// UpdateProgramContext updates the program context for the model and propagates it to the rows viewport.
func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = *ctx
	m.rowsViewport.UpdateProgramContext(ctx)
}

// LastUpdated retrieves the last updated timestamp of the rows viewport in the model.
func (m *Model) LastUpdated() time.Time {
	return m.rowsViewport.LastUpdated
}

// CreatedAt returns the timestamp of when the underlying rowsViewport data was created.
func (m *Model) CreatedAt() time.Time {
	return m.rowsViewport.CreatedAt
}

// UpdateLastUpdated updates the LastUpdated field of the rowsViewport with the provided time value.
func (m *Model) UpdateLastUpdated(t time.Time) {
	m.rowsViewport.LastUpdated = t
}

// UpdateTotalItemsCount updates the total number of items in the rows viewport with the given count.
func (m *Model) UpdateTotalItemsCount(count int) {
	m.rowsViewport.SetTotalItems(count)
}

// IsLoading returns a boolean indicating whether the model is currently in a loading state.
func (m *Model) IsLoading() bool {
	return m.isLoading
}
