package handlers

import (
	"GerenciadorDeUsuarios/dto"
	"GerenciadorDeUsuarios/models"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func PostUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body dto.CreateUserDTO

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid request body"}, http.StatusBadRequest)
			return
		}

		if body.FirstName == "" || body.LastName == "" || body.Biography == "" {
			sendJSON(w, Response{Error: "missing required fields"}, http.StatusBadRequest)
			return
		}

		createUser := dto.CreateUserDTO{
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Biography: body.Biography,
		}

		user, err := models.Insert(createUser)
		if err != nil {
			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusCreated)
	}
}

func GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.FindAll()
		if err != nil {
			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: users}, http.StatusOK)
	}
}

func GetUsersById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsed, err := uuid.Parse(id)
		if err != nil {
			sendJSON(w, Response{Error: "invalid id"}, http.StatusBadRequest)
			return
		}
		user, err := models.FindById(parsed)
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				sendJSON(w, Response{Error: "usuário não encontrado"}, http.StatusNotFound)
				return
			}

			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func UpdateUsersById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteUsersById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error ao fazer marshal de json", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error ao enviar a resposta", "error", err)
		return
	}
}
