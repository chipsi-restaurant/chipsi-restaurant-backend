package usecase

import (
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"chipsiBackend/internal/tokenutil"
	"chipsiBackend/pkg/httpErrors"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	cfg            *bootstrap.Config
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration, cfg *bootstrap.Config) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		cfg:            cfg,
	}
}

func (lu *loginUsecase) Login(ctx context.Context, request domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := lu.GetUserByEmail(ctx, request.Login)
	if err != nil {
		return nil, httpErrors.NotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)) != nil {
		return nil, httpErrors.Unauthorized
	}

	accessToken, err := lu.CreateAccessToken(user, lu.cfg.App.JwtSecretKey, lu.cfg.App.AccessTokenExpires)

	if err != nil {
		return nil, httpErrors.InternalServerError
	}

	refreshToken, err := lu.CreateRefreshToken(user, lu.cfg.App.JwtSecretKey, lu.cfg.App.RefreshTokenExpires)

	if err != nil {
		return nil, httpErrors.InternalServerError
	}

	response := &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (lu *loginUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
