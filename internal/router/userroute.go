package router

import "github.com/Sathish-30/authentication-and-authorization-golang/internal/handler"

func registerUserRoute() {
	router.HandleFunc("GET /users", handler.UsersHandler)
}
