package controller

import (
	"chipsiBackend/api/middleware"
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ReservationController struct {
	ReservationUsecase domain.ReservationUsecase
}

func (c *ReservationController) CreateReservation(w http.ResponseWriter, r *http.Request) {

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

	var reservationRequest domain.ReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&reservationRequest); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	reservationRequest.UserID = uint(id)

	response, err := c.ReservationUsecase.Create(r.Context(), &reservationRequest)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ReservationController) GetAllReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := c.ReservationUsecase.GetAll(r.Context())
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(reservations); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *ReservationController) GetReservationsByUserID(w http.ResponseWriter, r *http.Request) {
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

	reservations, err := c.ReservationUsecase.GetByUserID(r.Context(), uint(id))
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(reservations); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rc *ReservationController) UpdateReservationStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		httpErrors.JSONError(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "invalid id format", http.StatusBadRequest)
		return
	}

	var req domain.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpErrors.JSONError(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Status == "" {
		httpErrors.JSONError(w, "status is required", http.StatusBadRequest)
		return
	}

	fields := map[string]interface{}{"status": req.Status}
	if req.Status == "canceled" || req.Status == "rejected" {
		fields["comment"] = req.Comment
	}

	err = rc.ReservationUsecase.UpdateStatus(r.Context(), id, fields)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Reservation updated"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
