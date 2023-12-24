package routes

import (
	"go-server/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	users := controllers.InitUserRouter()
	auth := controllers.InitAuthRouter()

	r.Mount("/users", users)
	r.Mount("/auth", auth)

	return r
}
