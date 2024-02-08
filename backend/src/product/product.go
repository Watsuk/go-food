package product

import (
	"database/sql"

	"github.com/Watsuk/go-food/src/entity"
)

func GetProduct(db *sql.DB, productID int64) (entity.Chart, error) {
	var product entity.Chart
	err := db.QueryRow("SELECT * FROM products WHERE id = ?", productID).Scan(&product)
	if err != nil {
		return entity.Chart{}, err
	}
	return product, nil
}
