package server

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		fmt.Printf("Middleware: %s %s %s\n", r.Method, r.RequestURI, r.Proto)
		next.ServeHTTP(w, r)
	})
}
