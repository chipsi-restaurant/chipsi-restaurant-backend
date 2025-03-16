package usecase

import (
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"chipsiBackend/internal/tokenutil"
	"chipsiBackend/pkg/httpErrors"
	"context"
	"strconv"
	"time"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	cfg            *bootstrap.Config
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, cfg *bootstrap.Config, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		cfg:            cfg,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) RefreshToken(ctx context.Context, request domain.RefreshTokenRequest) (*domain.RefreshTokenResponse, error) {
	id, err := rtu.ExtractIDFromToken(request.RefreshToken, rtu.cfg.App.JwtSecretKey)
	if err != nil {
		return nil, httpErrors.Unauthorized
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := rtu.GetUserByID(ctx, userId)
	if err != nil {
		return nil, httpErrors.NotFound
	}

	accessToken, err := rtu.CreateAccessToken(user, rtu.cfg.App.JwtSecretKey, rtu.cfg.App.AccessTokenExpires)

	if err != nil {
		return nil, httpErrors.InternalServerError
	}

	refreshToken, err := rtu.CreateRefreshToken(user, rtu.cfg.App.JwtSecretKey, rtu.cfg.App.RefreshTokenExpires)

	if err != nil {
		return nil, httpErrors.InternalServerError
	}

	response := &domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil

}

func (rtu *refreshTokenUsecase) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, id)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry time.Duration) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry time.Duration) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
