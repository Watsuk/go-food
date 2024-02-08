package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Watsuk/go-food/src/entity"
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

		order, err := order.CreateOrder(db, newOrder.UserID, truckID, newOrder.Products, newOrder.Quantity, newOrder.Comment, newOrder.Hour)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order created with ID: " + strconv.Itoa(int(order.ID))))
	}
}

type OrderBody struct {
	UserID   int            `json:"user_id"`
	TruckID  int            `json:"truck_id"`
	Products []entity.Chart `json:"product"`
	Quantity []int          `json:"quantity"`
	Comment  string         `json:"comment"`
	Hour     time.Time      `json:"hour"`
}
