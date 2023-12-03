package handlers

import (
	"go-server/internal/storage/postgres"
	"io"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	w.Write([]byte(string(body)))
}
