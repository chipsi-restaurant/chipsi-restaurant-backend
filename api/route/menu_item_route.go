package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/bootstrap"
	"chipsiBackend/repository"
	"chipsiBackend/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewMenuItemRouter(client *minio.Client, db *gorm.DB,
	cfg *bootstrap.Config, handler func(http.Handler) http.Handler) chi.Router {

	r := chi.NewRouter()

	s3Usecase := usecase.NewS3Usecase(client, cfg)
	categoryRepository := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository, time.Second*10)

	menuItemRepository := repository.NewMenuItemRepository(db)
	menuItemUsecase := usecase.NewMenuItemUsecase(s3Usecase, categoryUsecase, menuItemRepository, time.Second*10)

	mc := controller.MenuItemController{MenuItemUsecase: menuItemUsecase}

	r.Get("/{id}", mc.GetByID)
	r.Get("/", mc.GetAll)
	r.With(handler).Post("/", mc.Create)
	r.With(handler).Delete("/{id}", mc.Delete)

	return r
}
