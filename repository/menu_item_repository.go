package repository

import (
	"chipsiBackend/domain"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type menuItemRepository struct {
	db *gorm.DB
}

func NewMenuItemRepository(db *gorm.DB) domain.MenuItemRepository {
	return &menuItemRepository{db: db}
}

func (r menuItemRepository) Create(ctx context.Context, menuItem *domain.MenuItem) (*domain.MenuItem, error) {
	result := r.db.WithContext(ctx).Create(menuItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return menuItem, nil
}

func (r menuItemRepository) GetByID(ctx context.Context, id int64) (*domain.MenuItem, error) {
	menuItem := domain.MenuItem{}
	var result *gorm.DB

	result = r.db.WithContext(ctx).Where("id = ?", id).First(&menuItem)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get categories: %v", result.Error)
	}
	return &menuItem, nil
}

func (r menuItemRepository) GetAllByCategory(ctx context.Context, categoryId int64) ([]*domain.MenuItem, error) {
	var menuItems []*domain.MenuItem
	var result *gorm.DB

	result = r.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&menuItems)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get categories: %v", result.Error)
	}
	return menuItems, nil
}

func (r menuItemRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&domain.MenuItem{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete menuItem: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("menuItem with id %d not found", id)
	}
	return nil
}
