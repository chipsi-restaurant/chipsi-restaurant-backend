package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email string `json:"Email"`
	Phone string `json:"Phone"`
	ID    uint   `json:"id"`
	jwt.Claims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.Claims
}
