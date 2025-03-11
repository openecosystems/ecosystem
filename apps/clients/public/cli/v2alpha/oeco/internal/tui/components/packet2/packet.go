package main

import (
	"apps/clients/public/cli/v2alpha/oeco/internal/data"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Model for Bubble Tea
type model struct {
	sub      chan data.PacketData // Channel to receive packets
	packets  []data.PacketData    // Store captured packets
	device   string               // Network interface
	quitting bool                 // Quit flag
}

func (m model) Init() tea.Cmd {
	packets, err := data.ListenForPackets(m.device, m.sub)
	if err != nil {
		return nil
	} // Start capturing packets

	return tea.Batch(
		packets,
		data.WaitForPacket(m.sub), // Wait for the first packet
	)
}

// Handle incoming messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	case data.PacketData:
		m.packets = append(m.packets, msg)

		// Keep only the last 20 packets
		if len(m.packets) > 20 {
			m.packets = m.packets[len(m.packets)-20:]
		}

		// Wait for the next packet
		return m, data.WaitForPacket(m.sub)
	default:
		return m, nil
	}
}

// Display the packet capture output
func (m model) View() string {
	s := "\nReal-Time Mesh Packet Analyzer\n" + "---------------------------------\n"

	// Show captured packets
	for i, p := range m.packets {
		s += fmt.Sprintf("%2d: %s -> %s | %s | %d bytes\n",
			i+1, p.SrcIP, p.DstIP, p.Protocol, p.Length)
	}

	s += "---------------------------------\n"
	s += "Press any key to exit.\n"

	if m.quitting {
		s += "\n"
	}

	return s
}

func main() {
	// Start the Bubble Tea program
	p := tea.NewProgram(model{
		sub:    make(chan data.PacketData),
		device: "en0",
	})

	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
