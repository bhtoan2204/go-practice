package client

import "github.com/gin-gonic/gin"

type ClientUserRouter struct {
}

func (p *ClientUserRouter) InitClientUserRouter(Router *gin.RouterGroup) {
	clientUserRouter := Router.Group("user")
	{
		clientUserRouter.POST("register")
		clientUserRouter.GET("")
		clientUserRouter.PUT(":id")
	}
}
