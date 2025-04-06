package usecase

import (
	"chipsiBackend/domain"
	"context"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (b userUsecase) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.userRepository.Create(ctx, user)
}

func (b userUsecase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.userRepository.GetByEmail(ctx, email)
}

func (b userUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.userRepository.GetByID(ctx, id)
}
