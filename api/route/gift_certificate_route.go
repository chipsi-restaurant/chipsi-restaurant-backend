package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
)

func NewGiftCertificateRouter(giftCertificateUsecase domain.GiftCertificateUsecase) chi.Router {
	r := chi.NewRouter()
	gc := controller.GiftCertificateController{
		GiftCertificateUsecase: giftCertificateUsecase,
	}

	r.Get("/mine", gc.GetMine)
	r.Get("/", gc.GetByPromoCode)
	r.Post("/", gc.Create)

	return r
}
