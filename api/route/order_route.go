package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
)

func NewOrderRouter(orderUsecase domain.OrderUsecase) chi.Router {
	oc := controller.OrderController{
		OrderUsecase: orderUsecase,
	}
	r := chi.NewRouter()

	r.Get("/mine", oc.GetUserOrders)
	r.Get("/status", oc.GetOrderStatuses)
	r.Post("/", oc.CreateOrder)
	return r
}
