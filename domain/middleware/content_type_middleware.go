package middleware

import (
	"net/http"
	"strings"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.RequestURI, "docs") && !strings.Contains(r.RequestURI, "image"){
			w.Header().Add("Content-type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}
