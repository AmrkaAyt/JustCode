package main

import (
	"Lecture7/services/serviceA/handlerA"
	"net/http"
)

func main() {
	http.Handle("/", handlerA.NewServiceAHandler())
	http.ListenAndServe(":8081", nil)
}
