package handler

import (
	"database/sql"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/Watsuk/go-food/src/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandlerUser(db *sql.DB, user *entity.User) *HandlerUser {
	handlers := &HandlerUser{
		chi.NewRouter(),
		user,
	}

	handlers.Use(middleware.Logger)
	handlers.Get("/users", http.GetUsersEndpoint(db))
	handlers.Post("/user", http.CreateUserEndpoint(db))

	return handlers
}

type HandlerUser struct {
	*chi.Mux
	user *entity.User
}
