package server

import (
	"fmt"
	"net/http"

	"github.com/Sathish-30/authentication-and-authorization-golang/pkg/log"
)

type Server struct {
	port   string
	router *http.ServeMux
}

func (s *Server) SetPort(port int) {
	s.port = fmt.Sprintf(":%d", port)
}

func (s *Server) SetRouter(router *http.ServeMux) {
	s.router = router
}

func (s *Server) GetPort() string {
	return s.port
}

func (s *Server) GetRouter() *http.ServeMux {
	return s.router
}

func NewServer(port int, router *http.ServeMux) *Server {
	return &Server{
		port:   fmt.Sprintf(":%d", port),
		router: router,
	}
}

func (s *Server) Start() {
	if err := http.ListenAndServe(s.port, s.router); err != nil {
		log.Logger.Fatal("Failed to start server: " + err.Error())
	}
}
