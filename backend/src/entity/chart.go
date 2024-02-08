package entity

import "time"

type Chart struct {
	ConsomableID int64     `json:"consumable_id"`
	TruckID      int64     `json:"truck_id"`
	Label        string    `json:"label"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
