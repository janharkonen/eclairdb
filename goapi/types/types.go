package types

type Sha string
type SchemaName string
type ColumnName string
type TableName string

type Value string
type Row map[ColumnName]Value
type Rows []Row
type Table struct {
	Done bool
	Rows Rows
}
type Tables map[TableName]Table
type Schemas map[SchemaName]Tables
type Databases map[Sha]Schemas

type DoneSchemaAndTableChannel chan [2]string
type DoneSchemaAndTableChannelMap map[Sha]DoneSchemaAndTableChannel
