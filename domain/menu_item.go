package domain

import (
	"context"
	"mime/multipart"
)

type MenuItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:100;not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	CategoryID  uint    `json:"categoryId"`
	ImageURL    string  `json:"imageUrl"`
}

type MenuItemRepository interface {
	Create(ctx context.Context, menuItem *MenuItem) (*MenuItem, error)
	GetByID(ctx context.Context, id int64) (*MenuItem, error)
	GetAllByCategory(ctx context.Context, categoryId int64) ([]*MenuItem, error)
	Delete(ctx context.Context, id int64) error
}

type MenuItemUsecase interface {
	Create(ctx context.Context, menuItem *MenuItem, image multipart.File, handler *multipart.FileHeader) (*MenuItem, error)
	GetByID(ctx context.Context, id int64) (*MenuItem, error)
	GetAllByCategory(ctx context.Context, categoryId int64) ([]*MenuItem, error)
	Delete(ctx context.Context, id int64) error
}
