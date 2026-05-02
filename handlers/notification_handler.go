package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var restyClient = resty.New()

func SendNotification(c *gin.Context) {
	var req struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// ВНИМАНИЕ: в Docker используем имя сервиса "notification", а не localhost
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"message": req.Message}).
		Post("http://notification:8081/notify")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notification service unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notification_response": string(resp.Body()),
	})
}

func GetNotifications(c *gin.Context) {
	resp, err := restyClient.R().Get("http://notification:8081/notifications")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notification service unavailable"})
		return
	}

	c.Data(http.StatusOK, "application/json", resp.Body())
}
