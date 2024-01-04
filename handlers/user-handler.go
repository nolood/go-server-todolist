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

	err := postgres.Db.Model(&users).Select()
	if err != nil {
		log.Fatal(err)
	}

	//userId, ok := r.Context().Value("user_id").(uuid.UUID)
	//if !ok {
	//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//	return
	//}
	//
	//log.Println(userId)

	w.Write(toJson(users))
}

func CreateUser(user *postgres.User) error {

	isUser := postgres.User{}

	postgres.Db.Model(&isUser).Where("username = ?", user.Username).First()

	if isUser.Username != "" {
		return fmt.Errorf("user_exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = postgres.Db.Model(user).Insert()

	return err
}
