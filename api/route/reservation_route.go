package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewReservationRouter(reservationUsecase domain.ReservationUsecase, adminHandler func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()
	rc := controller.ReservationController{
		ReservationUsecase: reservationUsecase,
	}

	r.Post("/", rc.CreateReservation)
	r.Get("/me", rc.GetReservationsByUserID)

	r.With(adminHandler).Get("/", rc.GetAllReservations)
	r.With(adminHandler).Patch("/{id}/status", rc.UpdateReservationStatus)

	return r
}
