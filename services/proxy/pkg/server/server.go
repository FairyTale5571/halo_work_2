package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/fairytale5571/halo_work_2/services/proxy/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	logger *logger.Wrapper
	app    *gin.Engine
	client *http.Client
}

func New() *Server {
	return &Server{
		logger: logger.New("auth_server"),
		app:    gin.Default(),
		client: &http.Client{},
	}
}

func (s *Server) Run() error {
	s.logger.Infof("server run on port: %s", os.Getenv("PORT_PROXY"))
	s.mainRouter()
	return s.app.Run(":" + os.Getenv("PORT_PROXY"))
}

func (s *Server) mainRouter() {
	s.app.GET("/auth", s.auth)
	s.app.GET("/microservice/name", s.microServiceName)
	s.app.GET("/user/profile", s.userProfile)
}

func (s *Server) redirect(c *gin.Context, method, port string) {
	u, err := url.Parse(fmt.Sprintf("%s:%s%s", os.Getenv("URL"), port, method))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = u.Host
		req.URL.Scheme = u.Scheme
		req.URL.Host = u.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
