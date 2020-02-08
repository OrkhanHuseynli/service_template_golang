package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/microservice_template/src/handlers"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Start() {
	port := 8080
	handler := handlers.NewServiceHandler(handlers.NewServiceSubHandler())
	http.Handle("/product", handler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
