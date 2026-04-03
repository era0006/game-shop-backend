package handlers

import (
	"net/http"

	"github.com/era0006/game-shop-backend/models"
	"github.com/gin-gonic/gin"
)

func GetDevelopers(c *gin.Context) {
	var developers []models.Developer
	db.Find(&developers)
	c.JSON(http.StatusOK, developers)
}

func CreateDeveloper(c *gin.Context) {
	var dev models.Developer
	if err := c.ShouldBindJSON(&dev); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if dev.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	result := db.Create(&dev)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create developer"})
		return
	}
	c.JSON(http.StatusCreated, dev)
}
