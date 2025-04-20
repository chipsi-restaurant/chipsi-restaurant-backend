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

func (r *bonusRepository) ChangeAmount(ctx context.Context, userID uint, used int, earned int) error {
	return r.db.WithContext(ctx).Model(&domain.Bonus{}).
		Where("user_id = ?", userID).
		UpdateColumn("amount", gorm.Expr("amount - ? + ?", used, earned)).Error
}
