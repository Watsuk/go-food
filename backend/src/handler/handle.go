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

func NewHandler(db *sql.DB, ref entity.Reference) *HandlerReference {

	handlers := &HandlerReference{
		chi.NewRouter(),
		ref.User,
		ref.Truck,
		ref.Product,
		ref.Order,
	}

	// Cors middleware, the goal is to allow the front-end to access the API
	handlers.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Met l'URL de ton front-end ici
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
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
	handlers.Get("/user/{userID:[0-9]+}", myhttp.GetUserByIdEndpoint(db))
	handlers.Post("/register", myhttp.CreateUserEndpoint(db))
	handlers.Post("/login", myhttp.LoginEndpoint(db))
	handlers.Patch("/delete-account/{userID:[0-9]+}", myhttp.DeleteAccountEndpoint(db))

	handlers.Post("/create-truck", myhttp.CreateTrucksEndpoint(db))
	handlers.Delete("/delete-truck/{truckID:[0-9]+}", myhttp.DeleteTruckEndpoint(db))

	handlers.Patch("/order/accept/{orderID:[0-9]+}/{accept:[0-1]}", myhttp.AcceptOrderEndpoint(db))
	handlers.Get("/order/{orderID:[0-9]+}", myhttp.GetOrdersByIdEndpoint(db))
	handlers.Post("/order", myhttp.CreateOrderEndpoint(db))
	handlers.Get("/orders/truck/{truckID:[0-9]+}", myhttp.GetOrdersByTruckEndpoint(db))
	handlers.Get("/orders/user/{userID:[0-9]+}", myhttp.GetOrdersByUserEndpoint(db))
	handlers.Get("/orders", myhttp.GetOrdersEndpoint(db))
	handlers.Patch("/order/{orderID:[0-9]+}/completed", myhttp.CompletedOrderEndpoint(db))
	handlers.Patch("/order/{orderID:[0-9]+}/handedover", myhttp.HandedOverOrderEndpoint(db))

	handlers.Post("/product", myhttp.CreateProductEndpoint(db))
	handlers.Get("/product/{productID:[0-9]+}", myhttp.GetProductByIdEndpoint(db))
	handlers.Get("/products/truck/{truckID:[0-9]+}", myhttp.GetProductsByTruckEndpoint(db))

	handlers.Patch("/users/{userID:[0-9]+}", myhttp.AdminEditEndpoint(db))
	handlers.Delete("/users/{userID:[0-9]+}", myhttp.AdminDeleteEndpoint(db))

	handlers.Get("/trucks", myhttp.GetTrucksEndpoint(db))
	handlers.Get("/trucks/{truckID:[0-9]+}", myhttp.GetTruckByIDEndpoint(db))
	handlers.Get("/trucks/user/{userID:[0-9]+}", myhttp.GetTrucksByUserIDEndpoint(db))
	handlers.Patch("/trucks/{truckID:[0-9]+}", myhttp.EditTruckEndpoint(db))

	return handlers
}

type HandlerReference struct {
	*chi.Mux
	user  *entity.User
	truck *entity.Truck
	chart *entity.Product
	order *entity.Order
}
