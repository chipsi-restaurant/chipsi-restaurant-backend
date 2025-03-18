package middleware

import (
	"chipsiBackend/internal/tokenutil"
	"chipsiBackend/pkg/httpErrors"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "x-user-id"

func JwtAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := tokenutil.IsAuthorized(authToken, secret)
				if authorized {
					userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
					if err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						err := json.NewEncoder(w).Encode(httpErrors.NewRestError(http.StatusUnauthorized, err.Error(), err))
						if err != nil {
							http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
							return
						}
						return
					}
					ctx := context.WithValue(r.Context(), UserIDKey, userID)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				w.WriteHeader(http.StatusUnauthorized)
				err = json.NewEncoder(w).Encode(httpErrors.NewRestError(http.StatusUnauthorized, err.Error(), err))
				if err != nil {
					http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
					return
				}
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		})
	}
}
