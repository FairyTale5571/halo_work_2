package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) microserviceName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "user-microservice",
	})
}

func (h *Handler) profile(c *gin.Context) {
	username := c.Request.Header.Get("Username")
	if h.checkAuth(username) != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid username",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"dob":      "01/01/1990",
		"age":      "30",
		"phone":    "1234567890",
	})
}

func (h *Handler) checkAuth(username string) int {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s:%s/auth", os.Getenv("URL"), os.Getenv("PORT_AUTH")), nil)
	if err != nil {
		return http.StatusInternalServerError
	}
	req.Header.Add("Username", username)
	res, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError
	}
	if res.StatusCode != http.StatusOK {
		return 401
	}
	return http.StatusOK
}
