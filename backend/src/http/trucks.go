package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/Watsuk/go-food/src/truck"
	"github.com/go-chi/chi"
)

func CreateTrucksEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newTruck entity.Truck // On crée une variable de type Truck
		err := json.NewDecoder(r.Body).Decode(&newTruck)

		if err != nil {
			log.Printf("Erreur de décodage JSON : %v", err)
			http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
			return
		}

		truck, err := truck.CreateTruck(db, newTruck.Name, newTruck.SlotBuffer, newTruck.OpenTime, newTruck.CloseTime)

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
