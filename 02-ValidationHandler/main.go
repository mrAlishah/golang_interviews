package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {

	// Use the http.NewServeMux() function to create an empty servemux.
	mux := http.NewServeMux()

	// I used Validation middleware handler to validate and
	// after that I use handler to save data(userPayload entity)
	// Note: just for professionally I used WithContext to send value between two handler
	handler := NewValidationHandler(NewUserHandler())
	mux.Handle("/save", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), mux))
}
