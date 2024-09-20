package controllers

import (
	"net/http"

	"github.com/bhtoan2204/go-practice.git/internal/services"
	"github.com/gin-gonic/gin"
)

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

type UserController struct {
	userService services.UserService
}

func (ctrl *UserController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Sign up successful",
		})
	}
}

func (ctrl *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	}
}
