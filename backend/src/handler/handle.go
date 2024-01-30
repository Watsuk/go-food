package handler

import (
	"database/sql"
	"net/http"

	"github.com/Watsuk/go-food/src/entity"
	myhttp "github.com/Watsuk/go-food/src/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func NewHandlerUser(db *sql.DB, user *entity.User) *HandlerUser {

	handlers := &HandlerUser{
		chi.NewRouter(),
		user,
	}

	// Cors middleware, the goal is to allow the front-end to access the API
	handlers.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Met l'URL de ton front-end ici
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	handlers.Use(middleware.Logger)

	handlers.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API Go-Food"))
	})

	handlers.Get("/users", myhttp.GetUsersEndpoint(db))
	handlers.Post("/register", myhttp.CreateUserEndpoint(db))
	handlers.Post("/login", myhttp.LoginEndpoint(db))
	handlers.Delete("/delete-account/{userID:[0-9]+}", myhttp.DeleteAccountEndpoint(db))

	handlers.Post("/create-truck", myhttp.CreateTrucksEndpoint(db))
	handlers.Delete("/delete-truck/{truckID:[0-9]+}", myhttp.DeleteTruckEndpoint(db))

	return handlers
}

type HandlerUser struct {
	*chi.Mux
	user *entity.User
}
