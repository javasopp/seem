package main

import (
	"fmt"
	"log"
	"net"
)

func main1() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() || ipNet.IP.To4() == nil {
				continue
			}

			broadcast := calculateBroadcastAddr(ipNet)
			fmt.Printf("Interface: %s, IP: %s, Broadcast: %s\n", iface.Name, ipNet.IP, broadcast)
		}
	}
}

func calculateBroadcastAddr(ipNet *net.IPNet) net.IP {
	mask := ipNet.Mask
	ip := ipNet.IP.To4()
	broadcast := make(net.IP, len(ip))
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}
	return broadcast
}
