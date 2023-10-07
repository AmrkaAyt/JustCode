package handlerB

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServiceBHandler struct{}

func NewServiceBHandler() *ServiceBHandler {
	return &ServiceBHandler{}
}

func (h *ServiceBHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.welcome(w, r)
	case "/products":
		h.getProducts(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *ServiceBHandler) welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Service B!")
}

func (h *ServiceBHandler) getProducts(w http.ResponseWriter, r *http.Request) {

	products := []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}{
		{ID: 4, Name: "Root", Price: 25000},
		{ID: 5, Name: "Star Realms", Price: 9500},
		{ID: 6, Name: "Orleans", Price: 20000},
	}

	json.NewEncoder(w).Encode(products)
}
