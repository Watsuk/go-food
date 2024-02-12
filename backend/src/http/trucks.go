package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Watsuk/go-food/src/auth"
	"github.com/Watsuk/go-food/src/permissions"
	"github.com/Watsuk/go-food/src/truck"
	"github.com/go-chi/chi"
)

func CreateTrucksEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Restaurateur, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		var newTruck newTruck // On crée une variable de type Truck
		err = json.NewDecoder(r.Body).Decode(&newTruck)

		if err != nil {
			log.Printf("Erreur de décodage JSON : %v", err)
			http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
			return
		}

		truck, err := truck.CreateTruck(db, newTruck.Name, newTruck.UserID, newTruck.SlotBuffer, newTruck.OpenTime, newTruck.CloseTime)

		if err != nil {
			log.Printf("Erreur lors la création du camion : %v", err)
			http.Error(w, "Erreur lors la création du camion", http.StatusBadRequest)
			return
		}

		jsonTruck, err := json.Marshal(truck)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTruck)
	}
}

func DeleteTruckEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Restaurateur, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		truckIDToDelete := chi.URLParam(r, "truckID")

		if truckIDToDelete == "" {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		truckID, err := strconv.ParseInt(truckIDToDelete, 10, 64)
		if err != nil {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		err = truck.DeleteTruck(db, truckID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "truck deleted successfully")
	}
}

func GetTrucksEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.User, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		trucks, err := truck.GetTrucks(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonTrucks, err := json.Marshal(trucks)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTrucks)
	}
}

func GetTruckByIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.User, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		truckIDString := chi.URLParam(r, "truckID")
		if truckIDString == "" {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		truckID, err := strconv.ParseInt(truckIDString, 10, 64)
		if err != nil {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		truck, err := truck.GetTruckByID(db, truckID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonTruck, err := json.Marshal(truck)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTruck)
	}
}

func GetTrucksByUserIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Restaurateur, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userIDString := chi.URLParam(r, "userID")
		if userIDString == "" {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		userID, err := strconv.ParseInt(userIDString, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		trucks, err := truck.GetTrucksByUserID(db, userID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonTrucks, err := json.Marshal(trucks)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTrucks)
	}
}

func EditTruckEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Restaurateur, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		truckIDString := chi.URLParam(r, "truckID")
		var newTruck newTruck
		err = json.NewDecoder(r.Body).Decode(&newTruck)

		if truckIDString == "" || err != nil {
			http.Error(w, "Invalid truck ID or data", http.StatusBadRequest)
			return
		}

		truckID, err := strconv.ParseInt(truckIDString, 10, 64)
		if err != nil {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		err = truck.EditTruck(db, truckID, newTruck.Name, newTruck.UserID, newTruck.SlotBuffer, newTruck.OpenTime, newTruck.CloseTime)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Truck edited"))
	}
}

func GetNumberCurrentOrdersByTruckIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.User, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		truckIDString := chi.URLParam(r, "truckID")
		if truckIDString == "" {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		truckID, err := strconv.ParseInt(truckIDString, 10, 64)
		if err != nil {
			http.Error(w, "Invalid truck ID", http.StatusBadRequest)
			return
		}

		nbrCurrentOrder, err := truck.NumberCurrentOrdersByTruckID(db, truckID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		response := struct {
			Count int `json:"count"`
		}{
			Count: nbrCurrentOrder,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

type newTruck struct {
	Name       string `json:"name"`
	UserID     int64  `json:"user_id"`
	SlotBuffer int64  `json:"slot_buffer"`
	OpenTime   string `json:"open_time"`
	CloseTime  string `json:"close_time"`
}
