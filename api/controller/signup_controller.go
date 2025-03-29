package controller

import (
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type SignupController struct {
	SignupUseCase domain.SignupUsecase
	Log           *slog.Logger
	Cfg           *bootstrap.Config
}

func (sc *SignupController) Signup(w http.ResponseWriter, r *http.Request) {

	var request domain.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	user, err := sc.SignupUseCase.Create(r.Context(), request)
	if err != nil {
		if errors.Is(err, httpErrors.ExistsEmailError) {
			restError := httpErrors.NewRestError(http.StatusConflict, err.Error(), err)
			w.WriteHeader(http.StatusConflict)
			if err := json.NewEncoder(w).Encode(restError); err != nil {
				sc.Log.Error("can't encode json")
			}
		} else if err := json.NewEncoder(w).Encode(httpErrors.NewInternalServerError(err)); err != nil {
			sc.Log.Error("can't encode json")
		}
		return
	}

	accessToken, err := sc.SignupUseCase.CreateAccessToken(user, sc.Cfg.App.JwtSecretKey, sc.Cfg.App.AccessTokenExpires)

	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := sc.SignupUseCase.CreateRefreshToken(user, sc.Cfg.App.JwtSecretKey, sc.Cfg.App.RefreshTokenExpires)

	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		sc.Log.Error("can't encode json")
	}

	w.WriteHeader(http.StatusOK)
	return
}
