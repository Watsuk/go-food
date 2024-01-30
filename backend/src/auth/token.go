package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ValidateTokenAndGetUserID(tokenString string) (int64, error) {
    // check if token is empty
    if tokenString == "" {
        return 0, fmt.Errorf("token is empty")
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte("VotreCléSecrète"), nil
    })

    if err != nil {
        return 0, err
    }

    // Check if token is valid
    if token == nil {
        return 0, fmt.Errorf("token is nil after parsing")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Check if the token contains the userId claim
        userID, ok := claims["userId"].(int64)
        if !ok {
            return 0, fmt.Errorf("userId claim is missing or not an int64")
        }
        return userID, nil
    }

    return 0, fmt.Errorf("invalid token")
}
