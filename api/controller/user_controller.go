package controller

import (
	"chipsiBackend/api/middleware"
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strconv"
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

func (uc *UserController) GetMe(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIDKey).(string)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(httpErrors.NewInternalServerError(err)); err != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}

	user, err := uc.UserUsecase.GetByID(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if encodeError := json.NewEncoder(w).Encode(httpErrors.NewNotFoundError(err)); encodeError != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}
	userDTO := domain.ToUserDTO(user)
	if err := json.NewEncoder(w).Encode(userDTO); err != nil {
		uc.Log.Error("can't encode json")
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) PatchMe(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIDKey).(string)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(httpErrors.NewInternalServerError(err)); err != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpErrors.JSONError(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := uc.UserUsecase.Patch(r.Context(), id, body)
	if err != nil {
		if errors.Is(err, httpErrors.BadRequest) {
			httpErrors.JSONError(w, "invalid json", http.StatusBadRequest)
			return

		}
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return

	}

	userDTO := domain.ToUserDTO(user)
	if err := json.NewEncoder(w).Encode(userDTO); err != nil {
		uc.Log.Error("can't encode json")
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(httpErrors.NewBadRequestError("missing id parameter")); err != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(httpErrors.NewBadRequestError("invalid id format")); err != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}

	user, err := uc.UserUsecase.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if encodeError := json.NewEncoder(w).Encode(httpErrors.NewNotFoundError(err)); encodeError != nil {
			uc.Log.Error("can't encode json")
		}
		return
	}

	userDTO := domain.ToUserDTO(user)
	if err := json.NewEncoder(w).Encode(userDTO); err != nil {
		uc.Log.Error("can't encode json")
	}
	w.WriteHeader(http.StatusOK)
}
