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
	users := controllers.InitUserRouter()

	r.Mount("/tasks", tasks)
	r.Mount("/users", users)

	return r
}
