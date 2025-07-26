package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Struct user sebagai representasi data pengguna
type User struct {
	ID 	 int    `json:"id"`
	Name string `json:"name"`
}

// Slice users sebagai data dummy
var users = []User{
	{ID: 1, Name: "Clara"},
	{ID: 2, Name: "Miki"},
	{ID: 3, Name: "Chiki"},
	{ID: 4, Name: "Suneo"},
	{ID: 5, Name: "Jojo"},
}

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	r.GET("/v1/greeting", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Zahra!",
		})
	})

	r.GET("/v1/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ini adalah hands-on deployment VPS pada bootcamp dibimbing",
		})
	})

	// Endpoint baru untuk GET dan POST users
	r.GET("/v1/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.POST("/v1/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			// Misalkan jika input tidak valid, return HTTP 400 Bad Request
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		// Untuk memberikan ID baru berdasarkan panjang slice users sebagai jumlah user yang ada
		newUser.ID = len(users) + 1
		// Tambahkan user baru ke slice users
		users = append(users, newUser)
		// Return HTTP 201 Created dengan user baru
		c.JSON(201, newUser)
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Server is running on port " + port)
	r.Run(":" + port)
}
