package company

import (
	"fmt"
	"github.com/archangel/SMx/config"
	_ "github.com/lib/pq"
)

type Company struct {
	ID          int
	Name        string
	Country     string
	City        string
	State       string
	Address     string
	PhoneNumber string
}

func init() {
	_, err := config.DB.Exec(
		"SELECT 1 " +
		"FROM pg_catalog.pg_class c " +
		"JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace " +
		"WHERE n.nspname = 'smx' " +
   		"AND c.relname = 'companies' " +
		"AND c.relkind = 'r' -- only tables",
		)

	if err != nil {
		fmt.Println("Companies Table Doesn't Exist, Creating....")
		seedCompanyData()
	}
}

func seedCompanyData() {
	_, err := config.DB.Exec("DROP TABLE companies")

	_, err = config.DB.Exec(
		"CREATE TABLE companies (" +
			"id SERIAL PRIMARY KEY, " +
			"name varchar(50) NOT NULL, " +
			"country varchar(25) NOT NULL, " +
			"city varchar(25), " +
			"state varchar(25), " +
			"address varchar(100), " +
			"phone_number varchar(15)" +
			")")

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Created Company Table")

	errs := make([]error, 0)
	_, err = config.DB.Exec("INSERT INTO companies (id, name, country, city, state, address, phone_number) VALUES (1, 'Solution Matrix', 'United States', 'Tucson', 'Arizona', '143 Beautiful Way', '520-445-3214')")
	errs = append(errs, err)
	_, err = config.DB.Exec("INSERT INTO companies (id, name, country, city, state, address, phone_number) VALUES (2, 'Knights', 'United States', 'Boise', 'Idaho', '123 Knights Way', '208-616-4741')")
	errs = append(errs, err)

	if err != nil {
		panic(errs)
	}
}

func AllCompanies() ([]*Company, error) {
	rows, err := config.DB.Query("SELECT * FROM companies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	xc := make([]*Company, 0)
	for rows.Next() {
		c := new(Company)
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Country,
			&c.City,
			&c.State,
			&c.Address,
			&c.PhoneNumber,
		)

		if err != nil {
			return nil, err
		}

		xc = append(xc, c)
	}

	return xc, nil
}
