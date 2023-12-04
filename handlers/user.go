package handlers

import (
	"go-server/internal/storage/postgres"
	"go-server/models"
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

	var user models.CreateUserDto
	err := fromBody(r.Body, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Write(toJson(user))
}
