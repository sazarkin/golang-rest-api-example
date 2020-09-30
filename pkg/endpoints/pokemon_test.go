// endpoints_test.go
package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthCheckHandler(t *testing.T) {
	name := "charizard"
	req, err := http.NewRequest("GET", fmt.Sprintf("/pokemon/%s", name), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/pokemon/{name}", PokemonHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}

	var resp pockemonResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to unmarshal response, got error %v", err)
	}

	if resp.Name != name {
		t.Errorf("expected response with name `%s`, got `%s`", name, resp.Name)
	}
}
