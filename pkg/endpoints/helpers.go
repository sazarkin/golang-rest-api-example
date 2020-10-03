package endpoints

import (
	"encoding/json"
	"net/http"
)

type errorResp struct {
	Error string `json:"error"`
}

func returnJSONError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	resp := errorResp{
		Error: err.Error(),
	}
	json.NewEncoder(w).Encode(resp)
}
