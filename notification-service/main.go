package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Notification struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

var notifications = []Notification{}

func main() {
	r := gin.Default()

	r.POST("/notify", func(c *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		notif := Notification{
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
			Message:   req.Message,
			CreatedAt: time.Now(),
		}
		notifications = append(notifications, notif)

		c.JSON(200, gin.H{
			"status":  "sent",
			"message": req.Message,
			"id":      notif.ID,
		})
	})

	r.GET("/notifications", func(c *gin.Context) {
		c.JSON(200, notifications)
	})

	log.Println("📧 Notification Service running on :8081")
	r.Run(":8081")
}
