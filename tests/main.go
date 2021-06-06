package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	// Open sqlite3 database containing the bibles
	db, err := sql.Open("sqlite3", "../bible-sqlite.db")
	if err != nil {
		fmt.Println("Had trouble opening database")
		fmt.Println(err)
		os.Exit(1)
	}

	generateTests(db)
}
