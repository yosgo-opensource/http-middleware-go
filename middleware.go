package middleware

import (
	"net/http"
)

func HttpJSONResponse(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rw.Header().Set("Content-Type", "text/html")
		case http.MethodPost:
			rw.Header().Set("Content-Type", "application/json")
		}

		next.ServeHTTP(rw, r)
	})
}

func HttpCORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(rw, r)
	})
}
