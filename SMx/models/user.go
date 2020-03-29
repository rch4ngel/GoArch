package models

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	Role      string
	*Company
}

func (db *DB) seedUsers() {
	_, err := db.Exec("CREATE TABLE users (" +
		"id SERIAL PRIMARY KEY NOT NULL, " +
		"first_name varchar(25) NOT NULL, " +
		"last_name varchar(25) NOT NULL, " +
		"username varchar(25) NOT NULL, " +
		"password varchar(100) NOT NULL, " +
		"email varchar(50) NOT NULL, " +
		"role varchar(25) NOT NULL, " +
		"company_id INTEGER REFERENCES companies(id)" +
		")")
	if err != nil {
		panic(err)
	}

	fmt.Println("User Table Successfully Created!")
}