package models

type Game struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `gorm:"not null" json:"title"`
	DeveloperID uint    `json:"developer_id"`
	GenreID     uint    `json:"genre_id"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	InStock     bool    `json:"in_stock"`
}
