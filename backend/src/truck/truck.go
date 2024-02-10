package truck

import (
	"database/sql"
	"log"

	"github.com/Watsuk/go-food/src/entity"
)

func CreateTruck(db *sql.DB, name string, userID int64, slotBuffer int64, openTime string, closeTime string) (entity.Truck, error) {
	var truck entity.Truck
	query := "INSERT INTO trucks (name, user_id, slot_buffer, opening, closing) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, name, userID, slotBuffer, openTime, closeTime)
	if err != nil {
		log.Printf("Erreur lors de la création du camion : %v", err)
		return truck, err
	}
	truckID, _ := result.LastInsertId()
	truck.ID = truckID
	truck.Name = name
	truck.UserID = userID
	truck.SlotBuffer = slotBuffer
	truck.OpenTime = openTime
	truck.CloseTime = closeTime

	return truck, nil
}

func DeleteTruck(db *sql.DB, truckID int64) error {
	_, err := db.Exec("DELETE FROM trucks WHERE id = ?", truckID)
	return err
}
