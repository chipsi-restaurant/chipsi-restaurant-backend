package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	ID         uint   `json:"id"`
	jwt.Claims `json:"claims"`
}

type JwtCustomRefreshClaims struct {
	ID         uint `json:"id"`
	jwt.Claims `json:"claims"`
}
