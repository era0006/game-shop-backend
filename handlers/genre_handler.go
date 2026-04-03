package handlers

import (
	"net/http"

	"github.com/era0006/game-shop-backend/models"
	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	var genres []models.Genre
	db.Find(&genres)
	c.JSON(http.StatusOK, genres)
}

func CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if genre.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	result := db.Create(&genre)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create genre"})
		return
	}
	c.JSON(http.StatusCreated, genre)
}
