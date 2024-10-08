package initialize

import (
	"simple_bank/global"
	"simple_bank/internal/controller"
	"simple_bank/internal/middleware"
	"simple_bank/internal/router"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func InitRouter() *gin.Engine {
	// Initialize the Gin router
	var r *gin.Engine
	requestsPerSecond := rate.Limit(5)
	burstSize := 10

	rl := middleware.NewRateLimiter(requestsPerSecond, burstSize)

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
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimitMiddleware(rl))

	clientRouter := router.RouterGroupApp.Client
	adminRouter := router.RouterGroupApp.Admin
	// publicRouter := router.RouterGroupApp.Public
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}
	{
		clientGroup := MainGroup.Group("/client")
		clientRouter.InitClientAuthRouter(clientGroup)
		clientRouter.InitClientUserRouter(clientGroup)
	}
	{
		adminGroup := MainGroup.Group("/admin")
		adminRouter.InitAdminAccountRouter(adminGroup, controller.NewAccountController())
	}
	return r
}
