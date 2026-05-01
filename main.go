package main

import (
	"fmt"

	"github.com/era0006/game-shop-backend/database"
	"github.com/era0006/game-shop-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	handlers.SetDB(database.DB)

	r := gin.Default()

	r.GET("/games", handlers.GetGames)
	r.GET("/games/:id", handlers.GetGameByID)
	r.GET("/developers", handlers.GetDevelopers)
	r.GET("/genres", handlers.GetGenres)

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.POST("/games", handlers.CreateGame)
		protected.PUT("/games/:id", handlers.UpdateGame)
		protected.DELETE("/games/:id", handlers.DeleteGame)
		protected.POST("/developers", handlers.CreateDeveloper)
		protected.POST("/genres", handlers.CreateGenre)
	}

	r.POST("/notify", handlers.SendNotification)
	r.GET("/notifications", handlers.GetNotifications)

	fmt.Println("🚀 Server starting on http://localhost:8080")
	r.Run(":8080")
}
