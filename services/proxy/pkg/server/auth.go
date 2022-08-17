package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) auth(c *gin.Context) {
	s.redirect(c, "/auth", os.Getenv("PORT_AUTH"))
}
