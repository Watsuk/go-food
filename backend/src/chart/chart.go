package chart

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
)

func GetCharts(db *sql.DB) ([]entity.Chart, error) {
	rows, err := db.Query("SELECT * FROM chart")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var charts []entity.Chart
	var createdAt []uint8
	var updatedAt []uint8
	var deletedAt []uint8
	for rows.Next() {
		var chart entity.Chart
		err := rows.Scan(&chart.ConsomableID, &chart.TruckID, &chart.Label, &chart.Description, &chart.Price, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		createdAtString := string(createdAt)

		chart.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			log.Fatal(err)
		}

		updatedAtString := string(updatedAt)
		chart.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
		if err != nil {
			log.Fatal(err)
		}
		if deletedAt != nil {
			deletedAtString := string(deletedAt)
			chart.DeletedAt, err = time.Parse("2006-01-02 15:04:05", deletedAtString)
			if err != nil {
				log.Fatal(err)
			}
		}

		charts = append(charts, chart)
	}

	return charts, err
}

func CreateChart(db *sql.DB, consomableID int64, truckID int64, label string, description string, price float64) (entity.Chart, error) {

	chart := entity.Chart{
		ConsomableID: consomableID,
		TruckID:      truckID,
		Label:        label,
		Description:  description,
		Price:        price,
	}

	_, err := db.Exec("INSERT INTO chart (consumable_id, truck_id, label, description, price) VALUES (?, ?, ?, ?, ?)", chart.ConsomableID, chart.TruckID, chart.Label, chart.Description, chart.Price)
	if err != nil {
		log.Printf("Erreur lors de la création du produit : %v", err)
		return entity.Chart{}, fmt.Errorf("could not create chart: %v", err)
	}

	return chart, err
}
