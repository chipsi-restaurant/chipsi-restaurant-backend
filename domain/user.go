package domain

import "time"

type User struct {
	ID                   uint                 `gorm:"primaryKey"`
	Phone                string               `gorm:"unique;not null"`
	Email                string               `gorm:"unique;not null"`
	PasswordHash         string               `gorm:"not null"`
	FirstName            string               `gorm:"not null"`
	LastName             string               `gorm:"not null"`
	CreatedAt            time.Time            `gorm:"autoCreateTime"`
	Orders               []Order              `gorm:"constraint:OnDelete:CASCADE;"`
	Reservations         []Reservation        `gorm:"constraint:OnDelete:CASCADE;"`
	Events               []Event              `gorm:"constraint:OnDelete:CASCADE;"`
	SentCertificates     []GiftCertificate    `gorm:"foreignKey:SenderID"`
	ReceivedCertificates []GiftCertificate    `gorm:"foreignKey:ReceiverID"`
	Bonuses              Bonus                `gorm:"constraint:OnDelete:CASCADE;"`
	Admin                Admin                `gorm:"constraint:OnDelete:CASCADE;"`
	PasswordResetTokens  []PasswordResetToken `gorm:"constraint:OnDelete:CASCADE;"`
}
