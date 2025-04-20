package domain

import (
	"context"
	"time"
)

type GiftCertificate struct {
	ID         uint `gorm:"primaryKey"`
	Amount     uint
	Code       string
	SenderID   uint
	Sender     *User `gorm:"foreignKey:SenderID;constraint:OnDelete:SET NULL;"`
	ReceiverID uint
	Receiver   *User             `gorm:"foreignKey:ReceiverID;constraint:OnDelete:SET NULL;"`
	Status     CertificateStatus `gorm:"default:'active'"`
	CreatedAt  time.Time         `gorm:"autoCreateTime"`
}

type GiftCertificateRequest struct {
	Amount        uint   `json:"amount"`
	SenderID      uint   `json:"senderId"`
	ReceiverEmail string `json:"receiverEmail"`
}

type GiftCertificateResponse struct {
	Amount        uint      `json:"amount"`
	ReceiverEmail string    `json:"receiverEmail"`
	CreatedAt     time.Time `json:"createdAt"`
}

type GiftCertificateRepository interface {
	Create(ctx context.Context, giftCertificate *GiftCertificate) (*GiftCertificate, error)
	GetBySenderID(ctx context.Context, id int64) ([]*GiftCertificate, error)
	GetByReceiverID(ctx context.Context, id int64) ([]*GiftCertificate, error)
	GetByCode(ctx context.Context, code string) (*GiftCertificate, error)
	UseCertificate(ctx context.Context, code string) error
}

type GiftCertificateUsecase interface {
	Create(ctx context.Context, request *GiftCertificateRequest) (*GiftCertificateResponse, error)
	GetBySenderID(ctx context.Context, id int64) ([]*GiftCertificateResponse, error)
	GetByCode(ctx context.Context, code string, userID int64) (*GiftCertificateResponse, error)
}
