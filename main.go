package main

import (
	"fmt"
	"log"

	"github.com/era0006/game-shop-backend/database"
	"github.com/era0006/game-shop-backend/handlers"
	"github.com/era0006/game-shop-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	handlers.SetDB(database.DB)

	err := database.DB.AutoMigrate(&models.Game{}, &models.Developer{}, &models.Genre{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Database migrated!")

	r := gin.Default()

	r.GET("/games", handlers.GetGames)
	r.POST("/games", handlers.CreateGame)
	r.GET("/games/:id", handlers.GetGameByID)
	r.PUT("/games/:id", handlers.UpdateGame)
	r.DELETE("/games/:id", handlers.DeleteGame)

	r.GET("/developers", handlers.GetDevelopers)
	r.POST("/developers", handlers.CreateDeveloper)

	r.GET("/genres", handlers.GetGenres)
	r.POST("/genres", handlers.CreateGenre)

	r.Run(":8080")
}
