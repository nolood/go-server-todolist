package handlers

import (
	"fmt"
	"go-server/internal/storage/postgres"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string, id uuid.UUID) (string, error) {
	secret := viper.GetString("SECRET_KEY")

	claims := Claims{
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func jwtAuthenticator(token *jwt.Token) (interface{}, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("Can't get SECRET_KEY")
		return nil, fmt.Errorf("can't get SECRET_KEY")
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(secret), nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user postgres.User
	err := fromBody(r.Body, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := generateToken(user.Username, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var isUser postgres.User
	err := fromBody(r.Body, &isUser)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var user postgres.User

	err = postgres.Db.Model(&user).Where("username = ?", isUser.Username).Select()
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(isUser.Password))
	if err != nil {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}

	token, err := generateToken(user.Username, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}
