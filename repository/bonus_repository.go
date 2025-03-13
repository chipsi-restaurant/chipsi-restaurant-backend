package repository

import (
	"chipsiBackend/domain"
	"context"
	"gorm.io/gorm"
)

type bonusRepository struct {
	db *gorm.DB
}

func NewBonusRepository(db *gorm.DB) domain.BonusRepository {
	return &bonusRepository{db: db}
}

func (r *bonusRepository) Create(ctx context.Context, bonus *domain.Bonus) (*domain.Bonus, error) {
	result := r.db.WithContext(ctx).Create(bonus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bonus, nil
}
