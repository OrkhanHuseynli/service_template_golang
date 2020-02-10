package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServiceHandler(t *testing.T) {
	handler := NewServiceHandler(NewServiceSubHandler())
	request := httptest.NewRequest("POST", "/product", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", response.Code)
	}
}