package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	router.GET("/filtered_paginated_products", getFilteredPaginatedProducts)
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
	hash := ginctx.Query("hash")
	if hash == "" {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing hash parameter"})
		return
	}
	dbmap, ok := db[types.Sha(hash)]
	if !ok {
		ginctx.JSON(http.StatusNotFound, gin.H{"error": "Hash not found"})
		return
	}
	ginctx.JSON(http.StatusOK, dbmap)
}

func getFilteredPaginatedProducts(ginctx *gin.Context) {
	var queryParams map[string][]string = ginctx.Request.URL.Query()

	var filterParams map[string]string = parseFilterParams(queryParams)
	hash := filterParams["hash"]
	schema := filterParams["schema"]
	table := filterParams["table"]

	var indexStart int
	var indexEnd int
	var err error
	indexStart, indexEnd, err = parseIndexes(queryParams)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(hash)
	fmt.Println(schema)
	fmt.Println(table)

	if db[types.Sha(hash)][types.Schema(schema)][types.Table(table)] == nil {
		fmt.Println("Table not found")
		return
	}

	var filteredRows = getFilteredPaginatedRows(filterParams, indexStart, indexEnd)

	fmt.Println(filterParams)
	fmt.Println(indexStart)
	fmt.Println(indexEnd)
	ginctx.JSON(http.StatusOK, filteredRows)
}

// Helper function
func parseFilterParams(queryParams map[string][]string) map[string]string {
	var filterParams map[string]string = make(map[string]string, len(queryParams))
	for key, value := range queryParams {
		if key[0:2] == "--" {
			continue
		}
		filterParams[key] = value[0]
	}
	return filterParams
}

// Helper function
func parseIndexes(queryParams map[string][]string) (int, int, error) {
	var indexStart int
	var indexEnd int
	var err error
	if indexesString, ok := queryParams["--indexes"]; ok {
		parts := strings.Split(indexesString[0], "-")
		indexStart, err = strconv.Atoi(parts[0])
		indexEnd, err = strconv.Atoi(parts[1])
	}
	return indexStart, indexEnd, err
}

func getFilteredPaginatedRows(filterParams map[string]string, indexStart int, indexEnd int) []map[string]string {
	return []map[string]string{
		{"id": "1", "name": "Product 1"},
		{"id": "2", "name": "Product 2"},
		{"id": "3", "name": "Product 3"},
	}
}
