package repository

import (
	"chipsiBackend/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type passwordResetTokenRepository struct {
	db *gorm.DB
}

func NewPasswordResetTokenRepository(db *gorm.DB) domain.PasswordResetTokenRepository {
	return &passwordResetTokenRepository{db: db}
}

func (r *passwordResetTokenRepository) Create(ctx context.Context, token *domain.PasswordResetToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *passwordResetTokenRepository) GetAll(ctx context.Context) ([]domain.PasswordResetToken, error) {
	var tokens []domain.PasswordResetToken
	if err := r.db.WithContext(ctx).
		Where("expires_at > ?", time.Now()).
		Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

func (r *passwordResetTokenRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.PasswordResetToken{}, id).Error
}
