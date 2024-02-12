package entity

import "time"

type Truck struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	UserID     int64     `json:"user_id"`
	SlotBuffer int64     `json:"slot_buffer"`
	OpenTime   string    `json:"open_time"`
	CloseTime  string    `json:"close_time"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
