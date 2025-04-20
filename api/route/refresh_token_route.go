package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func NewRefreshTokenRouter(refreshTokenUsecase domain.RefreshTokenUsecase, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()

	rtc := controller.RefreshTokenController{
		RefreshTokenUsecase: refreshTokenUsecase,
		Cfg:                 cfg,
		Log:                 log,
	}

	r.Post("/", rtc.RefreshToken)

	return r

}
