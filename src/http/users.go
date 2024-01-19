package http

import (
	"encoding/json"
	"net/http"

	"github.com/Watsuk/go-food/src/user"
)

func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	user := user.GetUsers()
	// Convertir l'utilisateur en JSON et le renvoyer en réponse
	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}
