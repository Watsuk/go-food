package product

import (
	"database/sql"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
)

func GetProduct(db *sql.DB, productID int64) (entity.Product, error) {
	var product entity.Product
	rows, err := db.Query("SELECT id, truck_id, naùe, label, description, price, created_at, updated_at FROM product WHERE id = ?", productID)
	if err != nil {
		return entity.Product{}, err
	}
	for rows.Next() {
		var createdAt []uint8
		var updatedAt []uint8
		err = rows.Scan(&product.ID, &product.TruckID, &product.Name, &product.Label, &product.Description, &product.Price, &createdAt, &updatedAt)
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

func CreateProduct(db *sql.DB, truckID int, name string, label string, description string, price int) (entity.Product, error) {
	res, err := db.Exec("INSERT INTO product (truck_id, name, label, description, price) VALUES (?, ?, ?, ?, ?)", int64(truckID), name, label, description, int64(price))
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
		Name:        name,
		Label:       label,
		Description: description,
		Price:       int64(price),
	}
	return product, nil
}

func GetProductsByTruckID(db *sql.DB, truckID int64) ([]entity.Product, error) {
	var product entity.Product
	var products []entity.Product
	rows, err := db.Query("SELECT id, truck_id, name, label, description, price, created_at, updated_at FROM product WHERE truck_id = ?", truckID)
	if err != nil {
		return []entity.Product{}, err
	}
	for rows.Next() {
		var createdAt []uint8
		var updatedAt []uint8
		err = rows.Scan(&product.ID, &product.TruckID, &product.Name, &product.Label, &product.Description, &product.Price, &createdAt, &updatedAt)
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
		products = append(products, product)
	}

	return products, nil
}

func DeleteProduct(db *sql.DB, productID int64) error {
	_, err := db.Exec("DELETE FROM product WHERE id = ?", productID)
	if err != nil {
		return err
	}
	return nil
}
