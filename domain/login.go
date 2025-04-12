package domain

import (
	"context"
	"time"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	Login(ctx context.Context, request LoginRequest) (*LoginResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry time.Duration) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry time.Duration) (refreshToken string, err error)
}
