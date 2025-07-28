package types

type Sha string
type SchemaName string
type ColumnName string
type TableName string

type Value string
type Row map[ColumnName]Value
type Tables map[TableName][]Row
type Schemas map[SchemaName]Tables
type Databases map[Sha]Schemas
