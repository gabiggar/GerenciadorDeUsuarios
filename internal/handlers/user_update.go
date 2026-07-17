package handlers

import (
	"GerenciadorDeUsuarios/internal/models"
	"GerenciadorDeUsuarios/internal/repository"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UpdateUsers(store repository.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsedID, err := uuid.Parse(id)
		if err != nil {
			sendJSON(w, Response{Error: "invalid id"}, http.StatusBadRequest)
			return
		}

		var body models.UpdateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid request body"}, http.StatusBadRequest)
			return
		}

		if body.FirstName == nil || body.LastName == nil || body.Biography == nil {
			sendJSON(w, Response{Error: "all fields are required"}, http.StatusBadRequest)
			return
		}
		userUpdated, err := store.Update(parsedID, body)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
				return
			}

			sendJSON(w, Response{Error: "unable to update user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: userUpdated}, http.StatusOK)
	}
}
