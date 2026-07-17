package handlers

import (
	"GerenciadorDeUsuarios/internal/repository"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUsersById(store repository.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsed, err := uuid.Parse(id)
		if err != nil {
			sendJSON(w, Response{Error: "invalid id"}, http.StatusBadRequest)
			return
		}
		user, err := store.FindById(parsed)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
				return
			}

			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
