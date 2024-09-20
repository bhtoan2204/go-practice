package manager

import (
	"github.com/bhtoan2204/go-practice.git/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (p *UserRouter) InitManagerRouter(Router *gin.RouterGroup) {
	// private routes
	userRouterPrivate := Router.Group("/admin/user")
	userRouterPrivate.Use(middleware.AuthenticationMiddleware())
	{
		userRouterPrivate.PUT("/something")
	}
}
