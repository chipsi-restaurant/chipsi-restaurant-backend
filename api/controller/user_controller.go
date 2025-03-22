package controller

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"log/slog"
	"net/http"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Log         *slog.Logger
}

func (uc *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(httpErrors.NewNotFoundError(httpErrors.NotFound)); err != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}
	user, err := uc.UserUsecase.GetByEmail(r.Context(), email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if encodeError := json.NewEncoder(w).Encode(httpErrors.NewNotFoundError(email)); encodeError != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}

	userDTO := domain.ToUserDTO(user)

	if err := json.NewEncoder(w).Encode(userDTO); err != nil {
		uc.Log.Error("can't encode json")
	}
	w.WriteHeader(http.StatusOK)
	return
}
