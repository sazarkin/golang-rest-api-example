package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type pockemonResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PokemonHandler Returns pokemon description by name
func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp := pockemonResponse{
		Name:        vars["name"],
		Description: "empty",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
