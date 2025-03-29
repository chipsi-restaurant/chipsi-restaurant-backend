package domain

import "context"

type Category struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Name      string      `gorm:"size:100;not null" json:"name"`
	MenuItems *[]MenuItem `json:"menuItems,omitempty"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	GetAll(ctx context.Context, includeMenuItems bool) ([]*Category, error)
	Delete(ctx context.Context, id int64) error
}

type CategoryUsecase interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	GetAll(ctx context.Context, includeMenuItems bool) ([]*Category, error)
	Delete(ctx context.Context, id int64) error
}
