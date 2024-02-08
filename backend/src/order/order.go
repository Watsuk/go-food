package order

import (
	"database/sql"
	"time"
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
