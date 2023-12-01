package controllers

import (
	"github.com/go-chi/chi/v5"
	"go-server/handlers"
)

func InitTaskRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", tasks.GetAllTasks)

	return r
}
