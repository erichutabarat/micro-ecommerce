package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func (app *Config) successResponse(c *gin.Context, code int, message string, data interface{}) {
    c.JSON(code, Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

func (app *Config) errorResponse(c *gin.Context, code int, errMessage string) {
    c.JSON(code, Response{
        Success: false,
        Error:   errMessage,
    })
}