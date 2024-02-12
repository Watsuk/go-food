package truck

import (
	"database/sql"
	"log"

	"github.com/Watsuk/go-food/src/entity"
)

func GetTrucks(db *sql.DB) ([]entity.Truck, error) {
	rows, err := db.Query("SELECT id, name, user_id, slot_buffer, opening, closing FROM trucks")
	if err != nil {
		log.Printf("Erreur lors de la récupération des camions : %v", err)
		return nil, err
	}
	defer rows.Close()

	var trucks []entity.Truck
	for rows.Next() {
		var truck entity.Truck
		err := rows.Scan(&truck.ID, &truck.Name, &truck.UserID, &truck.SlotBuffer, &truck.OpenTime, &truck.CloseTime)
		if err != nil {
			log.Printf("Erreur lors de la lecture des camions : %v", err)
			return nil, err
		}
		trucks = append(trucks, truck)
	}
	return trucks, nil
}

func GetTruckByID(db *sql.DB, truckID int64) (entity.Truck, error) {
	var truck entity.Truck
	err := db.QueryRow("SELECT * FROM trucks WHERE id = ?", truckID).Scan(&truck.ID, &truck.Name, &truck.UserID, &truck.SlotBuffer, &truck.OpenTime, &truck.CloseTime)
	if err != nil {
		log.Printf("Erreur lors de la récupération du camion : %v", err)
		return truck, err
	}
	return truck, nil
}

func GetTrucksByUserID(db *sql.DB, userID int64) ([]entity.Truck, error) {
	rows, err := db.Query("SELECT id, user_id, name, slot_buffer, opening, closing FROM trucks WHERE user_id = ?", userID)
	if err != nil {
		log.Printf("Erreur lors de la récupération des camions : %v", err)
		return nil, err
	}
	defer rows.Close()

	var trucks []entity.Truck
	for rows.Next() {
		var truck entity.Truck
		err := rows.Scan(&truck.ID, &truck.UserID, &truck.Name, &truck.SlotBuffer, &truck.OpenTime, &truck.CloseTime)
		if err != nil {
			log.Printf("Erreur lors de la lecture des camions : %v", err)
			return nil, err
		}
		trucks = append(trucks, truck)
	}
	return trucks, nil
}

func EditTruck(db *sql.DB, truckID int64, name string, userID int64, slotBuffer int64, openTime string, closeTime string) error {
	_, err := db.Exec("UPDATE trucks SET name = ?, user_id = ?, slot_buffer = ?, opening = ?, closing = ? WHERE id = ?",
		name, userID, slotBuffer, openTime, closeTime, truckID)
	return err
}

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

func NumberCurrentOrdersByTruckID(db *sql.DB, truckID int64) (int, error) {
	var response int
	err := db.QueryRow("SELECT COUNT(*) FROM orders WHERE truck_id = ? AND status IN ('pending', 'accepted', 'completed')", truckID).Scan(&response)
	if err != nil {
		log.Printf("Erreur lors de la récupération du camion : %v", err)
		return response, err
	}
	return response, nil
}
