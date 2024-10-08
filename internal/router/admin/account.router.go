package admin

import (
	"simple_bank/internal/controller"

	"github.com/gin-gonic/gin"
)

type AccountAccountRouter struct{}

func (p *AccountAccountRouter) InitAdminAccountRouter(router *gin.RouterGroup, accountController controller.AccountController) {
	accountAccountRouter := router.Group("account")
	{
		accountAccountRouter.GET("/:id", accountController.GetAccountByID)
		accountAccountRouter.POST("", accountController.CreateAccount)
		accountAccountRouter.PUT("/:id", accountController.UpdateAccount)
		accountAccountRouter.DELETE("/:id", accountController.DeleteAccount)
		accountAccountRouter.GET("/user/:userID", accountController.ListAccountsByUserID)
	}
}
