package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	Valid_Name_Length = 32
	Valid_Age_Year    = 16
)

type userResponse struct {
	Message string `json:"message"`
}

type userRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type userPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type validationContextKey string

// Create struct validation for handler
// we need to have a reference as an input Param (next http.Handler Interface{}) to the next in the chain
// as it has the responsibility for calling ServeHTTP or returning a response.
type validationHandler struct {
	next http.Handler
}

// Return constructor object with arguement
// @para: next pass http.Handler to validation.
// @return: http.Handler for validation
func NewValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

// Implement handler for validation.
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request userRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	log.Printf("request %v\n", request)
	if err := h.validateName(request.Name); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validateAge(request.Age); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	userEntity := userPayload{
		Name: request.Name,
		Age:  request.Age,
	}

	// Step 1: WithValue() method to get the parent Context and associated with key.
	// request which has the context changed to the given ctx context.
	// It contains connection between client and server.
	c := context.WithValue(r.Context(), validationContextKey("user"), userEntity)
	r = r.WithContext(c)

	// From h - validation handler, pass ResponseWriter, Request to Validation Handler.
	h.next.ServeHTTP(rw, r)
}

func (h validationHandler) validateName(name string) error {
	N := strings.TrimSpace(name)
	if N == "" {
		return errors.New("Error:Name is required")
	}

	if len(N) > Valid_Name_Length {
		return fmt.Errorf("Error:Name length must not langer than %d char", Valid_Name_Length)
	}

	return nil
}

func (h validationHandler) validateAge(age int) error {
	//a, err := strconv.Atoi(age)
	//if err != nil {
	//	return err
	//}
	if age <= 0 {
		return errors.New("Error:Age is required")
	}
	if age <= Valid_Age_Year {
		return fmt.Errorf("Error:Valid age must not less than %d char", Valid_Age_Year)
	}

	return nil
}

/*************************************************************************************************/
/* Handler 2: Send to API Database and Reply message response. This is ServHTTP last chain
/*************************************************************************************************/

// Create struct response for handler
type userHandler struct{}

// Return constructor for new handler
func NewUserHandler() http.Handler {
	return userHandler{}
}

// Implement handler for response.
func (h userHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	usr := r.Context().Value(validationContextKey("user")).(userPayload)

	//if err:=database.Save(); err!=nil {
	//http.Error(rw, err.Error(), http.StatusBadRequest)
	//return
	//} //save to database

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	response := userResponse{Message: fmt.Sprintf("Saved user %s ", usr.Name)}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
