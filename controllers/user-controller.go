package controllers

import (
	"go-server/handlers"
	"go-server/middlewares"

	"github.com/go-chi/chi/v5"
)

func InitUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.AuthMiddleware())

	r.Get("/", handlers.GetAllUsers)

	return r
}
