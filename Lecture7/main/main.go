package main

import (
	"Lecture7/services/serviceA/handlerA"
	"Lecture7/services/serviceB/handlerB"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/serviceA/", http.StripPrefix("/serviceA", handlerA.NewServiceAHandler()))

	mux.Handle("/serviceB/", http.StripPrefix("/serviceB", handlerB.NewServiceBHandler()))

	http.ListenAndServe(":8080", mux)
}
