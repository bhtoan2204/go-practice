package main

import (
	"log"
	"os"

	"github.com/bhtoan2204/go-practice.git/controllers"
	"github.com/bhtoan2204/go-practice.git/repositories"
	"github.com/bhtoan2204/go-practice.git/routes"
	"github.com/bhtoan2204/go-practice.git/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set the default port if not provided
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Initialize the Gin router
	router := gin.New()
	router.Use(gin.Logger())

	// Initialize the repository, service, and controller
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Set up routes with the initialized controller
	routes.AuthRoutes(router, userController)
	routes.UserRoutes(router, userController)

	// Define some example routes
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// Run the server
	router.Run(":" + port)
}
