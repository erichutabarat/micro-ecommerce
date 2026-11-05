package main

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/model"
)

func main() {
	// init db
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("Failed to auto-migrate model: %v", err)
	}
	log.Println("Database connection established and model auto-migrated")

}
