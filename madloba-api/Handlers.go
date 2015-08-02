package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"
	_ "github.com/lib/pq"
)


// Shows the content of all ads
// GET /ads
func AdIndex(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=etienne dbname=madloba password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select id, title, description from ads;")
	var ads Ads

	for rows.Next() {
		var id int
		var title string
		var description string
		err = rows.Scan(&id, &title, &description)
		ad := Ad{Id: id, Title: title, Description: description}
		ads = append(ads, ad)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ads); err != nil {
		panic(err)
	}
}

// Shows the content of one ad
// GET /ad/{adId}
func AdShow(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=etienne dbname=madloba password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	adId := vars["adId"]
	fmt.Println("Getting ad with ID", adId)

	rows, err := db.Query("select id, title, description from ads where id=$1;", adId)
	var id int
	var title, description string

	for rows.Next() {		
		err = rows.Scan(&id, &title, &description)
	}

	ad := Ad{Id: id, Title: title, Description: description}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ad); err != nil {
		panic(err)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Madloba API!")
}
