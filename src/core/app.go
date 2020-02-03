package core

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/microservice_template/src/models"
)

type App struct {
}

func New() *App {
	return &App{}
}

type serviceHandler struct {
	next http.Handler
}

func newServiceHandler(next http.Handler) http.Handler {
	return serviceHandler{next: next}
}

func (h serviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func newServiceSubHandler() http.Handler {
	return serviceSubHandler{}
}

func (h serviceSubHandler) ServeHTTP (w http.ResponseWriter, r *http.Request)  {
	product := r.Context().Value("product").(string)
	response := models.SimpleResponse{Message: "new product name: " + product}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func (a *App) Start() {
	port := 8080
	handler := newServiceHandler(newServiceSubHandler())
	http.Handle("/product", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
