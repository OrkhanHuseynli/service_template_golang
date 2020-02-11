package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServiceHandlerWhenNoSearchCriteriaIsSet(t *testing.T) {
	request, response, handler :=setupTest(nil)

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", response.Code)
	}
}

func setupTest(d interface{})(*http.Request, *httptest.ResponseRecorder, http.Handler){
	handler := NewServiceHandler(NewServiceSubHandler())
	request := httptest.NewRequest("POST", "/product", nil)
	if d != nil {
		body, _ := json.Marshal(d)
		bodyReader :=  bytes.NewReader(body)
		request = httptest.NewRequest("POST", "/product", bodyReader)
	}
	response := httptest.NewRecorder()
	return request, response, handler
}