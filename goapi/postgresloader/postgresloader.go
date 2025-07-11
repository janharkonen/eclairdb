package postgresloader

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/janharkonen/eclairdb/types"
)

func LoadData(postgresurl string, db *types.Database) error {

	dbclient, err := sql.Open("postgres", postgresurl)
	if err != nil {
		return err
	}
	defer dbclient.Close()

	hasher := sha256.New()
	hasher.Write([]byte(postgresurl))
	postgresurlsha := types.Sha(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))
	fmt.Println(postgresurlsha)
	(*db)[postgresurlsha] = make(map[types.Schema]map[types.Table]map[types.Key]types.Value)

	schemas, err := dbclient.Query("SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT LIKE 'pg_%' AND schema_name != 'information_schema'")
	if err != nil {
		return err
	}
	defer schemas.Close()

	for schemas.Next() {
		var schema_name string
		err = schemas.Scan(&schema_name)
		if err != nil {
			return err
		}
		(*db)[postgresurlsha][types.Schema(schema_name)] = make(map[types.Table]map[types.Key]types.Value)
		tables, err := dbclient.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = '" + schema_name + "'")
		if err != nil {
			return err
		}
		defer tables.Close()
		for tables.Next() {
			var table_name string
			err = tables.Scan(&table_name)
			if err != nil {
				return err
			}
			(*db)[postgresurlsha][types.Schema(schema_name)][types.Table(table_name)] = make(map[types.Key]types.Value)
		}
	}
	return nil
}

func addition(a int, b int) int {
	return a + b
}
