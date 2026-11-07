package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/auth"
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
)

func main() {
	server := gin.Default()

	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 5 {
		log.Fatal("JWT_SECRET must be set and at least 6 characters long")
	}

	verifier := auth.NewVerifier(secret)
	dispatchHandler := handler.NewDispatchHandler()

	api := server.Group("/api", middleware.OptionalAuth(verifier))
	{
		api.POST("/dispatch", dispatchHandler.HandleDispatch)
	}

	log.Println("API GATEWAY RUNNING ON PORT :8080")

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start API GATEWAY on port :8080, err: %v", err)
	}
}
