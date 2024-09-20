package user

import (
	"github.com/bhtoan2204/go-practice.git/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routes
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/signup")
		userRouterPublic.POST("/sms")
		userRouterPublic.POST("/login")
	}
	// private routes
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middleware.AuthenticationMiddleware())
	{
		userRouterPrivate.GET("/")
		userRouterPrivate.GET("/:id")
	}
}
