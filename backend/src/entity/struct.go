package entity

type Reference struct {
	User  *User
	Truck *Truck
	Chart *Chart
	Order *Order
}

type Product struct {
	ProductID Chart `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}
