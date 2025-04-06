package usecase

import (
	"chipsiBackend/domain"
	"context"
	"time"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}

func (c categoryUsecase) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.Create(ctx, category)
}

func (c categoryUsecase) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.GetByID(ctx, id)
}

func (c categoryUsecase) GetAll(ctx context.Context, includeMenuItems bool) ([]*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categories, err := c.categoryRepository.GetAll(ctx, includeMenuItems)
	if err != nil {
		return nil, err
	}

	if includeMenuItems {
		for _, category := range categories {
			if category.MenuItems == nil {
				category.MenuItems = &[]domain.MenuItem{}
			}
		}

	}

	return categories, nil

}

func (c categoryUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return c.categoryRepository.Delete(ctx, id)
}
