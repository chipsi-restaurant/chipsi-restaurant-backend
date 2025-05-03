package domain

import (
	"context"
	"time"
)

type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"size:255;unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type PasswordResetTokenRepository interface {
	Create(ctx context.Context, token *PasswordResetToken) error
	GetAll(ctx context.Context) ([]PasswordResetToken, error) // заменяем GetByToken на GetAll
	Delete(ctx context.Context, id uint) error
}

type PasswordResetTokenUsecase interface {
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token string, newPassword string) error
}
