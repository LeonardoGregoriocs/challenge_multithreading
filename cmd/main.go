package main

import (
	"net/http"

	"github.com/leonardogregoriocs/challenge_multithreading/core/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SearchZipCodeHandler)

	http.ListenAndServe(":8080", nil)
}
