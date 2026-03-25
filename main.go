package main

import (
	"fmt"
	"net/http"

	"github.com/era0006/game-shop-backend/handlers"
)

func main() {
	fmt.Println("🎮 Game Shop API starting on http://localhost:8080")

	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetGames(w, r)
		case http.MethodPost:
			handlers.CreateGame(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/games/", handlers.GetGameByID)

	http.HandleFunc("/developers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetDevelopers(w, r)
		case http.MethodPost:
			handlers.CreateDeveloper(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetGenres(w, r)
		case http.MethodPost:
			handlers.CreateGenre(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
