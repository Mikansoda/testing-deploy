package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-deploy-vps/models"
)

// Slice users sebagai data dummy
var users = []models.User{
	{ID: 1, Name: "Clara"},
	{ID: 2, Name: "Miki"},
	{ID: 3, Name: "Chiki"},
	{ID: 4, Name: "Suneo"},
	{ID: 5, Name: "Jojo"},
}


func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func AddNewUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		// Misalkan jika input tidak valid, return HTTP 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Untuk memberikan ID baru berdasarkan panjang slice users sebagai jumlah user yang ada
	newUser.ID = len(users) + 1
	// Tambahkan user baru ke slice users
	users = append(users, newUser)
	// Return HTTP 201 Created dengan user baru
	c.JSON(http.StatusCreated, newUser)
}