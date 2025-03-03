package main

import (
	"go-auth-api/config"
	"go-auth-api/controllers"
	"go-auth-api/middleware"
	"go-auth-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Auto migrate models
	db.AutoMigrate(&models.User{})

	// Routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected route
	r.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + username})
	})

	r.Run(":8080")
}
