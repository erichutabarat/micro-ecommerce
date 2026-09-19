package main

import (
    "errors"
	"encoding/json"
	"log"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

type BrokerRequest struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (app *Config) handleBroker(c *gin.Context) {
	var req BrokerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		app.errorResponse(c, 400, "Invalid request body: "+err.Error())
		return
	}

	switch req.Action {
	case "auth.login":
		res, err := handleLogin(req.Payload)
		if err != nil {
			app.errorResponse(c, 401, err.Error())
			return
		}

		app.successResponse(c, 200, res.Message, gin.H{
			"token": res.Token,
		})

	case "product":
		result := gin.H{"id": 1, "name": "Laptop", "price": 999.99}
		app.successResponse(c, 200, "Product fetched successfully", result)

	default:
		log.Println("Targeted service unknown")
		app.errorResponse(c, 404, "Targeted service unknown: "+req.Action)
	}
}

func handleLogin(payload map[string]interface{}) (LoginResponse, error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return LoginResponse{}, err
	}

	var loginReq LoginRequest
	if err := json.Unmarshal(jsonBytes, &loginReq); err != nil {
		return LoginResponse{}, err
	}

	client, err := rpc.Dial("tcp", "auth-service:5001")
	if err != nil {
		log.Printf("Error connecting to auth service: %v", err)
		return LoginResponse{}, err
	}
	defer client.Close()

	var reply LoginResponse
	err = client.Call("AuthService.Login", &loginReq, &reply)
	if err != nil {
		if reply.Message != "" {
			return reply, errors.New(reply.Message)
		}
		return LoginResponse{}, err
	}

	return reply, nil
}