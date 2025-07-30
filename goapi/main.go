package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/janharkonen/eclairdb/postgresloader"
	"github.com/janharkonen/eclairdb/types"
)

var db types.Databases

func main() {
	fmt.Println("Starting Go API")
	db = make(types.Databases)
	router := gin.Default()

	router.Use(corsConfig)
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong33\n") })
	router.GET("/data", getData)
	router.POST("/connect-postgres", connectPostgres)
	router.GET("/get-schemas-and-tables", getSchemasAndTables)
	router.GET("/get-schemas-and-tables-stream", getSchemasAndTablesStream)
	// TODO: change name of method
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
	schemaTableMap := make(map[string]map[string]bool)

	if dbData, ok := db[types.Sha(hash)]; ok {
		for schemaName, tables := range dbData {
			schemaTableMap[string(schemaName)] = make(map[string]bool)
			for tableName := range tables {
				tableIsDone := db[types.Sha(hash)][types.SchemaName(schemaName)][types.TableName(tableName)].Done
				schemaTableMap[string(schemaName)][string(tableName)] = tableIsDone
			}
		}
		ginctx.JSON(http.StatusOK, schemaTableMap)
	} else {
		ginctx.JSON(http.StatusNotFound, gin.H{"error": "Hash not found"})
		return
	}

}

func getSchemasAndTablesStream(ginctx *gin.Context) {

	ginctx.Header("Content-Type", "text/event-stream")
	ginctx.Header("Cache-Control", "no-cache")
	ginctx.Header("Connection", "keep-alive")
	//ginctx.Header("Transfer-Encoding", "chunked")
	//ginctx.Header("X-Accel-Buffering", "no")
	//ginctx.Header("X-Frame-Options", "SAMEORIGIN")
	//ginctx.Header("X-XSS-Protection", "1; mode=block")
	hash := ginctx.Query("hash")
	if hash == "" {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing hash parameter"})
		return
	}

	var seconds int = 1
	clientGone := ginctx.Writer.CloseNotify()
	for {
		ready := true
		select {
		case <-clientGone:
			fmt.Println("Client disconnected, stopping SSE stream")
			return
		default:
			schemasAndTables := db[types.Sha(hash)]
			time.Sleep(time.Duration(seconds) * time.Second)
			for schemaname, schema := range schemasAndTables {
				for tablename := range schema {
					if schema[tablename].Done {
						ginctx.SSEvent("table_ready", fmt.Sprintf("%s:%s", schemaname, tablename))
						ginctx.Writer.Flush()
					} else {
						ready = false
					}
				}
			}
		}
		if ready {
			break
		}
	}
	ginctx.SSEvent("complete", "complete")
	ginctx.Writer.Flush()
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

	tableData := db[types.Sha(hash)][types.SchemaName(schema)][types.TableName(table)]
	if !tableData.Done {
		fmt.Println("Table not found")
		return
	}

	var filteredRows = getFilteredPaginatedRows(tableData.Rows, filterParams, indexStart, indexEnd)

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

func getFilteredPaginatedRows(tableData []types.Row, filterParams map[string]string, indexStart int, indexEnd int) []types.Row {
	filteredRows := make([]types.Row, 0)
	lastIndex := min(indexEnd, len(tableData))
	for i := indexStart; i < lastIndex; i++ {
		filteredRows = append(filteredRows, tableData[i])
	}
	//for _, row := range tableData {
	//	if row[types.ColumnName(filterParams["id"])] == types.Value(filterParams["id"]) {
	//		filteredRows = append(filteredRows, row)
	//	}
	//}

	fmt.Println(filterParams)

	return filteredRows
}
