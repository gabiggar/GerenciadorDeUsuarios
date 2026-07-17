package handlers

import (
	"GerenciadorDeUsuarios/internal/repository"
	"net/http"
)

func GetUsers(store repository.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := store.FindAll()
		if err != nil {
			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: users}, http.StatusOK)
	}
}
