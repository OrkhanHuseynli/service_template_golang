package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/microservice_template/src/models"
)

type ServiceHandler struct {
	next http.Handler
}

func NewServiceHandler(next http.Handler) http.Handler {
	return ServiceHandler{next: next}
}

func (h ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.SimpleRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.Product == "" {
		http.Error(w, "Required body fields are empty", http.StatusUnprocessableEntity)
		return
	}

	c := context.WithValue(r.Context(), "product", req.Product)
	r = r.WithContext(c)
	h.next.ServeHTTP(w, r)
}

type serviceSubHandler struct {
}

func NewServiceSubHandler() http.Handler {
	return serviceSubHandler{}
}

func (h serviceSubHandler) ServeHTTP (w http.ResponseWriter, r *http.Request)  {
	product := r.Context().Value("product").(string)
	response := models.SimpleResponse{Message: "new product name: " + product}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
