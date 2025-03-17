package middleware

import (
	"context"
	"net/http"
	"observer-go/src/auth"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		tkParts := strings.Split(authorization, " ")
		if len(tkParts) != 2 || tkParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claim, err := auth.NewJWTAuth().VerifyJWT(tkParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", claim["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
