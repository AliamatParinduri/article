package middleware

import (
	"article_app/helper"
	"article_app/modules/auth/usecase"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

func AuthorizeJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		const BEARER_SCHEMA = "Bearer "
		authHeader := r.Header.Get("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := usecase.NewJWTUsecase().Validate(tokenString)

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.ServiceError{Message: "Error verifying JWT token: " + err.Error()})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := strconv.FormatFloat(claims["id"].(float64), 'g', 1, 64)
		userAdmin := strconv.FormatBool(claims["IsAdmin"].(bool))

		r.Header.Set("userId", userId)
		r.Header.Set("userAdmin", userAdmin)
		next.ServeHTTP(w, r)
	})
}
