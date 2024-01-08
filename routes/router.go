package routes

import (
	"github.com/go-chi/cors"
	"go-server/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	users := controllers.InitUserRouter()
	auth := controllers.InitAuthRouter()
	bills := controllers.InitBillRouter()
	articles := controllers.InitArticleRouter()

	r.Mount("/users", users)
	r.Mount("/auth", auth)
	r.Mount("/bills", bills)
	r.Mount("/articles", articles)

	return r
}
