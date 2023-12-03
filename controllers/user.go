package controllers

import (
	"github.com/go-chi/chi/v5"
	"go-server/handlers"
)

func InitUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handlers.GetAllUsers)
	r.Post("/", handlers.CreateUser)

	return r
}
