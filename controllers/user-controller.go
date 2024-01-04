package controllers

import (
	"github.com/go-chi/chi/v5"
	"go-server/handlers"
	"go-server/middlewares"
)

func InitUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.AuthMiddleware())

	r.Get("/", handlers.GetAllUsers)

	return r
}
