package handlers

import (
	"go-server/internal/storage/postgres"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []postgres.User

	err := db.Model(&users).Select()
	if err != nil {
		log.Fatal(err)
	}

	w.Write(toJson(users))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user postgres.User
	err := fromBody(r.Body, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = db.Model(&user).Returning("*").Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(toJson(user))
}
