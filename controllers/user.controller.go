package controllers

import (
	"net/http"

	"github.com/bhtoan2204/go-practice.git/dto"
	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.SignUpInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Sign up successful",
			"username": input.Username,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login",
		})
	}
}
