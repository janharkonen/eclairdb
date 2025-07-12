package types

type Sha string
type Schema string
type Table string
type Column string
type Key string
type Value string
type Database map[Sha]map[Schema]map[Table]map[Key]Value
