package usecase

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"chipsiBackend/pkg/utils"
	"context"
	"time"
)

var allowedFields = map[string]bool{
	"firstName": true,
	"lastName":  true,
	"email":     true,
	"phone":     true,
}

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

func (u *userUsecase) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.userRepository.Create(ctx, user)
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByEmail(ctx, email)
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByID(ctx, id)
}

func (u *userUsecase) Patch(ctx context.Context, id int64, fields map[string]interface{}) (*domain.User, error) {
	updates := make(map[string]interface{})
	for key, val := range fields {
		if allowedFields[key] {
			updates[utils.ToSnakeCase(key)] = val
		}
	}
	if len(updates) == 0 {
		return nil, httpErrors.BadRequest
	}

	err := u.userRepository.UpdateFields(ctx, id, updates)

	if err != nil {
		return nil, err
	}

	return u.userRepository.GetByID(ctx, id)
}
