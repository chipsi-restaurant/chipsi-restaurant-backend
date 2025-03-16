package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/bootstrap"
	"chipsiBackend/repository"
	"chipsiBackend/usecase"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

func NewRefreshTokenRouter(db *gorm.DB, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()

	ur := repository.NewUserRepository(db)

	rtu := usecase.NewRefreshTokenUsecase(ur, cfg, 5*time.Second)

	rtc := controller.RefreshTokenController{
		RefreshTokenUsecase: rtu,
		Cfg:                 cfg,
		Log:                 log,
	}

	r.Post("/", rtc.RefreshToken)

	return r

}
