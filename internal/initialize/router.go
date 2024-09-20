package initialize

import (
	"github.com/bhtoan2204/go-practice.git/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Initialize the Gin router
	var r *gin.Engine

	if global.Config.Server.Mode == "local" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	r.Use(gin.Logger())
	return r
}
