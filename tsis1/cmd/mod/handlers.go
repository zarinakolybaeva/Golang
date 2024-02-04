package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zarinakolybaeva/tsis1/pkg/models"

	"github.com/gorilla/mux"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! made by Zarina!\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is healthy!\n")
}

func PrintAllActors(w http.ResponseWriter, r *http.Request) {
	players := models.GetAllPlayers()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetActorByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	player := models.GetPlayerByName(name)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
