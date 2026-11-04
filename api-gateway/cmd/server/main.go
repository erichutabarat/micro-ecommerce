package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/handler"
)

func main() {
	server := gin.Default()

	dispatchHandler := handler.NewDispatchHandler()

	server.POST("/api/dispatch", dispatchHandler.HandleDispatch)

	log.Println("API GATEWAY RUNNING ON PORT :8080")

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start API GATEWAY on port :8080, err: %v", err)
	}
}
