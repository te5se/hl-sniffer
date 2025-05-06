package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	tgService *TgService
}

func NewSniffer(tgService *TgService) *Sniffer {
	return &Sniffer{tgService: tgService}
}

func (s *Sniffer) Listen() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}

	device := devices[0].Name

	for _, deviceLocal := range devices {
		if deviceLocal.Description == "Realtek PCIe GbE Family Controller" {
			device = deviceLocal.Name
		}
	}

	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}

	defer handle.Close()

	fmt.Printf("Listening on interface %s...\n", device)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	buffer := []byte{}

	for packet := range packetSource.Packets() {
		ipv4Layer := packet.Layer(layers.LayerTypeIPv4)

		if ipv4Layer == nil {
			continue
		}

		ipv4 := ipv4Layer.(*layers.IPv4)

		if ipv4.SrcIP.String() != "176.9.0.5" {
			continue
		}
		if packet.ApplicationLayer() == nil {
			continue
		}

		payload := string(packet.ApplicationLayer().Payload())

		/* if strings.Contains(payload, "Content-Type: application/json") == false {
			continue
		} */
		/* if strings.Contains(payload, "Content-Length") == false {
			continue
		}
		*/

		if strings.Contains(payload, "HTTP") {
			/* fmt.Println("------------")
			fmt.Println(string(buffer)) */

			split := strings.Split(string(buffer), "\n")

			body := []byte{}

			chunked := strings.Contains(string(buffer), "Transfer-Encoding: chunked")

			for i := len(split) - 1; i != 0; i-- {
				line := split[i]

				line = strings.Trim(line, "\n")
				line = strings.Trim(line, "\r")

				intValue, err := strconv.Atoi(line)
				if err == nil && intValue < 10000 {
					continue
				}

				if chunked && strings.Contains(line, "e68") && len(line) < 10 {
					break
				}
				if !chunked && line == "" {
					break
				}

				body = append([]byte(line), body...)
			}

			s.ProcessRouges(body)

			buffer = []byte{}
		} else {

		}

		buffer = append(buffer, []byte(payload)...)

	}
}
