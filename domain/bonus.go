package domain

import (
	"context"
	"time"
)

type Bonus struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"unique" json:"user_id"`
	Amount      int       `gorm:"not null" json:"amount"`
	LastUpdated time.Time `gorm:"autoUpdateTime" json:"last_updated"`
}

type BonusRepository interface {
	Create(ctx context.Context, bonus *Bonus) (*Bonus, error)
}

type BonusUsecase interface {
	Create(ctx context.Context, bonus *Bonus) (*Bonus, error)
}
