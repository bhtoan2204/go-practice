package initialize

import (
	"github.com/bhtoan2204/go-practice.git/global"
	"github.com/bhtoan2204/go-practice.git/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
}
