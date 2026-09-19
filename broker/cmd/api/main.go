package main

import (
    "time"

	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

type Config struct {

}

func main() {
    app := Config{}

    r := gin.Default()

    r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Your Vite frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

    r.POST("/handle", app.handleBroker)

    r.Run(":8080")
}