package server

import (
	log "github.com/sirupsen/logrus"
	"net"
)

// ListenBroadCastMessage 监听局域网广播消息
func ListenBroadCastMessage() {
	addr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		log.Error("Error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Error("Error:", err)
		return
	}

	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	buffer := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Error("Error:", err)
			continue
		}

		message := string(buffer[:n])
		log.Infof("Received broadcast message: %s\n", message)
	}
}
