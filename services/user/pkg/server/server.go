package server

import (
	"context"
	"net/http"
	"os"
)

type Server struct {
	http *http.Server
}

func InitServer() *Server {
	return &Server{
		http: &http.Server{
			Addr: ":" + os.Getenv("PORT_USER"),
		},
	}
}

func (s *Server) SetHandler(handler http.Handler) {
	s.http.Handler = handler
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
