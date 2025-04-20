package repository

import (
	"chipsiBackend/domain"
	"context"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) WithTransaction(ctx context.Context, fn func(domain.OrderRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &orderRepository{db: tx}
		return fn(txRepo)
	})
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) GetByUserID(ctx context.Context, userID uint) ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.db.WithContext(ctx).
		Preload("OrderItems.MenuItem").
		Preload("Delivery").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}
