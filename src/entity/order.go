package entity

type Order struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	TruckID int64 `json:"truck_id"`
	Price   int64 `json:"price"`
}

type OrderData struct {
	ID           int64  `json:"id"`
	ConsomableID int64  `json:"consomable_id"`
	Qty          int64  `json:"qty"`
	Comment      string `json:"comment"`
}
