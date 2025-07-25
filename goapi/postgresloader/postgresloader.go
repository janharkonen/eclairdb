package postgresloader

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"sync"

	_ "github.com/lib/pq"

	"github.com/janharkonen/eclairdb/types"
)

func LoadData(postgresurl string, db *types.Database) (types.Sha, error) {

	hasher := sha256.New()
	hasher.Write([]byte(postgresurl))
	postgresurlsha := types.Sha(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))

	dbclient, err := sql.Open("postgres", postgresurl)
	if err != nil {
		return "", err
	}
	defer dbclient.Close()
	mu := sync.Mutex{}
	mu.Lock()
	(*db)[postgresurlsha] = make(types.Schema)
	mu.Unlock()

	if err := dbclient.Ping(); err != nil {
		return "", errors.New("postgresloader: " + err.Error())
	}
	query := `
		SELECT 
			s.schema_name, 
			t.table_name 
		FROM 
			information_schema.schemata s
		LEFT JOIN 
			information_schema.tables t 
		ON 
			s.schema_name = t.table_schema
		WHERE 
			s.schema_name NOT LIKE 'pg_%' 
			AND s.schema_name != 'information_schema'
		ORDER BY 
			s.schema_name, t.table_name
	`

	schemaTablesRows, err := dbclient.Query(query)
	if err != nil {
		return "", err
	}
	defer schemaTablesRows.Close()

	// Process the results and build the data structure
	wg := sync.WaitGroup{}
	for schemaTablesRows.Next() {
		var schemaName, tableName sql.NullString
		if err := schemaTablesRows.Scan(&schemaName, &tableName); err != nil {
			return "", err
		}

		if schemaName.Valid {
			if tableName.Valid {
				mu.Lock()
				if _, exists := (*db)[postgresurlsha][types.SchemaName(schemaName.String)]; !exists {
					(*db)[postgresurlsha][types.SchemaName(schemaName.String)] = make(types.Table)
				}
				(*db)[postgresurlsha][types.SchemaName(schemaName.String)][types.TableName(tableName.String)] = make([]types.Row, 0)
				mu.Unlock()
				wg.Add(1)
				go addTableToDb(&wg, postgresurlsha, types.SchemaName(schemaName.String), types.TableName(tableName.String), db, dbclient)
			}
		}
	}
	wg.Wait()
	if err := schemaTablesRows.Err(); err != nil {
		return "", err
	}
	return postgresurlsha, nil
}

func addTableToDb(wg *sync.WaitGroup, postgresurlsha types.Sha, schema types.SchemaName, table types.TableName, db *types.Database, dbclient *sql.DB) {
	defer wg.Done()
	query := fmt.Sprintf("SELECT * FROM %s.%s", string(schema), string(table))
	fmt.Println(query)
	dbRows, err := dbclient.Query(query)
	if err != nil {
		return
	}
	defer dbRows.Close()

	var rows []types.Row
	var columnNames []string
	columnNames, err = dbRows.Columns()
	if err != nil {
		return
	}

	for dbRows.Next() {
		scanArgs := make([]interface{}, len(columnNames))
		values := make([]sql.NullString, len(columnNames))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		if err := dbRows.Scan(scanArgs...); err != nil {
			return
		}
		row := make(types.Row)
		for i, columnName := range columnNames {
			if values[i].Valid {
				// If the value is not NULL, use the string value
				valueStr := values[i].String
				row[types.ColumnName(columnName)] = types.Value(valueStr)
			} else {
				// If the value is NULL, use an empty string
				row[types.ColumnName(columnName)] = types.Value("")
			}
		}

		rows = append(rows, row)
	}

	(*db)[postgresurlsha][schema][table] = rows
}

func Addition(a int, b int) int {
	return a + b
}
