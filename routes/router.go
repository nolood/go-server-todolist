package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-server/controllers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tasks := controllers.InitTaskRouter()

	r.Mount("/tasks", tasks)

	return r
}
