package entity

type Chart struct {
	ConsomableID int64   `json:"consomable_id"`
	TruckID      int64   `json:"truck_id"`
	Label        string  `json:"label"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
}
