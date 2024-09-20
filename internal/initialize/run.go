package initialize

import "github.com/bhtoan2204/go-practice.git/global"

func Run() {
	InitViper()
	InitLogger()
	InitMysql()
	InitRedis()
	global.Logger.Info("Initialize all services successfully")
	r := InitRouter()
	r.Run(":8080")
}
