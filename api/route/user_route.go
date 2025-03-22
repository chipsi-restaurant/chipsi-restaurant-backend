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

func NewUserRouter(db *gorm.DB, log *slog.Logger, cfg *bootstrap.Config) chi.Router {
	r := chi.NewRouter()
	ur := repository.NewUserRepository(db)
	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, 5*time.Second),
		Log:         log,
	}

	r.Get("/", uc.GetUserByEmail)

	return r
}
