package main

import (
	"net/http"

	"github.com/gorilla/mux"
	
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/actor", WelcomeHandler)
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/actors", PrintAllActors).Methods("GET")
	r.HandleFunc("/actor/{name}", GetActorByName).Methods("GET")
	http.ListenAndServe(":8050", r)
}