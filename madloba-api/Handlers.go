package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"database/sql"
	_ "github.com/lib/pq"
)

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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}