package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type user struct {
	ID        int
	Firstname string
	Lastname  string
	Username  string
	Password  []byte
	Role      string
}

func main() {
	fmt.Println("Connecting to Postgres Database.")
	db, err := sql.Open("postgres", "postgres://localhost/go_postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Postgres Database.")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	xu := []user{}
	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Username, &u.Password, &u.Role)
		if err != nil {
			panic(err)
		}
		xu = append(xu, u)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, v := range xu {
		fmt.Println(v.ID, v.Firstname)
	}
}
