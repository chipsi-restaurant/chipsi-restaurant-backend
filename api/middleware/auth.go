package middleware

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"net/http"
)

func jwtAuth(secret string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var signupRequest domain.SignupRequest

		if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			restError := httpErrors.NewRestError(http.StatusBadRequest, err.Error(), err)
			if err := json.NewEncoder(w).Encode(restError); err != nil {
				http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
				return
			}
		}

	}
}
