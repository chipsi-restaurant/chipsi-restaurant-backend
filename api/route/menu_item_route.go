package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewMenuItemRouter(menuItemUsecase domain.MenuItemUsecase, handler func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	mc := controller.MenuItemController{MenuItemUsecase: menuItemUsecase}

	r.Get("/{id}", mc.GetByID)
	r.Get("/", mc.GetAll)
	r.With(handler).Post("/", mc.Create)
	r.With(handler).Delete("/{id}", mc.Delete)

	return r
}
