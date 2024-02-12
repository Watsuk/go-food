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
		_, err := db.Exec("UPDATE orders SET status = ?, updated_at = ? WHERE order_id = ?", "accepted", time.Now(), orderID)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
	} else {
		_, err := db.Exec("UPDATE orders SET status = ?, updated_at = ? WHERE order_id = ?", "rejected", time.Now(), orderID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func CreateOrder(db *sql.DB, userID int, truckID int, products []int, quantity []int, comment string, hour time.Time) (entity.Order, error) {
	var price int
	var productsDetails []entity.Product
	for i, prod := range products {
		productDetail, err := product.GetProduct(db, int64(prod))
		if err != nil {
			fmt.Println(err)
			return entity.Order{}, err
		}
		if productDetail.TruckID != int64(truckID) {
			fmt.Println(err)
			return entity.Order{}, fmt.Errorf("Product %d does not belong to truck %d", prod, truckID)
		}

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
		fmt.Println(err)
		return entity.Order{}, err
	}

	res, err := db.Exec("INSERT INTO orders (user_id, truck_id, price, hours, order_data) VALUES (?, ?, ?, ?, ?)", int64(userID), int64(truckID), int64(price), hour, orderDataJSON)
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return entity.Order{}, err
	}
	order := entity.Order{
		ID:        orderID,
		UserID:    int64(userID),
		TruckID:   int64(truckID),
		Price:     int64(price),
		Hours:     hour,
		OrderData: orderData,
		CreatedAt: time.Now(),
	}

	return order, nil
}

func GetOrderById(db *sql.DB, orderID int) (entity.Order, error) {
	var order entity.Order
	var hours []uint8
	var orderData []uint8
	var createdAt []uint8
	var updatedAt []uint8
	err := db.QueryRow("SELECT * FROM orders WHERE order_id = ?", orderID).Scan(&order.ID, &order.UserID, &order.TruckID, &order.Price, &hours, &order.Status, &orderData, &createdAt, &updatedAt)
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

func GetOrdersByTruck(db *sql.DB, truckID int) ([]entity.Order, error) {
	var orders []entity.Order
	rows, err := db.Query("SELECT * FROM orders WHERE truck_id = ?", truckID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order entity.Order
		var hours []uint8
		var orderData []uint8
		var createdAt []uint8
		var updatedAt []uint8
		err := rows.Scan(&order.ID, &order.UserID, &order.TruckID, &order.Price, &hours, &order.Status, &orderData, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		hoursString := string(hours)
		order.Hours, err = time.Parse("15:04:05", hoursString)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(orderData, &order.OrderData)
		if err != nil {
			return nil, err
		}
		createdAtString := string(createdAt)
		order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			return nil, err
		}
		updatedAtString := string(updatedAt)
		order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrdersByUser(db *sql.DB, userID int) ([]entity.Order, error) {
	var orders []entity.Order
	rows, err := db.Query("SELECT * FROM orders WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order entity.Order
		var hours []uint8
		var orderData []uint8
		var createdAt []uint8
		var updatedAt []uint8
		err := rows.Scan(&order.ID, &order.UserID, &order.TruckID, &order.Price, &hours, &order.Status, &orderData, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		hoursString := string(hours)
		order.Hours, err = time.Parse("15:04:05", hoursString)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(orderData, &order.OrderData)
		if err != nil {
			return nil, err
		}
		createdAtString := string(createdAt)
		order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			return nil, err
		}
		updatedAtString := string(updatedAt)
		order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrders(db *sql.DB) ([]entity.Order, error) {
	var orders []entity.Order
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order entity.Order
		var hours []uint8
		var orderData []uint8
		var createdAt []uint8
		var updatedAt []uint8
		err := rows.Scan(&order.ID, &order.UserID, &order.TruckID, &order.Price, &hours, &order.Status, &orderData, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		hoursString := string(hours)
		order.Hours, err = time.Parse("15:04:05", hoursString)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(orderData, &order.OrderData)
		if err != nil {
			return nil, err
		}
		createdAtString := string(createdAt)
		order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			return nil, err
		}
		updatedAtString := string(updatedAt)
		order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func CompletedOrder(db *sql.DB, orderID int) error {
	order, err := GetOrderById(db, orderID)
	if order.Status != "accepted" {
		return fmt.Errorf("Order not accepted")
	}
	_, err = db.Exec("UPDATE orders SET status = ?, updated_at = ? WHERE order_id = ?", "completed", time.Now(), orderID)
	if err != nil {
		return err
	}
	return nil
}

func HandedOverOrder(db *sql.DB, orderID int) error {
	order, err := GetOrderById(db, orderID)
	if order.Status != "completed" {
		return fmt.Errorf("Order not completed")
	}
	_, err = db.Exec("UPDATE orders SET status = ?, updated_at = ? WHERE order_id = ?", "handedover", time.Now(), orderID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(db *sql.DB, orderID int) error {
	_, err := db.Exec("DELETE FROM orders WHERE order_id = ?", orderID)
	if err != nil {
		return err
	}
	return nil
}
