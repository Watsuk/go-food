package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	TruckID   int64     `json:"truck_id"`
	Price     int64     `json:"price"`
	Accepted  bool      `json:"accepted"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderData struct {
	ID           int64     `json:"id"`
	ConsomableID int64     `json:"consomable_id"`
	Qty          int64     `json:"qty"`
	Comment      string    `json:"comment"`
	Hour         time.Time `json:"hour"`
	CreatedAt    time.Time `json:"created_at"`
}
