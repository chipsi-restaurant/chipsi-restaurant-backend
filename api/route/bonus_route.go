package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
)

func NewBonusRouter(bonusUsecase domain.BonusUsecase) chi.Router {
	r := chi.NewRouter()
	bc := controller.BonusController{
		BonusUsecase: bonusUsecase,
	}
	r.Post("/", bc.Create)

	return r
}
