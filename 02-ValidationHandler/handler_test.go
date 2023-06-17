package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// validate that search criteria
// go test -v -race .
func TestValidationHandlerReturnsBadRequestWhenNoNameIsSent(t *testing.T) {
	//Arrange = Setup
	r, rw, handler := setupTest(&userRequest{Age: 13})

	//Act = Execute
	handler.ServeHTTP(rw, r)
	t.Log(rw.Code, rw.Body)
	//Assert
	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestValidationHandlerReturnsBadRequestWhenInvalidAgeIsSent(t *testing.T) {
	//Arrange = Setup
	r, rw, handler := setupTest(&userRequest{Name: "Alishah", Age: 13})

	//Act = Execute
	handler.ServeHTTP(rw, r)
	t.Log(rw.Code, rw.Body)
	//Assert
	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestValidationHandlerValidAndSucceedIsSent(t *testing.T) {
	//Arrange = Setup
	r, rw, handler := setupTest(&userRequest{Name: "Alishah", Age: 20})

	//Act = Execute
	handler.ServeHTTP(rw, r)
	t.Log(rw.Code, rw.Body)
	//Assert
	if rw.Code != http.StatusCreated {
		t.Errorf("Expecte created user successfully got %v", rw.Code)
	}
}

// Arrange = Setup
func setupTest(data interface{}) (*http.Request, *httptest.ResponseRecorder, http.Handler) {
	/***** create HandlerSearch *****/
	handler := NewValidationHandler(NewUserHandler())

	/***** create httptest.ResponseRecorder ~ http.ResponseWriter *****/
	res := httptest.NewRecorder()

	if data == nil {
		// httptest.NewRequest ~ http.Request
		return httptest.NewRequest("POST", "/save", nil), res, handler
	}

	/***** create httptest.NewRequest by passing some JSON in the request body *****/
	body, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/save", bytes.NewReader(body))

	return req, res, handler
}
