package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/era0006/game-shop-backend/models"
)

var developers = []models.Developer{
	{ID: 1, Name: "CD Projekt Red"},
	{ID: 2, Name: "Nintendo"},
}
var nextDevID = 3

func GetDevelopers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(developers)
}

func CreateDeveloper(w http.ResponseWriter, r *http.Request) {
	var dev models.Developer
	json.NewDecoder(r.Body).Decode(&dev)
	dev.ID = nextDevID
	nextDevID++
	developers = append(developers, dev)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dev)
}
