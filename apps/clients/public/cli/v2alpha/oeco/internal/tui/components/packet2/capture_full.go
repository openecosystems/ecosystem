package main

//
//import (
//	"fmt"
//	"log"
//
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
//	"github.com/google/gopacket/pcap"
//)
//
//func main() {
//	handle, err := pcap.OpenLive("en0", 1600, true, pcap.BlockForever) // Change interface
//	if err != nil {
//		log.Fatal("Error opening device:", err)
//	}
//	defer handle.Close()
//
//	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//
//	for packet := range packetSource.Packets() {
//		processPacket(packet)
//	}
//}
//
//func processPacket(packet gopacket.Packet) {
//	if netLayer := packet.NetworkLayer(); netLayer != nil {
//		fmt.Println("Src IP:", netLayer.NetworkFlow().Src())
//		fmt.Println("Dst IP:", netLayer.NetworkFlow().Dst())
//	}
//
//	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
//		tcp, _ := tcpLayer.(*layers.TCP)
//		fmt.Println("TCP Packet - SrcPort:", tcp.SrcPort, "DstPort:", tcp.DstPort)
//
//		// Filter only HTTP traffic (Port 80)
//		if tcp.DstPort == 80 || tcp.SrcPort == 80 {
//			if len(tcp.Payload) > 0 {
//				fmt.Println("HTTP Payload:", string(tcp.Payload))
//			} else {
//				fmt.Println("Empty TCP Payload")
//			}
//		}
//	}
//}
