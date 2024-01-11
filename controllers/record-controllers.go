package controllers

import (
	"github.com/go-chi/chi/v5"
	"go-server/handlers"
	"go-server/middlewares"
)

func InitRecordRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware())

	r.Post("/", handlers.CreateRecord)
	r.Get("/{billId}", handlers.GetRecordsByBillId)

	return r
}
