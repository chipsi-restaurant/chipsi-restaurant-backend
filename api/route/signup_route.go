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

func NewSignupRouter(db *gorm.DB, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUseCase: usecase.NewSignupUsecase(ur, 5*time.Second),
		Log:           log,
		Cfg:           cfg,
	}

	r.Post("/", sc.Signup)

	return r
}
