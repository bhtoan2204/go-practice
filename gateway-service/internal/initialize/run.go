package initialize

import (
	"gateway-service/global"
)

func Run() {
	InitViper()
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8080")
	global.Logger.Info("Initialize all services successfully")
}
