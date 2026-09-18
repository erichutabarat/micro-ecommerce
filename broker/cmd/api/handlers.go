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
        app.errorResponse(c, 400, "Invalid request body: "+err.Error())
        return
    }

    switch req.Action {
    case "auth":
        result := gin.H{"token": "mock-jwt-token-123"}
        app.successResponse(c, 200, "Auth successful", result)

    case "product":
        result := gin.H{"id": 1, "name": "Laptop", "price": 999.99}
        app.successResponse(c, 200, "Product fetched successfully", result)

    default:
        log.Println("Targeted service unknown")
        app.errorResponse(c, 404, "Targeted service unknown: "+req.Action)
    }
}