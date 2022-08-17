package handler

import (
	"github.com/fairytale5571/halo_work_2/services/auth/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitHandlers() *gin.Engine {
	router := gin.New()
	router.GET("/auth", h.auth)

	return router
}
