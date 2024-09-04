package controllers

import (
	"net/http"

	"github.com/bhtoan2204/go-practice.git/dto"
	"github.com/bhtoan2204/go-practice.git/services"
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
		var input dto.SignUpInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := ctrl.userService.SignUp(input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Sign up successful",
		})
	}
}

func (ctrl *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user, err := ctrl.userService.Login(input)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"username": user.Username,
		})
	}
}
