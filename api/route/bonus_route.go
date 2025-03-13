package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/repository"
	"chipsiBackend/usecase"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"time"
)

func NewBonusRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	br := repository.NewBonusRepository(db)
	bc := controller.BonusController{
		BonusUsecase: usecase.NewBonusUsecase(br, time.Second*5),
	}
	r.Post("/bonuses", bc.Create)

	return r
}
