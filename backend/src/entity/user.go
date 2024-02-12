package entity

import (
	"time"

	"github.com/Watsuk/go-food/src/permissions"
)

type User struct {
	ID        int64                  `json:"id"`
	Username  string                 `json:"username"`
	Password  string                 `json:"password"`
	Email     string                 `json:"email"`
	Role      permissions.Permission `json:"role"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
