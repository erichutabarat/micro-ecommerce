package main

import (
    "github.com/gin-gonic/gin"
)

type Config struct {

}

func main() {
    app := Config{}

    r := gin.Default()

    r.POST("/handle", app.handleBroker)

    r.Run(":8080")
}