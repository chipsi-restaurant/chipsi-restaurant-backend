package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/repository"
	"chipsiBackend/usecase"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewCategoryRouter(db *gorm.DB, handler func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	cr := repository.NewCategoryRepository(db)

	cu := usecase.NewCategoryUsecase(cr, 5*time.Second)

	cc := controller.CategoryController{
		CategoryUsecase: cu,
	}
	r.Get("/", cc.GetAll)
	r.Get("/{id}", cc.GetByID)
	r.With(handler).Post("/", cc.Create)
	r.With(handler).Delete("/{id}", cc.Delete)

	return r

}
