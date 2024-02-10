package order

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func CreateOrder(db *sql.DB, userID int, truckID int, products []int, quantity []int, comment string, hour time.Time) (entity.Order, error) {
	var price int
	var productsDetails []entity.Product
	fmt.Println("ça marche de fou")
	for i, prod := range products {
		productDetail, err := product.GetProduct(db, int64(prod))
		if err != nil {
			return entity.Order{}, err
		}
		fmt.Println("ça marche pas")
		price += int(productDetail.Price) * quantity[i]
		productsDetails = append(productsDetails, productDetail)
	}

	orderData := entity.OrderData{
		Products: productsDetails,
		Comment:  comment,
		Hour:     hour,
	}
	orderDataJSON, err := json.Marshal(orderData)
	if err != nil {
		return entity.Order{}, err
	}
	fmt.Println("ça marche de fou de foufou")

	res, err := db.Exec("INSERT INTO orders (user_id, truck_id, price, hours, order_data) VALUES (?, ?, ?, ?, ?)", int64(userID), int64(truckID), int64(price), hour, orderDataJSON)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}

	fmt.Println("ça marche de fou furieux")
	orderID, err := res.LastInsertId()
	if err != nil {
		return entity.Order{}, err
	}
	order := entity.Order{
		ID:        orderID,
		UserID:    int64(userID),
		TruckID:   int64(truckID),
		Price:     int64(price),
		Hours:     hour,
		Accepted:  false,
		OrderData: orderData,
		CreatedAt: time.Now(),
	}
	fmt.Println("ça marche de fou malade")

	return order, nil
}

func GetOrderById(db *sql.DB, orderID int) (entity.Order, error) {
	var order entity.Order
	var hours []uint8
	var orderData []uint8
	var createdAt []uint8
	var updatedAt []uint8
	var deletedAt []uint8
	err := db.QueryRow("SELECT * FROM orders WHERE order_id = ?", orderID).Scan(&order.ID, &order.UserID, &order.TruckID, &order.Price, &hours, &order.Accepted, &order.Status, &orderData, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}
	hoursString := string(hours)
	order.Hours, err = time.Parse("15:04:05", hoursString)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}
	err = json.Unmarshal(orderData, &order.OrderData)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}
	createdAtString := string(createdAt)
	order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}
	updatedAtString := string(updatedAt)
	order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}
	return order, nil
}
