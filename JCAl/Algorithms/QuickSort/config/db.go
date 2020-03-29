package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://localhost/archangel?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database.")
}
