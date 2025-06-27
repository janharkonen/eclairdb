package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Go API")
	router := gin.Default()

	router.Use(corsConfig)
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong\n") })
	router.GET("/data", getData)
	router.Run(":8081")
}

func corsConfig(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func getData(c *gin.Context) {
	c.JSON(http.StatusOK, "data")
}
