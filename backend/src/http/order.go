package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Watsuk/go-food/src/order"
	"github.com/go-chi/chi"
)

func AcceptOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")
		acceptString := chi.URLParam(r, "accept")

		if orderIDString == "" || acceptString == "" {
			http.Error(w, "Invalid order ID or accept value", http.StatusBadRequest)
			return
		}

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		accept, err := strconv.ParseBool(acceptString)
		if err != nil {
			http.Error(w, "Invalid accept value", http.StatusBadRequest)
			return
		}

		accepted, err := order.AcceptOrder(db, orderID, accept)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if accepted {
			w.Write([]byte("Order accepted"))
		} else {
			w.Write([]byte("Order declined"))
		}
	}
}

func CreateOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newOrder OrderBody
		err := json.NewDecoder(r.Body).Decode(&newOrder)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		order, err := order.CreateOrder(db, newOrder.UserID, newOrder.TruckID, newOrder.Products, newOrder.Quantity, newOrder.Comment, newOrder.Hour)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order created with ID: " + strconv.Itoa(int(order.ID))))
	}
}

func GetOrdersByIdEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")
		if orderIDString == "" {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		order, err := order.GetOrderById(db, orderID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(order)
	}
}

func GetOrdersByTruckEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		truckIDString := chi.URLParam(r, "truckID")
		if truckIDString == "" {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		truckID, err := strconv.Atoi(truckIDString)
		if err != nil {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		orders, err := order.GetOrdersByTruck(db, truckID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}

func GetOrdersByUserEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDString := chi.URLParam(r, "userID")
		if userIDString == "" {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		orders, err := order.GetOrdersByUser(db, userID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}

func GetOrdersEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := order.GetOrders(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}

func CompletedOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")
		if orderIDString == "" {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		err = order.CompletedOrder(db, orderID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order completed"))
	}
}

func HandedOverOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")
		if orderIDString == "" {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}

		err = order.HandedOverOrder(db, orderID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order handed over"))
	}
}

type OrderBody struct {
	UserID   int       `json:"user_id"`
	TruckID  int       `json:"truck_id"`
	Products []int     `json:"product"`
	Quantity []int     `json:"quantity"`
	Comment  string    `json:"comment"`
	Hour     time.Time `json:"hour"`
}
