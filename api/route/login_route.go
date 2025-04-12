package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func NewLoginRouter(loginUsecase domain.LoginUsecase, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()

	lc := controller.LoginController{
		LoginUsecase: loginUsecase,
		Cfg:          cfg,
		Log:          log,
	}

	r.Post("/", lc.Login)

	return r

}
