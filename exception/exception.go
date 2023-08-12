package exception

import (
	log "github.com/sirupsen/logrus"
)

func CatchError() {
	if err := recover(); err != nil {
		// 进行异常处理
		log.Errorf("there is occurred error,the error info is: %f\n", err)
	}
}
