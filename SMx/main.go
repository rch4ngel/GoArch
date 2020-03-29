package main

// Created by John Nielsen
// Solution Matrix also known as SMx is going to be a product management tool. This tool will be used to gather data
// from customers to executives. The philosophy here is that to have an amazing product things flow bio-directionally.
// Customers have their problems and Executives have their problems on how to solve them.
// This application will bring down the vail and open up communication for all parties.

import (
	"github.com/archangel/SMx/company"
	"github.com/archangel/SMx/user"
	"net/http"
	"github.com/archangel/SMx/dashboard"
	"github.com/archangel/SMx/config"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/users", user.Index)
	http.HandleFunc("/companies", company.Index)
	http.HandleFunc("/dashboard", dashboard.Index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	config.TPL.ExecuteTemplate(w, "index.html", nil)
}
