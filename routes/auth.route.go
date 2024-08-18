package routes

import (
	"github.com/bhtoan2204/go-practice.git/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
}
