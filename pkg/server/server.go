package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/sazarkin/golang-rest-api-example/pkg/endpoints"
	"github.com/sazarkin/golang-rest-api-example/pkg/middleware"
)

// InitServer configures http server instance with all endpoints and returns it
func InitServer(addr string) *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/pokemon/{name}", endpoints.PokemonHandler)

	router.Use(middleware.LoggingMiddleware)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}
