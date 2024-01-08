package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type CreateArticleParam struct {
	gorm.Model
	ID        uint64 `json:"id" gorm:"autoIncrement"`
	Icon      string `json:"icon"`
	Title     string `json:"title"`
	Color     string `json:"color"`
	IsDefault bool   `json:"default"`
	UserID    uint64 `json:"user_id"`
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var article CreateArticleParam

	article.UserID = userId
	article.IsDefault = false

	err = fromBody(r.Body, &article)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Println(len(article.Title))

	if len(article.Title) > 38 {
		http.Error(w, "Title should be less than 20", http.StatusBadRequest)
		return
	}

	query := postgres.Db.Table("articles")

	var count int64

	query.Where("user_id = ?", userId).Count(&count)

	if count >= 5 {
		http.Error(w, "You can't create more than 5 articles", http.StatusBadRequest)
		return
	}

	log.Println(article)

	query = query.Create(&article)

	w.Write(toJson(article))
}

type GetArticleParam struct {
	ID        uint64 `json:"id"`
	Icon      string `json:"icon"`
	Title     string `json:"title"`
	Color     string `json:"color"`
	IsDefault bool   `json:"default"`
	UserID    uint64 `json:"user_id"`
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var articles []GetArticleParam

	query := postgres.Db.Table("articles")

	query = query.Where("is_default = ?", true).Or("user_id = ?", userId)

	query.Find(&articles)

	w.Write(toJson(articles))
}
