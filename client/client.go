package client

import (
	log "github.com/sirupsen/logrus"
	"net"
)

// CreateClient Create an udp client side to receive broadcast messages
func CreateClient() {
	_, _ = net.InterfaceAddrs()
	localAddr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0), //写局域网下分配IP，0.0.0.0可以用来测试
		Port: 8080,
	}
	conn, err := net.ListenUDP("udp", &localAddr)
	if err != nil {
		log.Error(err.Error())
	}
	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	buf := make([]byte, 1024)
	for {
		read, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Error(err.Error())
		}
		log.Info(read, remoteAddr)
		log.Info(string(buf))
	}
}
