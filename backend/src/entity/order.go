package entity

import (
	"time"
)

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	TruckID   int64     `json:"truck_id"`
	Price     int64     `json:"price"`
	Accepted  bool      `json:"accepted"`
	OrderData OrderData `json:"order_data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type OrderData struct {
	ID       int64     `json:"id"`
	Products []Product `json:"products"`
	Comment  string    `json:"comment"`
	Hour     time.Time `json:"hour"`
}
