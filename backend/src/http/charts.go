package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Watsuk/go-food/src/chart"
)

func GetChartsEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		charts, err := chart.GetCharts(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des charts : %v", err)
			http.Error(w, "Erreur lors de la récupération des charts", http.StatusInternalServerError)
			return
		}
		jsonUser, err := json.Marshal(charts)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("ça marche pas trop trop mal")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	}

}
