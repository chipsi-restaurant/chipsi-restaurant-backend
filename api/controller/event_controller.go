package controller

import (
	"chipsiBackend/api/middleware"
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strconv"
)

type EventController struct {
	EventUsecase domain.EventUsecase
	Log          *slog.Logger
}

func (c *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
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

	var eventRequest domain.EventRequest
	if err := json.NewDecoder(r.Body).Decode(&eventRequest); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	eventRequest.UserID = uint(id)

	response, err := c.EventUsecase.Create(r.Context(), &eventRequest)
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

func (c *EventController) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.EventUsecase.GetAll(r.Context())
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(events); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *EventController) GetEventsByUserID(w http.ResponseWriter, r *http.Request) {
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

	events, err := c.EventUsecase.GetByUserID(r.Context(), uint(id))
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(events); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EventController) UpdateEventStatus(w http.ResponseWriter, r *http.Request) {
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

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpErrors.JSONError(w, "invalid json", http.StatusBadRequest)
		return
	}

	status, ok := body["status"].(string)
	if !ok || status == "" {
		httpErrors.JSONError(w, "missing or invalid status", http.StatusBadRequest)
		return
	}

	fields := map[string]interface{}{
		"status": status,
	}

	if comment, ok := body["comment"].(string); ok && comment != "" {
		fields["comment"] = comment
	}

	if err := c.EventUsecase.UpdateStatus(r.Context(), id, fields); err != nil {
		c.Log.Error("failed to update event status", "error", err)
		httpErrors.JSONError(w, "failed to update event status", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Event status updated"}); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
