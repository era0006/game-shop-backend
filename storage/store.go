package storage

import "github.com/era0006/game-shop-backend/models"

type Store struct {
	games  map[int]*models.Game
	nextID int
}

func NewStore() *Store {
	return &Store{
		games:  make(map[int]*models.Game),
		nextID: 1,
	}
}

func (s *Store) Add(game *models.Game) {
	game.ID = s.nextID
	s.games[s.nextID] = game
	s.nextID++
}

func (s *Store) GetAll() []*models.Game {
	var result []*models.Game
	for _, g := range s.games {
		result = append(result, g)
	}
	return result
}
