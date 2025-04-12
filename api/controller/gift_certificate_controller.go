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
