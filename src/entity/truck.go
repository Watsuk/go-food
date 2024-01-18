package entity

import "time"

type Truck struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	SlotBuffer int64     `json:"slot_buffer"`
	OpenTime   time.Time `json:"open_time"`
	CloseTime  time.Time `json:"close_time"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
