package usecase

import (
	"chipsiBackend/domain"
	"context"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type menuItemUsecase struct {
	s3Usecase          S3Usecase
	categoryUsecase    domain.CategoryUsecase
	menuItemRepository domain.MenuItemRepository
	contextTimeout     time.Duration
}

func NewMenuItemUsecase(s3Usecase S3Usecase, categoryUsecase domain.CategoryUsecase, menuItemRepository domain.MenuItemRepository, timeout time.Duration) domain.MenuItemUsecase {
	return &menuItemUsecase{
		s3Usecase:          s3Usecase,
		categoryUsecase:    categoryUsecase,
		menuItemRepository: menuItemRepository,
		contextTimeout:     timeout,
	}
}

func (mu menuItemUsecase) Create(ctx context.Context, menuItem *domain.MenuItem, image multipart.File, handler *multipart.FileHeader) (*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	contentType := handler.Header.Get("Content-Type")
	size := handler.Size
	imageUrl, err := mu.s3Usecase.UploadImage(ctx, image, uuid.NewString(), contentType, size)
	if err != nil {
		return nil, err
	}

	menuItem.ImageURL = imageUrl
	newMenuItem, err := mu.menuItemRepository.Create(ctx, menuItem)
	if err != nil {
		return nil, err
	}
	return newMenuItem, nil
}

func (mu menuItemUsecase) GetByID(ctx context.Context, id int64) (*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepository.GetByID(ctx, id)
}

func (mu menuItemUsecase) GetAllByCategory(ctx context.Context, categoryId int64) ([]*domain.MenuItem, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	_, err := mu.categoryUsecase.GetByID(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return mu.menuItemRepository.GetAllByCategory(ctx, categoryId)
}

func (mu menuItemUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	return mu.menuItemRepository.Delete(ctx, id)
}
