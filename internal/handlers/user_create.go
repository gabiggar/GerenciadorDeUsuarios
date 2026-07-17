package handlers

import (
	"GerenciadorDeUsuarios/internal/models"
	"GerenciadorDeUsuarios/internal/repository"
	"encoding/json"
	"net/http"
)

func PostUsers(store repository.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid request body"}, http.StatusBadRequest)
			return
		}

		if body.FirstName == "" || body.LastName == "" || body.Biography == "" {
			sendJSON(w, Response{Error: "missing required fields"}, http.StatusBadRequest)
			return
		}

		user, err := store.Insert(body)
		if err != nil {
			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusCreated)
	}
}
