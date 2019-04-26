package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	println("Net-up!")
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Printf("Error get devises: %s", err.Error())
		os.Exit(1)
	}
	device := devices[0].Name
	handle, err := pcap.OpenLive(device, 1024, false, 30*time.Second)
	if err != nil {
		fmt.Printf("Error capturing packets: %s", err.Error())
	}
	fmt.Printf("capturing packets for: %s \n", device)
	filter := "tcp and port 3000"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		fmt.Printf("Error applying filter: %s", err.Error())
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

}
