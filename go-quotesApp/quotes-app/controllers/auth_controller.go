package controllers

import (
	"fmt"
	"net/http"
	"quotes-app/quotes-app/internal/database"
	"quotes-app/quotes-app/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt" // used to hash password
)

func CreateUser(c *gin.Context) {

	var authInput models.User
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//  check id user already exists
	db_user := GetUserByEmail(database.Db, authInput.Email)
	if db_user != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Id:          uuid.New(),
		Email:       authInput.Email,
		Password:    string(passwordHash),
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		IsDelete:    false,
		FirstName:   authInput.FirstName,
		LastName:    authInput.LastName,
	}

	// Sanving in Db

	if err := database.Db.Create(&user).Error; err != nil {
		fmt.Printf("error adding user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": user})
}
