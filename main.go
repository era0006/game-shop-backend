package main

import (
	"fmt"

	"github.com/era0006/game-shop-backend/models"
	"github.com/era0006/game-shop-backend/storage"
)

func main() {
	fmt.Println("Game Shop Backend starting...")

	store := storage.NewStore()

	games := []*models.Game{
		models.NewGame("The Witcher 3", "RPG", "PC", 39.99),
		models.NewGame("Cyberpunk 2077", "RPG", "PC", 49.99),
		models.NewGame("Mario Kart", "Racing", "Nintendo", 44.99),
	}

	for _, g := range games {
		store.Add(g)
	}

	fmt.Println("\nGames in store:")
	for _, g := range store.GetAll() {
		fmt.Printf("- %s | $%.2f | %s\n", g.Title, g.Price, g.Genre)
	}
}
