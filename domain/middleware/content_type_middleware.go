package middleware

import (
	"log"
	"net/http"
	"strings"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.RequestURI, "docs"){
			log.Println("Attaching header")
			log.Println(r.RequestURI)
			w.Header().Add("Content-type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}
