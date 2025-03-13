package packet

import (
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	data "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/data"
	table "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/components/table"
	config "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/config"
	constants "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/constants"
	context "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	contract "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/contract"
	utils "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
)

const device = "lo0"

// cellWidth represents the default width for table cells as an integer. It is used to configure column dimensions in tables.
var cellWidth = 6

// Model represents a structure that embeds another Model from the table package, providing extended functionalities.
type Model struct {
	*table.Model

	Packets []data.PacketData

	PacketChannel chan data.PacketData // Channel to receive packets
}

// NewModelOptions defines configuration options for initializing a new table model, including dimensions, data, and states.
type NewModelOptions struct {
	Dimensions     constants.Dimensions
	LastUpdated    time.Time
	CreatedAt      time.Time
	Rows           []table.Row
	ItemTypeLabel  string
	EmptyState     *string
	LoadingMessage string
	IsLoading      bool
}

// NewModel creates a new instance of the Model, embedding a table model initialized with the given context and options.
func NewModel(ctx *context.ProgramContext, options *NewModelOptions) contract.Table {
	limit := 20
	cfg := config.PacketContainerConfig{
		Title:   "Ecosystem Packets",
		Filters: "",
		Limit:   &limit,
		Layout:  ctx.Config.Defaults.Layout.Packets,
	}
	t := table.NewModel(
		ctx,
		options.Dimensions,
		options.LastUpdated,
		options.CreatedAt,
		GetSectionColumns(ctx, cfg),
		nil,
		options.ItemTypeLabel,
		options.EmptyState,
		options.LoadingMessage,
		options.IsLoading,
	)

	m := &Model{
		Model:         &t,
		Packets:       []data.PacketData{},
		PacketChannel: make(chan data.PacketData),
	}

	return m
}

// Init initializes the Model and returns a tea.Cmd batch for further processing or updates.
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	packetsCmd, err := data.ListenForPackets(device, m.PacketChannel)
	if err != nil {
		m.Ctx.Logger.Error(err)
	}

	cmds = append(cmds,
		packetsCmd,
		data.WaitForPacket(m.PacketChannel), // Wait for the first packet
	)
	return tea.Batch(cmds...)
}

// Update processes the incoming message, updates the model's state, and returns the updated model along with commands.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	_, cmd = m.Model.Update(msg)

	cmds = append(cmds,
		data.WaitForPacket(m.PacketChannel), // Wait for the first packet
	)

	switch message := msg.(type) {
	case data.PacketData:
		m.Packets = append(m.Packets, message)

		// Keep only the last StreamMaxRecords
		if len(m.Packets) > m.Ctx.Config.Defaults.StreamMaxRecordsToRetain {
			m.Packets = m.Packets[len(m.Packets)-m.Ctx.Config.Defaults.StreamMaxRecordsToRetain:]
		}

		m.SetIsLoading(false)
		m.SetRows(m.BuildRows())

		// Wait for the next packet
		return m, data.WaitForPacket(m.PacketChannel)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the current state of the model as a string for display in the terminal. It adapts based on the form state.
func (m *Model) View() string {
	return m.Model.View()
}

// BuildRows converts the packets contained in the model into rows for a table by processing each packet's data.
func (m *Model) BuildRows() []table.Row {
	var rows []table.Row
	for _, currPacket := range m.Packets {
		packetModel := NewPacket(m.Model.Ctx, currPacket)
		rows = append(rows, packetModel.ToTableRow())
	}

	if rows == nil {
		rows = []table.Row{}
	}

	return rows
}

// Packet represents a data unit that contains context and its associated payload for processing.
type Packet struct {
	Ctx  *context.ProgramContext
	Data data.PacketData
}

// NewPacket initializes and returns a new Packet instance using the provided program context and packet data.
func NewPacket(ctx *context.ProgramContext, data data.PacketData) *Packet {
	return &Packet{
		Ctx:  ctx,
		Data: data,
	}
}

// ToTableRow converts the Packet data and rendering logic into a table.Row structure for display in a table format.
func (packet *Packet) ToTableRow() table.Row {
	return table.Row{
		packet.renderStatus(),
		packet.renderRepoName(),
		packet.renderTitle(),
		packet.renderOpenedBy(),
		packet.renderAssignees(),
		packet.renderNumComments(),
		packet.renderNumReactions(),
		packet.renderUpdateAt(),
		packet.renderCreatedAt(),
	}
}

func (packet *Packet) getTextStyle() lipgloss.Style {
	return packet.Ctx.Styles.Table.RowStyle
}

func (packet *Packet) renderUpdateAt() string {
	timeFormat := packet.Ctx.Config.Defaults.DateFormat

	updatedAtOutput := ""
	if timeFormat == "" || timeFormat == "relative" {
		updatedAtOutput = utils.TimeElapsed(packet.Data.UpdatedAt)
	} else {
		updatedAtOutput = packet.Data.UpdatedAt.Format(timeFormat)
	}

	return packet.getTextStyle().Render(updatedAtOutput)
}

func (packet *Packet) renderCreatedAt() string {
	timeFormat := packet.Ctx.Config.Defaults.DateFormat

	createdAtOutput := ""
	if timeFormat == "" || timeFormat == "relative" {
		createdAtOutput = utils.TimeElapsed(packet.Data.CreatedAt)
	} else {
		createdAtOutput = packet.Data.CreatedAt.Format(timeFormat)
	}

	return packet.getTextStyle().Render(createdAtOutput)
}

func (packet *Packet) renderRepoName() string {
	repoName := packet.Data.Repository.Name
	return packet.getTextStyle().Render(repoName)
}

func (packet *Packet) renderTitle() string {
	return RenderTitle(packet.Ctx, packet.Data.State, packet.Data.Title, packet.Data.Number)
}

func (packet *Packet) renderOpenedBy() string {
	return packet.getTextStyle().Render(packet.Data.Author.Login)
}

func (packet *Packet) renderAssignees() string {
	assignees := make([]string, 0, len(packet.Data.Assignees.Nodes))
	for _, assignee := range packet.Data.Assignees.Nodes {
		assignees = append(assignees, assignee.Login)
	}
	return packet.getTextStyle().Render(strings.Join(assignees, ","))
}

func (packet *Packet) renderStatus() string {
	if packet.Data.State == "OPEN" {
		return lipgloss.NewStyle().Foreground(packet.Ctx.Theme.PrimaryColor500).Render("")
	}

	return packet.getTextStyle().Render("")
}

func (packet *Packet) renderNumComments() string {
	return packet.getTextStyle().Render(strconv.Itoa(packet.Data.Comments.TotalCount))
}

func (packet *Packet) renderNumReactions() string {
	return packet.getTextStyle().Render(strconv.Itoa(packet.Data.Reactions.TotalCount))
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
