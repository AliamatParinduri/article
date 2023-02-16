package middleware

import (
	"article_app/helper"
	"encoding/json"
	"net/http"
)

func AuthorizeIsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userAdmin := r.Header.Get("userAdmin")

		if userAdmin == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error your account not Authorization"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
