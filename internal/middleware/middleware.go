package middleware

import (
	"net/http"
	"time"

	"github.com/Sathish-30/authentication-and-authorization-golang/pkg/log"
)

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			duration := time.Since(startTime)
			log.Logger.Infof("Request %s %s took %s", r.Method, r.URL.Path, duration)
		}()
		handler.ServeHTTP(w, r)
	})
}

func ContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
