package product

import (
	"database/sql"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
)

func GetProduct(db *sql.DB, productID int64) (entity.Product, error) {
	var product entity.Product
	rows, err := db.Query("SELECT id, truck_id, label, description, price, created_at, updated_at FROM product WHERE id = ?", productID)
	if err != nil {
		return entity.Product{}, err
	}
	for rows.Next() {
		var createdAt []uint8
		var updatedAt []uint8
		err = rows.Scan(&product.ID, &product.TruckID, &product.Label, &product.Description, &product.Price, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		createdAtString := string(createdAt)
		product.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			log.Fatal(err)
		}
		updatedAtString := string(updatedAt)
		product.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
		if err != nil {
			log.Fatal(err)
		}
	}
	return product, nil
}

func CreateProduct(db *sql.DB, truckID int, label string, description string, price int) (entity.Product, error) {
	res, err := db.Exec("INSERT INTO product (truck_id, label, description, price) VALUES (?, ?, ?, ?)", int64(truckID), label, description, int64(price))
	if err != nil {
		return entity.Product{}, err
	}

	productID, err := res.LastInsertId()
	if err != nil {
		return entity.Product{}, err
	}
	product := entity.Product{
		ID:          productID,
		TruckID:     int64(truckID),
		Label:       label,
		Description: description,
		Price:       int64(price),
	}
	return product, nil
}
