package routes

import (
	"github.com/bhtoan2204/go-practice.git/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("auth/login", controllers.Login())
}
