package initialize

import (
	"gateway-service/global"

	routes "gateway-service/internal/router"

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
	// Logging
	r.Use(gin.Logger())
	manageRouter := routes.RouterGroupApp.Manage
	userRouter := routes.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}
	{
		userRouter.InitProductRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitManagerRouter(MainGroup)
	}
	return r
}
