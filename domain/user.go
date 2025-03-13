package domain

import (
	"context"
	"time"
)

type User struct {
	ID                   uint                 `gorm:"primaryKey" json:"id"`
	Phone                string               `gorm:"unique;not null" json:"phone"`
	Email                string               `gorm:"unique;not null" json:"email"`
	PasswordHash         string               `gorm:"not null" json:"password_hash"`
	FirstName            string               `gorm:"not null" json:"first_name"`
	LastName             string               `gorm:"not null" json:"last_name"`
	CreatedAt            time.Time            `gorm:"autoCreateTime" json:"created_at"`
	Orders               []Order              `gorm:"constraint:OnDelete:CASCADE;" json:"orders"`
	Reservations         []Reservation        `gorm:"constraint:OnDelete:CASCADE;" json:"reservations"`
	Events               []Event              `gorm:"constraint:OnDelete:CASCADE;" json:"events"`
	SentCertificates     []GiftCertificate    `gorm:"foreignKey:SenderID" json:"sent_certificates"`
	ReceivedCertificates []GiftCertificate    `gorm:"foreignKey:ReceiverID" json:"received_certificates"`
	Bonuses              Bonus                `gorm:"constraint:OnDelete:CASCADE;" json:"bonuses"`
	Admin                Admin                `gorm:"constraint:OnDelete:CASCADE;" json:"admin"`
	PasswordResetTokens  []PasswordResetToken `gorm:"constraint:OnDelete:CASCADE;" json:"password_reset_tokens"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
