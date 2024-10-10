package initialize

import (
	"fmt"
	"simple_bank/global"
	"simple_bank/internal/po"
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
	models := []interface{}{
		&po.Account{},
		&po.AuditLog{},
		&po.Event{},
		&po.Notification{},
		&po.Permission{},
		&po.RefreshToken{},
		&po.Role{},
		&po.SystemConfiguration{},
		&po.Transaction{},
		&po.User{},
	}
	err := global.MDB.AutoMigrate(models...)
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
	m := global.Config.MySQLConfig

	// Connect to MySQL without specifying database
	dsnNoDB := "%s:%s@tcp(%s:%v)/?charset=utf8mb4&parseTime=True&loc=Local"
	var sNoDB = fmt.Sprintf(dsnNoDB, m.User, m.Pass, m.Host, m.Port)
	dbNoDB, err := gorm.Open(mysql.Open(sNoDB), &gorm.Config{})
	checkErrorPanic(err, "Failed to connect to MySQL without specifying database")

	// Create database if not exists
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", m.Name)
	err = dbNoDB.Exec(createDBSQL).Error
	checkErrorPanic(err, "Failed to create database")

	// Close connection to MySQL without specifying database
	sqlDBNoDB, err := dbNoDB.DB()
	checkErrorPanic(err, "Failed to get db connection")
	sqlDBNoDB.Close()

	// Connect to MySQL with database
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Pass, m.Host, m.Port, m.Name)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to connect to MySQL with database")
	global.Logger.Info("Connected to MySQL successfully")
	global.MDB = db

	SetPool()

	// migrateTables()
}
