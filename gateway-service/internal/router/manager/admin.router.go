package manager

import (
	"gateway-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (p *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public routes
	userRouterPublic := Router.Group("/admin")
	{
		userRouterPublic.POST("/login")
	}
	// private routes
	userRouterPrivate := Router.Group("/admin/user")
	userRouterPrivate.Use(middleware.AuthenticationMiddleware())
	{
		userRouterPrivate.PUT("/ban_user")
	}
}
