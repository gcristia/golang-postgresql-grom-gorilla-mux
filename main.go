package main

import (
	"github.com/gcristia/golang-postgresql-grom-gorilla-mux/db"
	"github.com/gcristia/golang-postgresql-grom-gorilla-mux/models"
	"github.com/gcristia/golang-postgresql-grom-gorilla-mux/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteHandler).Methods("DELETE")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
