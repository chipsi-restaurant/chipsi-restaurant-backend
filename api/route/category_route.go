package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewCategoryRouter(categoryUsecase domain.CategoryUsecase, handler func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	cc := controller.CategoryController{
		CategoryUsecase: categoryUsecase,
	}
	r.Get("/", cc.GetAll)
	r.Get("/{id}", cc.GetByID)
	r.With(handler).Post("/", cc.Create)
	r.With(handler).Delete("/{id}", cc.Delete)

	return r

}
