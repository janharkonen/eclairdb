package postgresloader

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"sync"

	_ "github.com/lib/pq"

	"github.com/janharkonen/eclairdb/types"
)

func LoadData(postgresurl string, db *types.Database) error {

	hasher := sha256.New()
	hasher.Write([]byte(postgresurl))
	postgresurlsha := types.Sha(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))

	dbclient, err := sql.Open("postgres", postgresurl)
	if err != nil {
		return err
	}
	defer dbclient.Close()
	mu := sync.Mutex{}
	mu.Lock()
	(*db)[postgresurlsha] = make(map[types.Schema]map[types.Table]map[types.Key]types.Value)
	mu.Unlock()

	if err := dbclient.Ping(); err != nil {
		return err
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
		return err
	}
	defer schemaTablesRows.Close()

	// Process the results and build the data structure
	for schemaTablesRows.Next() {
		var schemaName, tableName sql.NullString
		if err := schemaTablesRows.Scan(&schemaName, &tableName); err != nil {
			return err
		}

		if schemaName.Valid {
			if tableName.Valid {
				mu.Lock()
				if _, exists := (*db)[postgresurlsha][types.Schema(schemaName.String)]; !exists {
					(*db)[postgresurlsha][types.Schema(schemaName.String)] = make(map[types.Table]map[types.Key]types.Value)
				}
				(*db)[postgresurlsha][types.Schema(schemaName.String)][types.Table(tableName.String)] = make(map[types.Key]types.Value)
				mu.Unlock()
			}
		}
	}

	if err := schemaTablesRows.Err(); err != nil {
		return err
	}
	return nil
}

func Addition(a int, b int) int {
	return a + b
}
