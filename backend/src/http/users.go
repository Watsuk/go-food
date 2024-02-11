package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Watsuk/go-food/src/auth"
	"github.com/Watsuk/go-food/src/permissions"
	"github.com/Watsuk/go-food/src/token"
	"github.com/Watsuk/go-food/src/user"
	"github.com/go-chi/chi"
)

func GetUsersEndpoint(db *sql.DB) http.HandlerFunc {
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
		// Récupérer l'utilisateur
		users, err := user.GetUsers(db)
		if err != nil {
			log.Printf("Erreur lors de la récupération des utilisateurs : %v", err)
			http.Error(w, "Erreur lors de la récupération des utilisateurs", http.StatusInternalServerError)
			return
		}
		// Convertir l'utilisateur en JSON et le renvoyer en réponse
		jsonUser, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println("ça marche pas trop mal")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	}
}

func GetUserByIdEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.User, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Récupérer l'ID de l'utilisateur
		userIDString := chi.URLParam(r, "userID")

		// Convertir l'ID de l'utilisateur en entier
		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}

		// Récupérer l'utilisateur par son ID
		user, err := user.GetUserByID(db, userID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.NotFound(w, r)
			return
		}

		// Convertir l'utilisateur en JSON et envoyer la réponse
		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	}
}

func CreateUserEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser UserBody

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			log.Printf("Erreur de décodage JSON : %v", err)
			http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
			return
		}

		user, err := user.CreateUser(db, newUser.Username, newUser.Password, newUser.Email, newUser.Role)
		if err != nil {
			log.Printf("Erreur lors de la création de l'utilisateur : %v", err)
			http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
			return
		}
		// Convertir l'utilisateur en JSON et le renvoyer en réponse
		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	}
}

func LoginEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&loginData)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		userID, tk, err := user.Login(db, loginData.Email, loginData.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		response := struct {
			UserID int64  `json:"userId"`
			Token  string `json:"token"`
		}{
			UserID: userID,
			Token:  tk,
		}

		err = token.AddToken(db, tk, userID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteAccountEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		perm, err := auth.CheckPerms(permissions.User, w, r, db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !perm {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userIDToDelete := chi.URLParam(r, "userID")

		if userIDToDelete == "" {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		userID, err := strconv.ParseInt(userIDToDelete, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		err = user.DeleteUser(db, userID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "User account deleted successfully")
	}
}

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}
