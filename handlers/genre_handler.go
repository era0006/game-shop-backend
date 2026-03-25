package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/era0006/game-shop-backend/models"
)

var genres = []models.Genre{
	{ID: 1, Name: "RPG"},
	{ID: 2, Name: "Racing"},
}
var nextGenreID = 3

func GetGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre
	json.NewDecoder(r.Body).Decode(&genre)
	genre.ID = nextGenreID
	nextGenreID++
	genres = append(genres, genre)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(genre)
}
