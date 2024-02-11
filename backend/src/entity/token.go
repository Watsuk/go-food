package entity

import "time"

type Token struct {
	Token    string    `json:"token"`
	UserID   int64     `json:"user_id"`
	Lifetime time.Time `json:"lifetime"`
}
