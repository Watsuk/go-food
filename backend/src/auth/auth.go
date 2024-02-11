package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Watsuk/go-food/src/permissions"
	"github.com/Watsuk/go-food/src/token"
	"github.com/Watsuk/go-food/src/tokens"
	"github.com/Watsuk/go-food/src/user"
)

func CheckPerms(perms permissions.Permission, w http.ResponseWriter, r *http.Request, db *sql.DB) (bool, error) {
	// Check if the user has the required permissions
	hasPerms, err := HasPerms(perms, r, db)
	if err != nil {
		fmt.Println("Error checking permissions: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return false, err
	}
	if !hasPerms {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false, nil
	}
	return true, nil

}

func HasPerms(perms permissions.Permission, r *http.Request, db *sql.DB) (bool, error) {
	//get the token from the request heade Authorization
	tk := r.Header.Get("Authorization")

	//check if we got the token in Token slice
	tokenValid := false
	Tokens, err := token.GetTokens(db)
	for _, t := range Tokens {
		if t == tk {
			tokenValid = true
			break
		}
	}

	//if the token is not valid return false
	if !tokenValid {
		fmt.Println("Token not valid", err)
		return false, nil
	}

	//read user id from JWT
	uid, err := tokens.ValidateTokenAndGetUserID(tk)
	if err != nil {
		fmt.Println("Error validating token: ", err)
		return false, err
	}

	//check if the user has the required permissions
	user, err := user.GetUserByID(db, int(uid))
	if err != nil {
		fmt.Println("Error getting user by id: ", err)
		return false, err
	}

	if user.Role != perms {
		return false, nil
	}

	return true, nil
}
