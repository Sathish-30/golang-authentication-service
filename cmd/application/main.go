package main

import (
	"net/http"

	"github.com/Sathish-30/authentication-and-authorization-golang/internal/router"
	"github.com/Sathish-30/authentication-and-authorization-golang/internal/server"
	"github.com/Sathish-30/authentication-and-authorization-golang/pkg/log"
)

func main() {
	router.SetRouter(http.NewServeMux())
	appRouter := router.GetRouter()
	server := server.NewServer(8000, appRouter)
	log.Logger.Info("Initializing server...")
	log.Logger.Info("Initializing router...")
	router.RegisterHomeRoute()
	router.RegisterAdminRoute()
	router.RegisterUserRoute()
	log.Logger.Infof("Starting server on port %s", server.GetPort())
	server.Start()
}
