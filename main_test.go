package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	testValue := "novian_test"
	req, err := http.NewRequest("GET", fmt.Sprintf("/%s", testValue), nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(greetingHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := fmt.Sprintf(messageTemplate, testValue)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
