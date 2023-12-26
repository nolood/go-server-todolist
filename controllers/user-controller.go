package controllers

import (
	"go-server/handlers"

	"github.com/go-chi/chi/v5"
)

func InitUserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handlers.GetAllUsers)

	return r
}
