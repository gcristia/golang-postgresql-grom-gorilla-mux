package routes

import (
	"encoding/json"
	"github.com/gcristia/golang-postgresql-grom-gorilla-mux/db"
	"github.com/gcristia/golang-postgresql-grom-gorilla-mux/models"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users !!"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User !!"))
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createduser := db.DB.Create(&user)
	err := createduser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete !!"))
}
