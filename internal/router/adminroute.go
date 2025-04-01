package router

import "github.com/Sathish-30/authentication-and-authorization-golang/internal/handler"

func RegisterAdminRoute() {
	router.HandleFunc("GET /admins", handler.AdminsHandler)
}
