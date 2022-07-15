package main

import (
	"github.com/gcristia/golang-postgresql-grom-gorilla-mux/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
