package initialize

import (
	"os"
	"simple_bank/global"
	"strconv"

	"go.uber.org/zap"
)

func Run() {
	InitViper()
	InitLogger()
	InitMysql()
	InitRedis()
	InitKafka()
	r := InitRouter()
	global.Logger.Info("Initialize all services successfully")
	if err := r.Run(":" + strconv.Itoa(global.Config.Server.Port)); err != nil {
		global.Logger.Error("Failed to start server", zap.Error(err))
		// Handle error
		os.Exit(1)
	}
}
