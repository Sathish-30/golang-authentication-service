package router

import (
	"net/http"

	"github.com/Sathish-30/authentication-and-authorization-golang/internal/handler"
)

var router *http.ServeMux

func SetRouter(r *http.ServeMux) {
	router = r
}

func GetRouter() *http.ServeMux {
	return router
}

func RegisterHomeRoute() {
	router.HandleFunc("GET /", handler.HomeHandler)
}
