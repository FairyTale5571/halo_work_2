package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) auth(c *gin.Context) {
	username := c.Request.Header.Get("Username")
	u := h.services.Authorize(username)
	if !u {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid username",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
