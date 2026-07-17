package main

import (
	"GerenciadorDeUsuarios/internal/api"
	"GerenciadorDeUsuarios/internal/database"
	"GerenciadorDeUsuarios/internal/repository"
	"log/slog"
	"net/http"
	"time"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	db, err := database.Open()
	if err != nil {
		return err
	}

	store := repository.NewStore(db)

	defer db.Close()

	handler := api.NewHandler(store)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
