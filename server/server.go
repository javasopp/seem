package server

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net"
	"seem/conf"
	"seem/entity"
	"seem/utils"
	"time"
)

func CreatServer() {
	// 此处返回的是一个ip数组
	localBroadCast, _ := utils.GetIpSetFromNic()
	if len(localBroadCast) <= 0 {
		log.Error("can't get the ip from nic")
		return
	}
	localIp := localBroadCast[0]
	broadcastIp := utils.GetBroadCastIp()
	// 本地ip
	ip := net.ParseIP(localIp)
	localAddr := net.UDPAddr{
		IP:   ip.To4(), // 写局域网下分配IP，0.0.0.0可以用来测试
		Port: conf.Conf.Server.Port,
	}

	// 局域网广播地址
	// 获取这个广播地址之前，先进行获取本机ip地址
	broadcastAddr := net.UDPAddr{
		IP:   net.ParseIP(broadcastIp[localIp]).To4(), // 局域网广播地址
		Port: conf.Conf.Udp.Port,
	}

	conn, err := net.DialUDP("udp", &localAddr, &broadcastAddr)

	if err != nil {
		log.Error(err.Error())
	}

	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	for {
		online := entity.Online{
			Id:   conf.Conf.Server.Id,
			Name: conf.Conf.Server.Name,
			Ip:   conf.Conf.Server.Address,
		}
		marshal, _ := json.Marshal(online)
		_, err = conn.Write(marshal)
		if err != nil {
			log.Error(err.Error())
		}
		time.Sleep(5 * time.Second)
	}

}
