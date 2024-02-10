package entity

type Reference struct {
	User    *User
	Truck   *Truck
	Product *Product
	Order   *Order
}

type Products struct {
	Product  Product `json:"product_id"`
	Quantity int64   `json:"quantity"`
}
