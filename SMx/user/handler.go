package users

import (
	"fmt"
	"net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	xu, err := AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	for _, u := range xu {
		fmt.Fprint(w, "%d, %s, %s, %s, %s, %s, %s, %s, %d", u.ID, u.FirstName, u.LastName, u.Username, u.Password, u.Email, u.CompanyId)
	}
}
