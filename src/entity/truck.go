package entity

type Truck struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	SlotBuffer int64     `json:"slot_buffer"`
	OpenTime   time.time `json:"open_time"`
	CloseTime  time.time `json:"close_time"`
}
