package controllers

import (
	"go-server/handlers"

	"github.com/go-chi/chi/v5"
)

func InitAuthRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", handlers.Login)
	r.Post("/register", handlers.Register)

	return r
}
