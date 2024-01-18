package handler

import (
	"github.com/Watsuk/go-food/src/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandlerUser(user *entity.User) *HandlerUser {
	handlers := &HandlerUser{
		chi.NewRouter(),
		user,
	}

	handlers.Use(middleware.Logger)

	return handlers
}

type HandlerUser struct {
	*chi.Mux
	user *entity.User
}
