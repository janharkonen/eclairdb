package postgresloader

import (
	"fmt"
)

func LoadData() error {
	fmt.Println("Loading data into PostgreSQL")
	fmt.Println(addition(1, 2))
	return nil
}

func addition(a int, b int) int {
	return a + b
}
