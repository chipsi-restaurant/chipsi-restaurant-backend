package domain

import (
	"context"
	"time"
)

type User struct {
	ID                   uint                 `gorm:"primaryKey" json:"id"`
	Phone                string               `gorm:"unique;not null" json:"phone"`
	Email                string               `gorm:"unique;not null" json:"email"`
	PasswordHash         string               `gorm:"not null" json:"passwordHash"`
	FirstName            string               `gorm:"not null" json:"firstName"`
	LastName             string               `gorm:"not null" json:"lastName"`
	CreatedAt            time.Time            `gorm:"autoCreateTime" json:"createdAt"`
	Orders               []Order              `gorm:"constraint:OnDelete:CASCADE;" json:"orders"`
	Reservations         []Reservation        `gorm:"constraint:OnDelete:CASCADE;" json:"reservations"`
	Events               []Event              `gorm:"constraint:OnDelete:CASCADE;" json:"events"`
	SentCertificates     []GiftCertificate    `gorm:"foreignKey:SenderID" json:"sentCertificates"`
	ReceivedCertificates []GiftCertificate    `gorm:"foreignKey:ReceiverID" json:"receivedCertificates"`
	Bonuses              Bonus                `gorm:"constraint:OnDelete:CASCADE;" json:"bonuses"`
	Admin                Admin                `gorm:"constraint:OnDelete:CASCADE;" json:"admin"`
	PasswordResetTokens  []PasswordResetToken `gorm:"constraint:OnDelete:CASCADE;" json:"passwordResetTokens"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
