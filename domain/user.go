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
	Admin                []Admin              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"admin"`
	PasswordResetTokens  []PasswordResetToken `gorm:"constraint:OnDelete:CASCADE;" json:"passwordResetTokens"`
}

type UserDTO struct {
	ID        uint      `json:"id"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	Bonuses   int       `json:"bonuses"`
	IsAdmin   bool      `json:"isAdmin"`
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Phone:     user.Phone,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Bonuses:   user.Bonuses.Amount,
		IsAdmin:   user.Admin != nil && len(user.Admin) > 0,
	}
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	UpdateFields(ctx context.Context, id int64, fields map[string]interface{}) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Patch(ctx context.Context, id int64, fields map[string]interface{}) (*User, error)
}
