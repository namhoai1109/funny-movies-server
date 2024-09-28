package main

import (
	"fmt"
	"funnymovies/internal/migration"
)

func main() {
	err := migration.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration completed")
}
