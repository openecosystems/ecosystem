package ecosystemdashboardcontent

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	data "apps/clients/public/cli/v2alpha/oeco/internal/data"
	packet "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/packet"
	table "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/table"
	config "apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	content "apps/clients/public/cli/v2alpha/oeco/internal/tui/content"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	utils "apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
)

// PageType defines the type of page as "dashboard" to differentiate it from other potential page types in the system.
const (
	PageType = "dashboard"
)

// cellWidth represents the default width for table cells as an integer. It is used to configure column dimensions in tables.
var cellWidth = 6

// Model represents a structure combining a dashboard base model with additional issue data for ecosystems.
type Model struct {
	*content.DashboardBaseModel

	Packets []data.PacketData

	packetChannel chan data.PacketData // Channel to receive packets
}

// NewModel initializes and returns a new Model with the provided context.
func NewModel(ctx *context.ProgramContext) *Model {
	viewportModel := viewport.New(ctx.MainContentBodyWidth, ctx.MainContentBodyHeight)
	limit := 20
	cfg := config.PacketContainerConfig{
		Title:   "Ecosystem Packets",
		Filters: "",
		Limit:   &limit,
		Layout:  ctx.Config.Defaults.Layout.Packets,
	}

	model := &Model{}
	model.DashboardBaseModel = content.NewDashboardBaseModel(
		ctx,
		&content.NewDashboardBaseOptions{
			NewBaseOptions: &content.NewBaseOptions{
				Viewport: &viewportModel,
			},
			ID:            0,
			Singular:      "packet",
			Plural:        "packets",
			Type:          PageType,
			SearchFilters: "",
			Columns:       GetSectionColumns(ctx, cfg),
			LastUpdated:   time.Now(),
			CreatedAt:     time.Now(),
		},
	)
	model.Packets = []data.PacketData{}
	model.packetChannel = make(chan data.PacketData)

	return model
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	packetsCmd, err := data.ListenForPackets("en0", m.packetChannel)
	if err != nil {
		m.DashboardBaseModel.Ctx.Logger.Error(err)
	}

	cmds = append(cmds,
		m.InitBase(),
		packetsCmd,
		data.WaitForPacket(m.packetChannel), // Wait for the first packet
	)

	return tea.Batch(cmds...)
}

// Update processes a message, updates the model state, and returns the updated model along with a command to be executed.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		baseCmd     tea.Cmd
		cmds        []tea.Cmd
		viewportCmd tea.Cmd
	)

	m.Ctx.Logger.Debug("Content - Ecosystem Dashboard Content - Update", msg)

	m.BaseModel, baseCmd = m.DashboardBaseModel.UpdateBase(msg)

	m.DashboardBaseModel.Viewport.SetContent(m.Table.View())
	v, c := m.DashboardBaseModel.Viewport.Update(msg)
	m.DashboardBaseModel.Viewport = &v
	viewportCmd = c

	cmds = append(
		cmds,
		baseCmd,
		viewportCmd,
	)

	switch message := msg.(type) {
	case tea.KeyMsg:

		if m.DashboardBaseModel.IsSearchFocused() {
			switch {
			case message.Type == tea.KeyCtrlC, message.Type == tea.KeyEsc:
				m.SearchBar.SetValue(m.SearchValue)
				blinkCmd := m.SetIsSearching(false)
				return m, blinkCmd

				// case message.Type == tea.KeyEnter:
				//	m.SearchValue = m.SearchBar.Value()
				//	m.SetIsSearching(false)
				//	m.ResetRows()
				//	return m, tea.Batch(m.FetchNextPageSectionRows()...)
			}

			break
		}

		if m.IsPromptConfirmationFocused() {
			switch {
			case message.Type == tea.KeyCtrlC, message.Type == tea.KeyEsc:
				m.PromptConfirmationBox.Reset()
				baseCmd = m.SetIsPromptConfirmationShown(false)
				return m, baseCmd

			case message.Type == tea.KeyEnter:
				input := m.PromptConfirmationBox.Value()
				action := m.GetPromptConfirmationAction()
				if input == "Y" || input == "y" {
					switch action {
					case "details":
						baseCmd = m.details()
					}
				}

				m.PromptConfirmationBox.Reset()
				blinkCmd := m.SetIsPromptConfirmationShown(false)

				return m, tea.Batch(baseCmd, blinkCmd)
			}
			break
		}
	case data.PacketData:
		m.Packets = append(m.Packets, message)

		// Keep only the last StreamMaxRecords
		if len(m.Packets) > m.DashboardBaseModel.Ctx.Config.Defaults.StreamMaxRecordsToRetain {
			m.Packets = m.Packets[len(m.Packets)-m.DashboardBaseModel.Ctx.Config.Defaults.StreamMaxRecordsToRetain:]
		}

		m.DashboardBaseModel.Ctx.Logger.Debug("LLLLLLLLLLContent - Ecosystem Dashboard Content - Update: Packet received")

		m.Table.SetIsLoading(false)
		m.Table.SetRows(m.BuildRows())

		// Wait for the next packet
		return m, data.WaitForPacket(m.packetChannel)
	}

	search, searchCmd := m.SearchBar.Update(msg)
	m.SearchBar = search

	prompt, promptCmd := m.PromptConfirmationBox.Update(msg)
	m.PromptConfirmationBox = prompt

	t, tableCmd := m.Table.Update(msg)
	m.Table = t

	cmds = append(cmds, baseCmd, searchCmd, promptCmd, tableCmd)

	return m, tea.Batch(cmds...)
}

// View returns the rendered string representation of the BaseModel by applying contextual styles and joining content vertically.
func (m *Model) View() string {
	//s := "\nReal-Time Mesh Packet Analyzer\n" + "---------------------------------\n"
	//
	//// Show captured packets
	//for i, p := range m.Packets {
	//	s += fmt.Sprintf("%2d: %s -> %s | %s | %d bytes\n",
	//		i+1, p.SrcIP, p.DstIP, p.Protocol, p.Length)
	//}
	//
	//return s
	return m.ViewBase()
}

// GetSectionColumns generates a slice of table columns configured based on default and section-specific layout settings.
func GetSectionColumns(ctx *context.ProgramContext, cfg config.PacketContainerConfig) []table.Column {
	dLayout := ctx.Config.Defaults.Layout.Packets
	sLayout := cfg.Layout

	updatedAtLayout := config.MergeColumnConfigs(
		dLayout.UpdatedAt,
		sLayout.UpdatedAt,
	)
	createdAtLayout := config.MergeColumnConfigs(
		dLayout.CreatedAt,
		sLayout.CreatedAt,
	)
	stateLayout := config.MergeColumnConfigs(dLayout.State, sLayout.State)
	repoLayout := config.MergeColumnConfigs(dLayout.Repo, sLayout.Repo)
	titleLayout := config.MergeColumnConfigs(dLayout.Title, sLayout.Title)
	creatorLayout := config.MergeColumnConfigs(dLayout.Creator, sLayout.Creator)
	assigneesLayout := config.MergeColumnConfigs(
		dLayout.Assignees,
		sLayout.Assignees,
	)
	commentsLayout := config.MergeColumnConfigs(
		dLayout.Comments,
		sLayout.Comments,
	)
	reactionsLayout := config.MergeColumnConfigs(
		dLayout.Reactions,
		sLayout.Reactions,
	)

	return []table.Column{
		{
			Title:  "",
			Width:  stateLayout.Width,
			Hidden: stateLayout.Hidden,
		},
		{
			Title:  "",
			Width:  repoLayout.Width,
			Hidden: repoLayout.Hidden,
		},
		{
			Title:  "Title",
			Grow:   utils.BoolPtr(true),
			Hidden: titleLayout.Hidden,
		},
		{
			Title:  "Creator",
			Width:  creatorLayout.Width,
			Hidden: creatorLayout.Hidden,
		},
		{
			Title:  "Assignees",
			Width:  assigneesLayout.Width,
			Hidden: assigneesLayout.Hidden,
		},
		{
			Title:  "",
			Width:  &cellWidth,
			Hidden: commentsLayout.Hidden,
		},
		{
			Title:  "",
			Width:  &cellWidth,
			Hidden: reactionsLayout.Hidden,
		},
		{
			Title:  "󱦻",
			Width:  updatedAtLayout.Width,
			Hidden: updatedAtLayout.Hidden,
		},
		{
			Title:  "󱡢",
			Width:  createdAtLayout.Width,
			Hidden: createdAtLayout.Hidden,
		},
	}
}

// BuildRows converts the packets contained in the model into rows for a table by processing each packet's data.
func (m *Model) BuildRows() []table.Row {
	var rows []table.Row
	for _, currPacket := range m.Packets {
		packetModel := packet.NewPacket(m.DashboardBaseModel.Ctx, currPacket)
		rows = append(rows, packetModel.ToTableRow())
	}

	if rows == nil {
		rows = []table.Row{}
	}

	return rows
}
