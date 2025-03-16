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

func NewLoginRouter(db *gorm.DB, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()

	ur := repository.NewUserRepository(db)

	lu := usecase.NewLoginUsecase(ur, 5*time.Second, cfg)

	lc := controller.LoginController{
		LoginUsecase: lu,
		Cfg:          cfg,
		Log:          log,
	}

	r.Post("/", lc.Login)

	return r

}
