package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) microServiceName(c *gin.Context) {
	s.redirect(c, "/microservice/name", os.Getenv("PORT_USER"))
}

func (s *Server) userProfile(c *gin.Context) {
	s.redirect(c, "/user/profile", os.Getenv("PORT_USER"))
}
