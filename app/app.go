package app

import (
	"context"
	log "github.com/sirupsen/logrus"
	"seem/exception"
	"seem/server"
	"time"
)

func Run() {
	defer exception.CatchError()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Info("Hello World")

	// 启动第一个协程
	go func() {
		defer cancel()
		server.CreatServer()
	}()

	// 启动第二个协程
	go func() {
		defer cancel()
		server.ListenBroadCastMessage()
	}()

	// 等待所有协程完成或上下文取消
	select {
	case <-ctx.Done():
		log.Info("All Goroutines have finished.")
	}

	time.Sleep(time.Second)
}
