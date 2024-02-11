package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Watsuk/go-food/src/auth"
	"github.com/Watsuk/go-food/src/permissions"
	"github.com/Watsuk/go-food/src/user"
	"github.com/go-chi/chi"
)

func AdminEditEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Admin, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userIDString := chi.URLParam(r, "userID")
		var userModel AdminUserBody
		err = json.NewDecoder(r.Body).Decode(&userModel)

		if userIDString == "" || err != nil {
			http.Error(w, "Invalid user ID or data", http.StatusBadRequest)
			return
		}

		userID, err := strconv.ParseInt(userIDString, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		err = user.EditUser(db, userID, userModel.Username, userModel.Email, userModel.Role)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("User edited"))
	}
}

func AdminDeleteEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.Admin, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userIDString := chi.URLParam(r, "userID")
		if userIDString == "" {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		userID, err := strconv.ParseInt(userIDString, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		err = user.DeleteUser(db, userID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("User deleted"))
	}
}

type AdminUserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}
