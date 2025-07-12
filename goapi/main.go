package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janharkonen/eclairdb/postgresloader"
	"github.com/janharkonen/eclairdb/types"
)

var db types.Database

func main() {
	fmt.Println("Starting Go API")
	db = make(types.Database)
	router := gin.Default()

	router.Use(corsConfig)
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong31\n") })
	router.GET("/data", getData)
	router.POST("/postgres", postgres)
	router.Run("0.0.0.0:8081")
}

func corsConfig(ginctx *gin.Context) {
	ginctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ginctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ginctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if ginctx.Request.Method == "OPTIONS" {
		ginctx.AbortWithStatus(204)
		return
	}
	ginctx.Next()
}

func getData(ginctx *gin.Context) {
	ginctx.JSON(http.StatusOK, "data")
}

func postgres(ginctx *gin.Context) {
	var requestBody struct {
		PostgresURL string `json:"postgresurl"`
	}
	if err := ginctx.ShouldBindJSON(&requestBody); err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postgresurl := requestBody.PostgresURL
	err := postgresloader.LoadData(postgresurl, &db)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginctx.JSON(http.StatusOK, db)
}
