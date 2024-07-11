package controllers

import (
	"log"
	"net/http"
	"quotes-app/quotes-app/internal/database"
	"quotes-app/quotes-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.Db.Where("id =?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserByEmail(db *gorm.DB, email string) *models.User {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func UpdateUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("User not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := database.Db.Save(&user).Error; err != nil {
		log.Println("Error saving user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user *models.User

	if err := database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("User not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := database.Db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully"})
}
