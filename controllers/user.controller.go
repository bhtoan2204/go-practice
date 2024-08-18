package controllers

import "github.com/gin-gonic/gin"

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Sign up",
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
