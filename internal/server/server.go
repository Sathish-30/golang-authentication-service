package server

import (
	"fmt"
	"net/http"

	"github.com/Sathish-30/authentication-and-authorization-golang/pkg/log"
)

type Server struct {
	port    string
	handler http.Handler
}

func (s *Server) SetPort(port int) {
	s.port = fmt.Sprintf(":%d", port)
}

func (s *Server) GetPort() string {
	return s.port
}

func (s *Server) SetHandler(handler http.Handler) {
	s.handler = handler
}
func (s *Server) GetHandler() http.Handler {
	return s.handler
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		port:    fmt.Sprintf(":%d", port),
		handler: handler,
	}
}

func (s *Server) Start() {
	if err := http.ListenAndServe(s.port, s.handler); err != nil {
		log.Logger.Fatal("Failed to start server: " + err.Error())
	}
}
