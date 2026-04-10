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
	// Подключаемся к БД
	database.Connect()
	handlers.SetDB(database.DB)

	// Автоматическая миграция (создание таблиц)
	err := database.DB.AutoMigrate(
		&models.Game{},
		&models.Developer{},
		&models.Genre{},
		&models.User{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("✅ Database migrated successfully!")

	r := gin.Default()

	// Публичные маршруты (не требуют токен)
	r.GET("/games", handlers.GetGames)
	r.GET("/games/:id", handlers.GetGameByID)
	r.GET("/developers", handlers.GetDevelopers)
	r.GET("/genres", handlers.GetGenres)

	// Маршруты для авторизации
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Защищённые маршруты (требуют токен)
	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.POST("/games", handlers.CreateGame)
		protected.PUT("/games/:id", handlers.UpdateGame)
		protected.DELETE("/games/:id", handlers.DeleteGame)
		protected.POST("/developers", handlers.CreateDeveloper)
		protected.POST("/genres", handlers.CreateGenre)
	}

	fmt.Println("🚀 Server starting on http://localhost:8080")
	r.Run(":8080")
}
