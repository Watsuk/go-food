package order

import (
	"database/sql"
	"time"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/Watsuk/go-food/src/product"
)

func AcceptOrder(db *sql.DB, orderID int, accept bool) (bool, error) {
	if accept {
		_, err := db.Exec("UPDATE orders SET accepted = ? WHERE id = ?", accept, orderID)
		if err != nil {
			return false, err
		}
	} else {
		_, err := db.Exec("UPDATE orders SET deleted_at = ?", time.Now())
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func CreateOrder(db *sql.DB, userID int, truckID int, products []entity.Chart, quantity []int, comment string, hour time.Time) (entity.Order, error) {
	var price int
	var prods []entity.Product

	for i, prod := range products {
		prod, err := product.GetProduct(db, prod.ConsomableID)
		if err != nil {
			return entity.Order{}, err
		}
		price += int(prod.Price) * quantity[i]
		product := entity.Product{
			ProductID: prod,
			Quantity:  int64(quantity[i]),
		}
		prods = append(prods, product)
	}

	productData := entity.OrderData{
		Products: prods,
		Comment:  comment,
		Hour:     hour,
	}

	res, err := db.Exec("INSERT INTO orders (user_id, truck_id, price, order_data) VALUES (?, ?, ?)", userID, truckID, price, productData)
	if err != nil {
		return entity.Order{}, err
	}
	orderID, err := res.LastInsertId()
	if err != nil {
		return entity.Order{}, err
	}
	order := entity.Order{
		ID:        orderID,
		UserID:    int64(userID),
		TruckID:   int64(truckID),
		Price:     int64(price),
		Accepted:  false,
		OrderData: productData,
		CreatedAt: time.Now(),
	}
	return order, nil
}
