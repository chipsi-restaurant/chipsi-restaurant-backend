package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func NewSignupRouter(signupUsecase domain.SignupUsecase, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()
	sc := controller.SignupController{
		SignupUseCase: signupUsecase,
		Log:           log,
		Cfg:           cfg,
	}

	r.Post("/", sc.Signup)

	return r
}
