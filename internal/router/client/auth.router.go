package client

import "github.com/gin-gonic/gin"

type ClientAuthRouter struct {
}

func (p *ClientAuthRouter) InitClientAuthRouter(Router *gin.RouterGroup) {
	clientAuthRouter := Router.Group("auth")
	{
		clientAuthRouter.POST("login")
		clientAuthRouter.POST("refresh")
		clientAuthRouter.POST("logout")
	}
}
