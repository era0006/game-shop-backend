package models

type Game struct {
	ID       int
	Title    string
	Genre    string
	Platform string
	Price    float64
	InStock  bool
}

func NewGame(title, genre, platform string, price float64) *Game {
	return &Game{
		Title:    title,
		Genre:    genre,
		Platform: platform,
		Price:    price,
		InStock:  true,
	}
}
