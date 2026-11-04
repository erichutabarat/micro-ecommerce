package server

import (
	"auth-service/internal/config"
	"auth-service/internal/model"
	"log"
)

func main() {
	// Connect Database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// auto-migrate model
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate model: %v", err)
	}

	log.Println("Database connection established and model auto-migrated")
}
