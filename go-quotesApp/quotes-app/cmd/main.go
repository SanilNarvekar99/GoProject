package main

import (
	"net/http"
	"quotes-app/quotes-app/controllers"
	"quotes-app/quotes-app/internal/config"
	"quotes-app/quotes-app/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config_credential := config.LoadConfig()
	database.ConnectDB(config_credential)

	r := gin.Default()                //function initialises a new Gin router instance by calling gin.Default(), which sets up a default middleware stack for handling HTTP requests
	r.GET("/", func(c *gin.Context) { // a route is defined using r.GET("/"), specifying that the handler function should be executed when a GET request is made to the root URL ("/").
		c.JSON(http.StatusOK, gin.H{"data": "Hello, world"})
	})

	r.POST("/auth/signup", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserById)
	r.PATCH("/users/:id", controllers.UpdateUserById)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run(":3000")
}
