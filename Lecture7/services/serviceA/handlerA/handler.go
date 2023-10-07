package handlerA

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServiceAHandler struct{}

func NewServiceAHandler() *ServiceAHandler {
	return &ServiceAHandler{}
}

func (h *ServiceAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.welcome(w, r)
	case "/products":
		h.getProducts(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *ServiceAHandler) welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to serviceA!")
}

func (h *ServiceAHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	products := []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}{
		{ID: 1, Name: "Lightning McQueen", Price: 4500},
		{ID: 2, Name: "Barbie", Price: 15000},
		{ID: 3, Name: "Gun", Price: 2500},
	}

	json.NewEncoder(w).Encode(products)
}
