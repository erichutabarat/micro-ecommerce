package main

import (
    "log"
    "github.com/gin-gonic/gin"
)

type BrokerRequest struct {
    Action  string                 `json:"action"`
    Payload map[string]interface{} `json:"payload"`
}

func (app *Config) handleBroker(c *gin.Context) {
    var req BrokerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    switch req.Action {
    case "auth":
        log.Println("Calling auth service...")
        c.JSON(200, gin.H{"status": "success", "service": "auth", "received": req.Payload})

    case "product":
        log.Println("Calling product service...")
        c.JSON(200, gin.H{"status": "success", "service": "product", "received": req.Payload})

    default:
        log.Println("Targeted service unknown")
        c.JSON(400, gin.H{"error": "Targeted service unknown"})
    }
}