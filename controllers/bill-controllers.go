package controllers

import (
	"github.com/go-chi/chi/v5"
	"go-server/handlers"
	"go-server/middlewares"
)

func InitBillRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.AuthMiddleware())

	r.Post("/", handlers.CreateBill)
	r.Get("/", handlers.GetAllBills)
	r.Get("/{id}", handlers.GetBill)

	return r
}
