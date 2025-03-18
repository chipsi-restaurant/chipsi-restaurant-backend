package usecase

import (
	"chipsiBackend/domain"
	"chipsiBackend/internal/tokenutil"
	"chipsiBackend/pkg/httpErrors"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, request domain.SignupRequest) (*domain.User, error) {

	_, err := su.GetUserByEmail(c, request.Email)
	if err == nil {
		return nil, httpErrors.ExistsEmailError
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        request.Email,
		Phone:        request.Phone,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		PasswordHash: string(encryptedPassword),
	}

	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	user, err = su.userRepository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (su *signupUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
