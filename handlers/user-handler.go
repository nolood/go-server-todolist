package handlers

import (
	"fmt"
	"go-server/internal/storage/postgres"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []postgres.User

	err := db.Model(&users).Select()
	if err != nil {
		log.Fatal(err)
	}

	w.Write(toJson(users))
}

func CreateUser(user *postgres.User) error {

	isUser := postgres.User{}

	db.Model(&isUser).Where("username = ?", user.Username).First()

	if isUser.Username != "" {
		return fmt.Errorf("user_exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = db.Model(user).Insert()

	return err
}
