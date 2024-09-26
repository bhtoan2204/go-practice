package initialize

import (
	"gateway-service/global"
	"gateway-service/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LogConfig)
}
