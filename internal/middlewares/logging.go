package middlewares

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		log.Printf("%s %s %s", r.Method, r.URL, r.Proto)
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
