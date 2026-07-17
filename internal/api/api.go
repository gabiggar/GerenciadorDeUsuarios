package api

import (
	"GerenciadorDeUsuarios/internal/handlers"
	"GerenciadorDeUsuarios/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store repository.Store) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", handlers.PostUsers(store))
		r.Get("/users", handlers.GetUsers(store))
		r.Get("/users/{id}", handlers.GetUsersById(store))
		r.Put("/users/{id}", handlers.UpdateUsers(store))
		r.Delete("/users/{id}", handlers.DeleteUser(store))
	})

	return r
}
