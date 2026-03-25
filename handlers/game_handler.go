package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/era0006/game-shop-backend/models"
)

var games = []models.Game{
	{ID: 1, Title: "The Witcher 3", DeveloperID: 1, GenreID: 1, Price: 39.99, InStock: true},
	{ID: 2, Title: "Mario Kart", DeveloperID: 2, GenreID: 2, Price: 44.99, InStock: true},
}
var nextGameID = 3

func GetGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	json.NewDecoder(r.Body).Decode(&game)
	game.ID = nextGameID
	nextGameID++
	games = append(games, game)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func GetGameByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/games/")
	id, _ := strconv.Atoi(idStr)
	for _, game := range games {
		if game.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(game)
			return
		}
	}
	http.Error(w, "Game not found", http.StatusNotFound)
}
