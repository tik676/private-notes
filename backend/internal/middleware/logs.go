package middleware

import (
	"context"
	"net/http"
	"os"
	"private-notes/api/authorization"
	"strings"
)

func MiddlewareCheckJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует токен", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		jwtMaker := authorization.NewJWTMaker(os.Getenv("JWT_SECRET"))

		user, err := jwtMaker.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Невалидный токен", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
