package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(store *database.Store) *Handler {
	handlers := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	return &Handler{}
}

type Handler struct {
	*chi.Mux
	store *database.Store
}
