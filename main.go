package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"seem/exception"
	"seem/server"
)

func main() {
	defer exception.CatchError()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Info("Hello World")
	go func() {
		server.CreatServer()
		cancel()
	}()

	go func() {
		server.ListenBroadCastMessage()
		cancel()
	}()

	// 等待所有协程完成或上下文取消
	select {
	case <-ctx.Done():
		log.Info("All Goroutines have finished.")
	}
}
