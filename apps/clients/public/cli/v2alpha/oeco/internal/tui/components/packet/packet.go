// package packet
package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Message type for Bubble Tea
type packetMsg struct {
	SrcIP    string
	DstIP    string
	Protocol string
	Length   int
}

// Model for Bubble Tea
type model struct {
	packets []packetMsg // Store packet information
}

func (m model) Init() tea.Cmd {
	// Start packet capture
	return capturePacketsCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case packetMsg:
		// Add the new packet to the list
		m.packets = append(m.packets, msg)

		// Keep only the last 20 packets
		if len(m.packets) > 20 {
			m.packets = m.packets[len(m.packets)-20:]
		}

		return m, nil

	case tea.KeyMsg:
		// Quit the program when "q" is pressed
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	// Header
	b.WriteString("Real-Time Packet Sniffer\n")
	b.WriteString(strings.Repeat("-", 50) + "\n")

	// Display captured packets
	for i, p := range m.packets {
		b.WriteString(fmt.Sprintf(
			"%2d: %s -> %s | %s | %d bytes\n",
			i+1, p.SrcIP, p.DstIP, p.Protocol, p.Length,
		))
	}

	// Footer
	b.WriteString(strings.Repeat("-", 50) + "\n")
	b.WriteString("Press 'q' to quit.\n")

	return b.String()
}

// Command to capture packets
func capturePacketsCmd() tea.Cmd {
	return func() tea.Msg {
		// Open a live capture
		handle, err := pcap.OpenLive("eth0", 1024, true, pcap.BlockForever)
		if err != nil {
			log.Fatalf("Error opening device: %v", err)
		}
		defer handle.Close()

		// Use pcapgo to create a packet source
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for packet := range packetSource.Packets() {
			// Extract packet metadata
			networkLayer := packet.NetworkLayer()
			if networkLayer == nil {
				continue
			}

			src, dst := networkLayer.NetworkFlow().Endpoints()
			length := len(packet.Data())
			protocol := packet.TransportLayer().LayerType().String()

			// Return the packet information as a message
			return packetMsg{
				SrcIP:    src.String(),
				DstIP:    dst.String(),
				Protocol: protocol,
				Length:   length,
			}
		}

		return nil
	}
}

func main() {
	// Check for available network interfaces
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Failed to list interfaces: %v", err)
	}

	if len(interfaces) == 0 {
		log.Fatal("No network interfaces found.")
	}

	// Print available interfaces
	fmt.Println("Available network interfaces:")
	for _, iface := range interfaces {
		fmt.Printf("- %s (%s)\n", iface.Name, iface.Description)
	}

	// Prompt the user to enter the interface name
	var device string
	fmt.Print("Enter the interface name to sniff: ")
	fmt.Scanln(&device)

	// Start the Bubble Tea program
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		log.Fatalf("Error starting program: %v", err)
	}
}
