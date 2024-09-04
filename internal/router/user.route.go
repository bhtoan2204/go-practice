package routes

import (
	"github.com/bhtoan2204/go-practice.git/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine, userController *controllers.UserController) {
	incomingRoutes.POST("users/signup", userController.SignUp())
}
