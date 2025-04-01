package main

import (
	"net/http"

	"github.com/Sathish-30/authentication-and-authorization-golang/internal/middleware"
	"github.com/Sathish-30/authentication-and-authorization-golang/internal/router"
	"github.com/Sathish-30/authentication-and-authorization-golang/internal/server"
	"github.com/Sathish-30/authentication-and-authorization-golang/pkg/log"
)

func main() {
	log.Logger.Info("Initializing router...")
	router.SetRouter(http.NewServeMux())
	handler := middleware.LoggingMiddleware(
		middleware.ContentTypeMiddleware(router.GetRouter()),
	)
	server := server.NewServer(8000, handler)
	router.RegisterRoutes()
	log.Logger.Infof("Starting server on port %s", server.GetPort())
	server.Start()
}
