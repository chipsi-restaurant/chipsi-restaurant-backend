package repository

import (
	"chipsiBackend/domain"
	"context"
	"gorm.io/gorm"
)

type giftCertificateRepository struct {
	db *gorm.DB
}

func NewGiftCertificateRepository(db *gorm.DB) domain.GiftCertificateRepository {
	return &giftCertificateRepository{db: db}
}

func (r *giftCertificateRepository) Create(ctx context.Context, giftCertificate *domain.GiftCertificate) (*domain.GiftCertificate, error) {
	result := r.db.WithContext(ctx).Create(giftCertificate)
	if result.Error != nil {
		return nil, result.Error
	}
	return giftCertificate, nil
}

func (r *giftCertificateRepository) GetBySenderID(ctx context.Context, id int64) ([]*domain.GiftCertificate, error) {
	var certificates []*domain.GiftCertificate
	err := r.db.WithContext(ctx).
		Preload("Receiver").
		Where("sender_id = ?", id).
		Find(&certificates).Error
	return certificates, err
}

func (r *giftCertificateRepository) GetByReceiverID(ctx context.Context, id int64) ([]*domain.GiftCertificate, error) {
	var certificates []*domain.GiftCertificate
	err := r.db.WithContext(ctx).
		Preload("Receiver").
		Where("receiver_id = ?", id).
		Find(&certificates).Error
	return certificates, err
}
