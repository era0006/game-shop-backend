package models

type Game struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	DeveloperID int     `json:"developer_id"`
	GenreID     int     `json:"genre_id"`
	Price       float64 `json:"price"`
	InStock     bool    `json:"in_stock"`
}
