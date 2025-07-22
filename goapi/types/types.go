package types

type Sha string
type SchemaName string
type ColumnName string
type TableName string

type Value string
type Row map[ColumnName]Value
type Table map[TableName][]Row
type Schema map[SchemaName]Table
type Database map[Sha]Schema
