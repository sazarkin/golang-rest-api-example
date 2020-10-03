package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sazarkin/golang-rest-api-example/pkg/services"
)

type pockemonResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PokemonHandler Returns pokemon description by name
func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	desc, err := services.GetPokemonDesc(name)
	if err != nil {
		returnJSONError(w, err)
		return
	}

	descTranslated, err := services.TranslateToShakespeare(desc)
	if err != nil {
		returnJSONError(w, err)
		return
	}

	resp := pockemonResponse{
		Name:        name,
		Description: descTranslated,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
