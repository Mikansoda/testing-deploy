package routes

import (
	"github.com/gin-gonic/gin"
	"go-deploy-vps/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Pong!",})
	})

	r.GET("/v1/greeting", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Zahra!",})
	})

	r.GET("/v1/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ini adalah hands-on deployment VPS pada bootcamp dibimbing",})
	})

	// Endpoint baru untuk GET dan POST users
	r.GET("/v1/users", handlers.GetUsers)
	r.POST("/v1/users", handlers.AddNewUser)
}
