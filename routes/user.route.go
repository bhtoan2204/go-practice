package routes

import (
	"github.com/bhtoan2204/go-practice.git/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
}
