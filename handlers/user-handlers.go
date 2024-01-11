package handlers

import (
	"fmt"
	"go-server/internal/storage/postgres"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []postgres.User

	query := postgres.Db.Table("users")

	query = query.Find(&users)

	w.Write(toJson(users))
}

func CreateUser(user *postgres.User) error {

	isUser := postgres.User{}

	query := postgres.Db.Table("users")

	query = query.Where("username = ?", user.Username)

	query.Find(&isUser)

	if isUser.Username != "" {
		return fmt.Errorf("user_exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	postgres.Db.Table("users").Create(&user)

	return err
}
