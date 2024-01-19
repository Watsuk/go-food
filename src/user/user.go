package user

import (
	"time"

	"github.com/Watsuk/go-food/src/entity"
)

func GetUsers() entity.User {
	return entity.User{
		Username:  "watsuk",
		Password:  "1234",
		Email:     "test@gmail.com",
		Role:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
