package usecase

import (
	"chipsiBackend/domain"
	"context"
	"time"
)

type bonusUsecase struct {
	bonusRepository domain.BonusRepository
	contextTimeout  time.Duration
}

func NewBonusUsecase(bonusRepository domain.BonusRepository, timeout time.Duration) domain.BonusUsecase {
	return &bonusUsecase{
		bonusRepository: bonusRepository,
		contextTimeout:  timeout,
	}
}

func (b bonusUsecase) Create(ctx context.Context, bonus *domain.Bonus) (*domain.Bonus, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.bonusRepository.Create(ctx, bonus)
}
