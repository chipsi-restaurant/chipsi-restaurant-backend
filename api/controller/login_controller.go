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

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Cfg          *bootstrap.Config
	Log          *slog.Logger
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	request := domain.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
	}

	response, err := lc.LoginUsecase.Login(r.Context(), request)

	if err != nil {
		if errors.Is(err, httpErrors.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			restError := httpErrors.NewRestError(http.StatusNotFound, err.Error(), err)
			if jsonError := json.NewEncoder(w).Encode(restError); jsonError != nil {
				lc.Log.Error("can't encode json")
			}
		} else if errors.Is(err, httpErrors.Unauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			restError := httpErrors.NewRestError(http.StatusUnauthorized, err.Error(), err)
			if jsonError := json.NewEncoder(w).Encode(restError); jsonError != nil {
				lc.Log.Error("can't encode json")
			}
		} else {
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(*response); err != nil {
		lc.Log.Error("can't encode json")
	}
	w.WriteHeader(http.StatusOK)
	return
}
