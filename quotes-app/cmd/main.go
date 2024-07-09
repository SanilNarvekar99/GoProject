package main

import (
	"fmt"
	"quotes-app/quotes-app/internal/config"
	"quotes-app/quotes-app/internal/database"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a Go program")

	config_credential := config.LoadConfig()
	database.ConnectDB(config_credential)
}
