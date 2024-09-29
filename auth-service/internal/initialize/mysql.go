package initialize

import (
	"auth-service/global"
	"auth-service/internal/po"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func migrateTables() {
	err := global.MDB.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(err)
	}
}

func SetPool() {
	sqlDb, err := global.MDB.DB()
	if err != nil {
		global.Logger.Error("Failed to set pool", zap.Error(err))
		panic(err)
	}
	sqlDb.SetMaxIdleConns(global.Config.MySQLConfig.MaxIdleConns)
	sqlDb.SetMaxOpenConns(global.Config.MySQLConfig.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Config.MySQLConfig.MaxLifetime))
}

func InitMysql() {
	// connect to MySQL
	m := global.Config.MySQLConfig
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Pass, m.Host, m.Port, m.Name)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to connect to MySQL")
	global.Logger.Info("Connected to MySQL successfully")
	global.MDB = db

	// set Pool
	SetPool()
	// migrate tables
	migrateTables()
}
