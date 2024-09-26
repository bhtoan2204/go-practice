package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (p *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public routes
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/")
		productRouterPublic.GET("/:id")
	}
	// private routes

}
