package main

import (
	"Lecture7/services/serviceB/handlerB"
	"net/http"
)

func main() {
	http.Handle("/", handlerB.NewServiceBHandler())
	http.ListenAndServe(":8082", nil)
}
