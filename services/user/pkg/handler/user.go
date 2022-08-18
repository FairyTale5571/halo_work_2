package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) microserviceName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "user-microservice",
	})
}

func (h *Handler) profile(c *gin.Context) {
	username := c.Request.Header.Get("Username")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"dob":      "01/01/1990",
		"age":      "30",
		"phone":    "1234567890",
	})
}
