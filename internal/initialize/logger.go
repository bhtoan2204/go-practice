package initialize

import (
	"simple_bank/global"
	"simple_bank/package/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
}
