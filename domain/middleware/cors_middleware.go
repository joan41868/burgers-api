package middleware

import (
	"log"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqOrigin := r.Header.Get("Origin")
		// colorful log
		log.Println("\x1b[31m" + r.Method + " \x1b[32m" + r.RequestURI + " \x1b[33m" + reqOrigin)

		// I hope this is ok?
		w.Header().Add("Access-Control-Allow-Origin", reqOrigin) // <- this instead of *
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
