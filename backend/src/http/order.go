package http

import (
	"database/sql"
	"net/http"
	"strconv"

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
