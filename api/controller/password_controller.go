package controller

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"net/http"
)

type PasswordController struct {
	PasswordUsecase domain.PasswordResetTokenUsecase
}

func (pc *PasswordController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpErrors.JSONError(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := pc.PasswordUsecase.ForgotPassword(r.Context(), body.Email); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpErrors.JSONError(w, "Password reset email sent", http.StatusOK)
}

func (pc *PasswordController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Token       string `json:"token"`
		NewPassword string `json:"newPassword"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpErrors.JSONError(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := pc.PasswordUsecase.ResetPassword(r.Context(), body.Token, body.NewPassword); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpErrors.JSONError(w, "Password has been reset", http.StatusOK)
}
