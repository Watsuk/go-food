package tokens

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

const TOKEN_EXPIRATION = int64(time.Hour * 24)

func ValidateTokenAndGetUserID(tokenString string) (int64, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}
