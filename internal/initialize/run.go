package initialize

import (
	"github.com/bhtoan2204/go-practice.git/global"
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
