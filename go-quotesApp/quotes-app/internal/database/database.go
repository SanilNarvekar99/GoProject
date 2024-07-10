package database

import (
	"database/sql"
	"fmt"
	"log"

	"quotes-app/quotes-app/internal/config"

	_ "github.com/lib/pq" // important to import postgres driver
)

func ConnectDB(config *config.Config) {
	// log.Println("Connecting to database...")
	// log.Println("Host:", config.DBHost, "Port:", config.DBPort, "Name:", config.DBName, "User:", config.DBUser)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//Ensure the connection is available by pinging the database
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to database successfully")
	defer db.Close()
}
