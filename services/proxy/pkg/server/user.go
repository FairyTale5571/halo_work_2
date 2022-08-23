package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) microServiceName(c *gin.Context) {
	s.redirect(c, "user", os.Getenv("PORT_USER"), "/microservice/name")
}

func (s *Server) checkAuth(username string) int {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://auth:%s/auth", os.Getenv("PORT_AUTH")), http.NoBody)
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

func (s *Server) userProfile(c *gin.Context) {
	if s.checkAuth(c.Request.Header.Get("Username")) != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	s.redirect(c, "user", os.Getenv("PORT_USER"), "/user/profile")
}
