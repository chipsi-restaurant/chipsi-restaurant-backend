package repository

import (
	"chipsiBackend/domain"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return categoryRepository{db: db}
}

func (r categoryRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	result := r.db.WithContext(ctx).Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil

}

func (r categoryRepository) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	category := domain.Category{}
	result := r.db.WithContext(ctx).Preload("MenuItems").Where("id = ?", id).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		} else {
			return nil, fmt.Errorf("db error: %v", result.Error)
		}
	}
	return &category, nil
}

func (r categoryRepository) GetAll(ctx context.Context, includeMenuItems bool) ([]*domain.Category, error) {
	var categories []*domain.Category
	var result *gorm.DB

	if includeMenuItems {
		result = r.db.WithContext(ctx).Preload("MenuItems").Find(&categories)
	} else {
		result = r.db.WithContext(ctx).Find(&categories)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get categories: %v", result.Error)
	}

	return categories, nil
}

func (r categoryRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&domain.Category{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete category: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("category with id %d not found", id)
	}
	return nil
}
