package user

import (
	"fmt"
	"github.com/archangel/SMx/config"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	Role      string
	CompanyId int
}

func init() {
	_, err := config.DB.Exec(
			"SELECT 1 " +
			"FROM pg_catalog.pg_class c " +
			"JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace " +
			"WHERE n.nspname = 'smx' " +
			"AND c.relname = 'users' " +
			"AND c.relkind = 'r' -- only tables",
			)

	if err != nil {
		fmt.Println("User Table Doesn't Exist, Creating....")
		seedUserData()
	}
}

func seedUserData() {
	_, err := config.DB.Exec("CREATE TABLE users (" +
		"id SERIAL PRIMARY KEY NOT NULL, " +
		"first_name varchar(25) NOT NULL, " +
		"last_name varchar(25) NOT NULL, " +
		"username varchar(25) NOT NULL, " +
		"password varchar(100) NOT NULL, " +
		"email varchar(50) NOT NULL, " +
		"role varchar(25) NOT NULL, " +
		"company_id INTEGER, " +
		"FOREIGN KEY (company_id) REFERENCES companies(id)" +
		")")
	if err != nil {
		panic(err)
	}

	fmt.Println("User Table Successfully Created!")

	_, err = config.DB.Exec("INSERT INTO users (id, first_name, last_name, username, password, email, role, company_id) VALUES (1, 'Jubei', 'Kibagami', 'Samurai','vagabond','jubei.kibagami@ninja.com', 'Customer', 1)")
	_, err = config.DB.Exec("INSERT INTO users (id, first_name, last_name, username, password, email, role, company_id) VALUES (2, 'Clark', 'Kent', 'CKent','Superman','clark.kent@dc.com', 'Customer', 2)")

	if err != nil {
		panic(err)
	}
}

func AllUsers() ([]*User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	xu := make([]*User, 0)
	for rows.Next() {
		u := new(User)
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Password,
			&u.Email,
			&u.Role,
			&u.CompanyId,
		)

		if err != nil {
			return nil, err
		}

		xu = append(xu, u)
	}

	return xu, nil
}
