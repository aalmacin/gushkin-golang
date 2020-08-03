package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const userIdKey = "userIdKey"

// CurrentUserMiddleware Middleware to get current user from jwt token
func CurrentUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
		tokenString := authHeaderParts[1]

		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			cert, err := getPemCert(token)
			if err != nil {
				return nil, err
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		})

		if err != nil {
			fmt.Println("Failed to parse with claims")
		}

		req := r
		if token != nil {
			userID := token.Claims.(jwt.MapClaims)["sub"]

			ctx := context.WithValue(r.Context(), userIdKey, userID)
			req = r.WithContext(ctx)
		}

		next.ServeHTTP(w, req)
	})
}

// GetCurrentUser Get userID from context
func GetCurrentUser(ctx context.Context) string {
	return ctx.Value(userIdKey).(string)
}
