package main

import (
	"fmt"
	"net/http"
	"sort"
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

	clientGone := ginctx.Writer.CloseNotify()
	for {
		ready := true
		select {
		case <-clientGone:
			fmt.Println("Client disconnected, stopping SSE stream")
			return
		default:
			schemasAndTables := db[types.Sha(hash)]
			time.Sleep(time.Duration(1000) * time.Millisecond)
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
	hash := queryParams["--hash"][0]
	schema := queryParams["--schema"][0]
	table := queryParams["--table"][0]

	var filterParams map[string]string = parseFilterParams(queryParams)

	for key, value := range filterParams {
		fmt.Println(key, value)
	}

	var indexStart int
	var indexEnd int
	var err error
	indexStart, indexEnd, err = parseIndexes(queryParams)
	fmt.Println("indexStart", indexStart)
	fmt.Println("indexEnd", indexEnd)
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

func getFilteredPaginatedRows(
	tableRows []types.Row,
	filterParams map[string]string,
	indexStart int,
	indexEnd int,
) map[string]any {
	//lastIndex := min(indexEnd, len(tableRows))
	columnList := make([]string, len(tableRows[0]))
	i := 0
	for key := range tableRows[0] {
		columnList[i] = string(key)
		i++
	}
	sort.Strings(columnList)

	filteredRows := make([]types.Row, 0)
	rowCount := 0
	for _, row := range tableRows {
		match := true
		for key, value := range filterParams {
			if rowValue, ok := row[types.ColumnName(key)]; ok {
				if !strings.Contains(strings.ToLower(string(rowValue)), strings.ToLower(value)) {
					match = false
					break
				}
			} else {
				match = false
				break
			}
		}
		if match {
			rowCount++
			if rowCount >= indexStart && rowCount <= indexEnd {
				filteredRows = append(filteredRows, row)
			}
		}
	}
	//for _, row := range tableData {
	//	if row[types.ColumnName(filterParams["id"])] == types.Value(filterParams["id"]) {
	//		filteredRows = append(filteredRows, row)
	//	}
	//}

	fmt.Println(filterParams)
	var rowListWithColumns map[string]any = make(map[string]any, 0)
	rowListWithColumns["columnList"] = columnList
	rowListWithColumns["rowList"] = filteredRows
	return rowListWithColumns
}
