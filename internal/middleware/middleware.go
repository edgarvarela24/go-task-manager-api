package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get JWT from Auth header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		}

		// Split the Authorization header to get the token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// Verify the token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey := os.Getenv("JWT_SECRET")
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract the user ID from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID := claims["userID"]

		// Set the user ID in the request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		// Call the next handler
		next(w, r)
	}
}
