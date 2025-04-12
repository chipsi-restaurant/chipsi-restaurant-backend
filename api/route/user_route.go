package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func NewUserRouter(userUsecase domain.UserUsecase, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()
	uc := controller.UserController{
		UserUsecase: userUsecase,
		Log:         log,
	}

	r.Get("/", uc.GetUserByEmail)
	r.Get("/me", uc.GetMe)
	r.Patch("/me", uc.PatchMe)

	return r
}
