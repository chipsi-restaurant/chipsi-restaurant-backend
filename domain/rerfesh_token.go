package domain

import (
	"context"
	"time"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type RefreshTokenUsecase interface {
	RefreshToken(ctx context.Context, request RefreshTokenRequest) (*RefreshTokenResponse, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	CreateAccessToken(user *User, secret string, expiry time.Duration) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry time.Duration) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
