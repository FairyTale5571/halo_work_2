package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitHandlers() *gin.Engine {
	router := gin.New()

	router.GET("/user/profile", h.profile)
	router.GET("/microservice/name", h.microserviceName)
	return router
}
