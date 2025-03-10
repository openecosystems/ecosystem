package packet

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"

	data "apps/clients/public/cli/v2alpha/oeco/internal/data"
	table "apps/clients/public/cli/v2alpha/oeco/internal/tui/components/table"
	context "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	utils "apps/clients/public/cli/v2alpha/oeco/internal/tui/utils"
)

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
