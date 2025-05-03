package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func NewEventRouter(eventUsecase domain.EventUsecase, adminHandler func(http.Handler) http.Handler, log *slog.Logger) chi.Router {
	r := chi.NewRouter()
	ec := controller.EventController{
		EventUsecase: eventUsecase,
		Log:          log,
	}

	r.Post("/", ec.CreateEvent)
	r.Get("/me", ec.GetEventsByUserID)

	r.With(adminHandler).Get("/", ec.GetAllEvents)
	r.With(adminHandler).Patch("/{id}/status", ec.UpdateEventStatus)

	return r
}
