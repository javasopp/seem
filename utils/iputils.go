package utils

import (
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

// GetIpSetFromNic Get the IP address set of the current network interface card
// 获取网卡的ip地址
func GetIpSetFromNic() ([]string, error) {
	ipAddress, err := getLocalIpV4()
	if err != nil {
		return nil, err
	}
	return ipAddress, nil
}

// 获取本机的ip v4的地址
func getLocalIpV4() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		// 无法获取本地网卡
		log.Error("can't get network interface card!")
		return nil, err
	}
	var ipAddress []string
	for _, inter := range interfaces {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			address, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range address {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					if ipNet.IP.To4() != nil {
						ipAddress = append(ipAddress, ipNet.IP.String())
					}
				}
			}
		}
	}
	if len(ipAddress) > 0 {
		return ipAddress, nil
	} else {
		// 并无任何ip地址
		log.Error("There is no IP address under the network interface card")
		return nil, nil
	}
}

// GetBroadCastIp 获取广播ip
func GetBroadCastIp() map[string]string {
	mapInfo := make(map[string]string)
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
			mapInfo[ipNet.IP.String()] = broadcast.String()
		}
	}
	if len(mapInfo) > 0 {
		return mapInfo
	}
	return nil
}

// 对比获得broadcast ip
func calculateBroadcastAddr(ipNet *net.IPNet) net.IP {
	mask := ipNet.Mask
	ip := ipNet.IP.To4()
	broadcast := make(net.IP, len(ip))
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}
	return broadcast
}

// ReplaceLastOctetWith255 替换最后一位的ip地址
func ReplaceLastOctetWith255(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ""
	}

	ipParts := strings.Split(ipStr, ".")
	if len(ipParts) != 4 {
		return ""
	}

	ipParts[3] = "255"
	newIPStr := strings.Join(ipParts, ".")

	return newIPStr
}
