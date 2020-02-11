package handlers

import (
	"github.com/microservice_template/src/models"
	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServiceHandlerWhenNoSearchCriteriaIsSet(t *testing.T) {
	request, response, handler := setupTest(nil)

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", response.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T){
	request, response, handler := setupTest(&models.SimpleRequest{})

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected StatusUnprocessableEntity got %v", response.Code)
	}
}

func TestSearchHandlerCallsDataStoreWithValidQuery(t *testing.T) {
	productName := "Golden Coin"
	request, response, handler := setupTest(&models.SimpleRequest{productName})
	handler.ServeHTTP(response, request)
	var res models.SimpleResponse
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&res)
	if err != nil {
		t.Errorf("Error while decoding req. Error : %v", err)
	}
	assert.Equal(t, getMockResponse(productName), &res, "Response body must match the expected outcome")
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

func getMockResponse(productName string) *models.SimpleResponse {
	res := models.SimpleResponse{}
	res.Message = "new product name: " + productName
	res.Date = ""
	res.Id = 0
	return  &res
}