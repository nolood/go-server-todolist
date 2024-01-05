package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string, id uint64) (string, error) {
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

func Register(w http.ResponseWriter, r *http.Request) {
	var user postgres.User

	err := fromBody(r.Body, &user)
	log.Println(user)

	if err != nil {
		config.Logger.Error(err.Error())
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

	query := postgres.Db.Table("users")

	query = query.Where("username = ?", isUser.Username)

	query.Find(&user)

	if user.Username == "" {
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

func Vkminiapp(w http.ResponseWriter, r *http.Request) {
	var isUser postgres.User

	err := fromBody(r.Body, &isUser)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	query := postgres.Db.Table("users")

	query = query.Where("vk_id = ?", isUser.VKID)

	var user postgres.User

	query.Find(&user)

	if user.Username == "" {
		err = CreateUser(&isUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	token, err := generateToken(isUser.Username, isUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}
