package controller

import (
	"chipsiBackend/api/middleware"
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type GiftCertificateController struct {
	GiftCertificateUsecase domain.GiftCertificateUsecase
}

func (c *GiftCertificateController) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIDKey).(string)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "can't parse userId", http.StatusInternalServerError)
		return
	}

	var request domain.GiftCertificateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httpErrors.JSONError(w, "can't decode request body", http.StatusBadRequest)
		return
	}

	request.SenderID = uint(id)

	response, err := c.GiftCertificateUsecase.Create(r.Context(), &request)

	if err != nil {
		if errors.Is(err, httpErrors.EmailNotExistsError) {
			httpErrors.JSONError(w, "email doesn't exists", http.StatusBadRequest)
			return
		}
		httpErrors.JSONError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		httpErrors.JSONError(w, "can't encode json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (c *GiftCertificateController) GetMine(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		httpErrors.JSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "can't parse userId", http.StatusInternalServerError)
		return
	}

	certificates, err := c.GiftCertificateUsecase.GetBySenderID(r.Context(), id)
	if err != nil {
		httpErrors.JSONError(w, "failed to fetch certificates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(certificates); err != nil {
		httpErrors.JSONError(w, "can't encode response", http.StatusInternalServerError)
		return
	}
}

func (c *GiftCertificateController) GetByPromoCode(w http.ResponseWriter, r *http.Request) {
	promoCode := r.URL.Query().Get("code")
	if promoCode == "" {
		httpErrors.JSONError(w, "promo code is required", http.StatusBadRequest)
		return
	}

	userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		httpErrors.JSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	id, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "can't parse userId", http.StatusInternalServerError)
		return
	}

	response, err := c.GiftCertificateUsecase.GetByCode(r.Context(), promoCode, id)
	if err != nil {
		if errors.Is(err, httpErrors.NotFound) {
			httpErrors.JSONError(w, "gift certificate not found", http.StatusNotFound)
			return
		} else if errors.Is(err, httpErrors.BadRequest) {
			httpErrors.JSONError(w, "gift certificate is used", http.StatusBadRequest)
			return
		}
		httpErrors.JSONError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		httpErrors.JSONError(w, "can't encode response", http.StatusInternalServerError)
		return
	}
}
