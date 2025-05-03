package domain

import (
	"context"
	"time"
)

type SignupRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(ctx context.Context, request SignupRequest) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry time.Duration) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry time.Duration) (refreshToken string, err error)
}
