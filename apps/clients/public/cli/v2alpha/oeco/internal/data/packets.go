package data

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// PacketData represents the data structure containing information about a specific issue or item in the ecosystem.
type PacketData struct {
	SrcIP    string
	DstIP    string
	Protocol string
	Length   int

	Number int
	Title  string
	Body   string
	State  string
	Author struct {
		Login string
	}
	CreatedAt time.Time
	UpdatedAt time.Time

	URL        string
	Repository Repository
	Assignees  Assignees      `graphql:"assignees(first: 3)"`
	Comments   IssueComments  `graphql:"comments(first: 15)"`
	Reactions  IssueReactions `graphql:"reactions(first: 1)"`
	Labels     IssueLabels    `graphql:"labels(first: 3)"`
}

// PacketResponse represents the response containing a collection of PacketData and the total count of packets.
type PacketResponse struct {
	Packets    []PacketData
	TotalCount int
}

// GetTitle returns the title of the PacketData.
func (data PacketData) GetTitle() string {
	return data.Title
}

// GetRepoNameWithOwner returns the repository name along with the owner's name from the PacketData structure.
func (data PacketData) GetRepoNameWithOwner() string {
	return data.Repository.NameWithOwner
}

// GetNumber returns the number associated with the PacketData instance.
func (data PacketData) GetNumber() int {
	return data.Number
}

// GetURL returns the Url field from the PacketData structure.
func (data PacketData) GetURL() string {
	return data.URL
}

// GetUpdatedAt returns the time when the PacketData was last updated.
func (data PacketData) GetUpdatedAt() time.Time {
	return data.UpdatedAt
}

// makeIssuesQuery formats a GitHub issues search query string by appending "is:issue" and "sort:updated" to the input query.
//
//nolint:unused
func makeIssuesQuery(query string) string {
	return fmt.Sprintf("is:issue %s sort:updated", query)
}

//
//func streamPackets(_ string, _ int, _ *PageInfo) (PacketResponse, error) {
//	packets := make([]PacketData, 0, 1)
//	packets = append(packets, PacketData{
//		Number: 0,
//		Title:  "Hello World",
//		Body:   "Test Bod",
//		State:  "Active",
//		Author: struct{ Login string }{
//			Login: "12345",
//		},
//		UpdatedAt: time.Now(),
//		URL:       "https://openecosystems.com",
//		Repository: Repository{
//			Name:          "test",
//			NameWithOwner: "world",
//			IsArchived:    false,
//		},
//		Assignees: Assignees{
//			Nodes: nil,
//		},
//		Comments: IssueComments{
//			Nodes:      nil,
//			TotalCount: 0,
//		},
//		Reactions: IssueReactions{
//			TotalCount: 0,
//		},
//		Labels: IssueLabels{
//			Nodes: nil,
//		},
//	})
//
//	return PacketResponse{
//		Packets:    packets,
//		TotalCount: 1,
//	}, nil
//}

// ListenForPackets Start packet capturing in a goroutine and send packets through the channel
func ListenForPackets(device string, sub chan PacketData) (tea.Cmd, error) {
	var err error

	return func() tea.Msg {
		// TODO: Uncomment and fix this
		//handle, _err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
		//if _err != nil {
		//	err = _err
		//}
		//
		//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		//
		//// Start a goroutine to send packets into the channel
		//go func() {
		//	for packet := range packetSource.Packets() {
		//		networkLayer := packet.NetworkLayer()
		//		if networkLayer == nil {
		//			continue
		//		}
		//
		//		src, dst := networkLayer.NetworkFlow().Endpoints()
		//		protocol := ""
		//		if packet.TransportLayer() != nil {
		//			protocol = packet.TransportLayer().LayerType().String()
		//		}
		//
		//		// Send packet data to the channel
		//		sub <- PacketData{
		//			SrcIP:    src.String(),
		//			DstIP:    dst.String(),
		//			Protocol: protocol,
		//			Length:   len(packet.Data()),
		//
		//			Title: src.String(),
		//		}
		//	}
		//}()

		return nil
	}, err
}

// WaitForPacket Command that waits for a new packet from the channel
func WaitForPacket(sub chan PacketData) tea.Cmd {
	return func() tea.Msg {
		return <-sub // Wait for and return the next packet
	}
}
