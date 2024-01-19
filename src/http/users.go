package users

import (
	"encoding/json"
	"net/http"
)

func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	user := user.GetUser()
	// Convertir l'utilisateur en JSON et le renvoyer en réponse
	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}
