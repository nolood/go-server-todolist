package controllers

import (
	"github.com/go-chi/chi/v5"
	"go-server/handlers"
	"go-server/middlewares"
)

func InitArticleRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.AuthMiddleware())

	r.Post("/", handlers.CreateArticle)
	r.Get("/", handlers.GetArticles)

	return r
}
