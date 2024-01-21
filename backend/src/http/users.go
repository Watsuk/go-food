package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Watsuk/go-food/src/user"
)

func GetUsersEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func CreateUserEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser User

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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}
