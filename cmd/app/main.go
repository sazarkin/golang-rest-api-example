package main

import (
	"log"

	"github.com/sazarkin/golang-rest-api-example/pkg/server"
)

func main() {
	srv := server.InitServer("0.0.0.0:8080")
	log.Println("Service started running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
