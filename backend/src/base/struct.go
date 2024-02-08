package base

import "github.com/Watsuk/go-food/src/entity"

type Reference struct {
	User  *entity.User
	Truck *entity.Truck
	Chart *entity.Chart
	Order *entity.Order
}
