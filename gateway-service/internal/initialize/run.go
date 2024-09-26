package initialize

import (
	"gateway-service/global"
	"strconv"
)

func Run() {
	InitViper()
	InitLogger()
	InitMysql()
	InitRedis()
	InitKafka()
	r := InitRouter()
	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
	global.Logger.Info("Initialize all services successfully")
}
