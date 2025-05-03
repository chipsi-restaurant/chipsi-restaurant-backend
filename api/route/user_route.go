package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func NewUserRouter(userUsecase domain.UserUsecase, log *slog.Logger, adminHandler func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()
	uc := controller.UserController{
		UserUsecase: userUsecase,
		Log:         log,
	}
	r.With(adminHandler).Get("/{id}", uc.GetUserByID)
	r.Get("/", uc.GetUserByEmail)
	r.Get("/me", uc.GetMe)
	r.Patch("/me", uc.PatchMe)

	return r
}
