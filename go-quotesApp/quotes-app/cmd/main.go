package main

import (
	"quotes-app/quotes-app/internal/config"
	"quotes-app/quotes-app/internal/database"
)

func main() {
	config_credential := config.LoadConfig()
	database.ConnectDB(config_credential)
}
