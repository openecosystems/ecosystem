package main

//
//import (
//	"crypto/x509"
//	"encoding/pem"
//	"fmt"
//	"log"
//	"strings"
//	"time"
//
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
//	"github.com/google/gopacket/pcap"
//)
//
//var connectionTimes = make(map[string]time.Time)
//
//func main() {
//	// Open the device for capturing
//	handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever)
//	if err != nil {
//		log.Fatal("Error opening device:", err)
//	}
//	defer handle.Close()
//
//	// Create a packet source to read packets
//	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//
//	for packet := range packetSource.Packets() {
//		processPacket(packet)
//	}
//}
//
//func processPacket(packet gopacket.Packet) {
//	// Extract metadata
//	if netLayer := packet.NetworkLayer(); netLayer != nil {
//		fmt.Println("Src IP:", netLayer.NetworkFlow().Src())
//		fmt.Println("Dst IP:", netLayer.NetworkFlow().Dst())
//	}
//
//	if transportLayer := packet.TransportLayer(); transportLayer != nil {
//		fmt.Println("Protocol:", transportLayer.LayerType())
//	}
//
//	// Capture TCP Handshake Time (RTT Calculation)
//	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
//		tcp, _ := tcpLayer.(*layers.TCP)
//		flow := packet.NetworkLayer().NetworkFlow().String()
//		if tcp.SYN && !tcp.ACK {
//			connectionTimes[flow] = packet.Metadata().Timestamp
//		} else if tcp.SYN && tcp.ACK {
//			if startTime, exists := connectionTimes[flow]; exists {
//				rtt := packet.Metadata().Timestamp.Sub(startTime)
//				fmt.Println("RTT:", rtt)
//			}
//		}
//	}
//
//	// Extract HTTP Requests/Responses safely
//	if appLayer := packet.ApplicationLayer(); appLayer != nil {
//		payload := string(appLayer.Payload())
//		if len(payload) > 4 {
//			if strings.HasPrefix(payload, "GET") || strings.HasPrefix(payload, "POST") {
//				fmt.Println("HTTP Request:", extractFirstLine(payload))
//			} else if strings.HasPrefix(payload, "HTTP/1.1") {
//				fmt.Println("HTTP Response:", extractFirstLine(payload))
//			}
//		}
//	}
//
//	// DNS Query Capture
//	if dnsLayer := packet.Layer(layers.LayerTypeDNS); dnsLayer != nil {
//		dns, _ := dnsLayer.(*layers.DNS)
//		if !dns.QR && len(dns.Questions) > 0 {
//			fmt.Println("DNS Query:", string(dns.Questions[0].Name))
//		}
//	}
//
//	// TLS Certificate Inspection
//	if tlsLayer := packet.Layer(layers.LayerTypeTLS); tlsLayer != nil {
//		parseTLSCertificate(tlsLayer.LayerContents())
//	}
//
//	// Estimate network emissions
//	emissions := estimateEmissions(packet.Metadata().Length)
//	fmt.Printf("Packet Emissions: %.6f gCO2\n", emissions)
//}
//
//// Extracts the first line of an HTTP request/response
//func extractFirstLine(payload string) string {
//	lines := strings.Split(payload, "\n")
//	if len(lines) > 0 {
//		return strings.TrimSpace(lines[0])
//	}
//	return ""
//}
//
//// Parses TLS certificates to check validity
//func parseTLSCertificate(certData []byte) {
//	block, _ := pem.Decode(certData)
//	if block == nil {
//		fmt.Println("No valid TLS certificate found")
//		return
//	}
//
//	cert, err := x509.ParseCertificate(block.Bytes)
//	if err != nil {
//		fmt.Println("Error parsing certificate:", err)
//		return
//	}
//
//	fmt.Println("TLS Certificate Issuer:", cert.Issuer)
//	fmt.Println("TLS Certificate Valid Until:", cert.NotAfter)
//}
//
//// Estimates carbon emissions based on packet size
//func estimateEmissions(packetSize int) float64 {
//	energyPerByte := 0.0000001 // Example value in kWh
//	carbonIntensity := 0.4     // kgCO2 per kWh
//	return float64(packetSize) * energyPerByte * carbonIntensity
//}
