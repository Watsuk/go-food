package entity

import (
	"time"
)

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	TruckID   int64     `json:"truck_id"`
	Price     int64     `json:"price"`
	OrderData OrderData `json:"order_data"`
	Status    string    `json:"status"`
	Hours     time.Time `json:"hour"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderData struct {
	ID       int64     `json:"id"`
	Products []Product `json:"products"`
	Comment  string    `json:"comment"`
	Hour     time.Time `json:"hour"`
}
