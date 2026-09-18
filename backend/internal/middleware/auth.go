package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userId"


func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Autorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Unauthorized invalid header format", http.StatusUnauthorized)
			}
			tokenString := parts[1]

			token, err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{}, error) {
				if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(jwtSecret), nil

			})
			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized: invalid or expired token", http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized: invalid token claims", http.StatusUnauthorized)
			}
			userId, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, "Unauthorized: invalid user ID in token", http.StatusUnauthorized) 
				return 
			}
			ctx := context.WithValue(r.Context(), UserIDKey, userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		}) 
	}
}


func GetUserIdFromContext(ctx context.Context) (string,bool) {
	userId, ok := ctx.Value(UserIDKey).(string)
	return userId, ok
}