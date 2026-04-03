package handlers

import (
	"net/http"
	"strconv"

	"github.com/era0006/game-shop-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetGames(c *gin.Context) {
	var games []models.Game

	query := db.Model(&models.Game{})

	devFilter := c.Query("developer")
	if devFilter != "" {
		query = query.Joins("JOIN developers ON developers.id = games.developer_id").
			Where("developers.name ILIKE ?", "%"+devFilter+"%")
	}

	genreFilter := c.Query("genre")
	if genreFilter != "" {
		query = query.Joins("JOIN genres ON genres.id = games.genre_id").
			Where("genres.name ILIKE ?", "%"+genreFilter+"%")
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query.Offset(offset).Limit(pageSize).Find(&games)
	c.JSON(http.StatusOK, games)
}

func CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if game.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if game.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be positive"})
		return
	}

	result := db.Create(&game)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}
	c.JSON(http.StatusCreated, game)
}

func GetGameByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var game models.Game
	result := db.First(&game, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func UpdateGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var game models.Game
	if result := db.First(&game, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	var updated models.Game
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if updated.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if updated.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be positive"})
		return
	}

	db.Model(&game).Updates(updated)
	c.JSON(http.StatusOK, game)
}

func DeleteGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := db.Delete(&models.Game{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
