package middleware

import (
	"chipsiBackend/domain"
	"net/http"
	"strconv"
)

func IsAdmin(userUsecase domain.UserUsecase) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value(UserIDKey).(string)
			id, err := strconv.ParseInt(userId, 10, 64)
			if err != nil {
				http.Error(w, `{"error": "access denied"}`, http.StatusForbidden)
				return
			}
			user, err := userUsecase.GetByID(r.Context(), id)
			if user.Admin == nil || len(user.Admin) == 0 {
				http.Error(w, `{"error": "access denied"}`, http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
