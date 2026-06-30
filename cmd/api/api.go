package api

import (
	"GerenciadorDeUsuarios/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", handlers.PostUsers())
		r.Get("/users", handlers.GetUsers())
		r.Get("/users/{id}", handlers.GetUsersById())
		r.Put("/users/{id}", handlers.UpdateUsers())
		r.Delete("/users/{id}", handlers.DeleteUser())
	})

	return r
}
