package truck

import (
	"database/sql"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
)

func CreateTruck(db *sql.DB, name string, slotBuffer int64, openTime time.Time, closeTime time.Time) (entity.Truck, error) {
	var truck entity.Truck
	query := "INSERT INTO trucks (name, slot_buffer, open_time, close_time) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, name, slotBuffer, openTime, closeTime)
	if err != nil {
		log.Printf("Erreur lors de la création du camion : %v", err)
		return truck, err
	}
	truckID, _ := result.LastInsertId()
	truck.ID = truckID
	truck.Name = name
	truck.SlotBuffer = slotBuffer
	truck.OpenTime = openTime
	truck.CloseTime = closeTime

	return truck, nil
}

func DeleteTruck(db *sql.DB, truckID int64) error {
	_, err := db.Exec("DELETE FROM trucks WHERE id = ?", truckID)
	return err
}
