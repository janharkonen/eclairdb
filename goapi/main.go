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
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong33\n") })
	router.GET("/data", getData)
	router.POST("/connect-postgres", connectPostgres)
	router.GET("/get-schemas-and-tables", getSchemasAndTables)
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

func connectPostgres(ginctx *gin.Context) {
	var requestBody struct {
		URI string `json:"uri"`
	}
	if err := ginctx.ShouldBindJSON(&requestBody); err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postgresurl := requestBody.URI
	postgresurlsha, err := postgresloader.LoadData(postgresurl, &db)
	if err != nil {
		ginctx.JSON(http.StatusInternalServerError, gin.H{"error": "GIN ERROR: " + err.Error()})
		return
	}

	ginctx.JSON(http.StatusOK, postgresurlsha)
}

func getSchemasAndTables(ginctx *gin.Context) {
	fmt.Println("part 1")
	hash := ginctx.Query("hash")
	fmt.Println("part 2")
	if hash == "" {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing hash parameter"})
		return
	}
	fmt.Println("part 3")
	dbmap, ok := db[types.Sha(hash)]
	fmt.Println("part 4")
	if !ok {
		ginctx.JSON(http.StatusNotFound, gin.H{"error": "Hash not found"})
		return
	}
	fmt.Println("part 5")
	ginctx.JSON(http.StatusOK, dbmap)
}
