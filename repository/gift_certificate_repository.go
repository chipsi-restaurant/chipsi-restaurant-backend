package repository

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"context"
	"errors"
	"fmt"
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

func (r *giftCertificateRepository) GetByCode(ctx context.Context, code string) (*domain.GiftCertificate, error) {
	var giftCertificate domain.GiftCertificate
	result := r.db.WithContext(ctx).
		Preload("Receiver").
		Where("code = ?", code).
		First(&giftCertificate)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, httpErrors.NotFound
		} else {
			return nil, fmt.Errorf("db error: %v", result.Error)
		}
	}
	return &giftCertificate, nil
}

func (r *giftCertificateRepository) UseCertificate(ctx context.Context, code string) error {
	result := r.db.WithContext(ctx).
		Model(&domain.GiftCertificate{}).
		Where("code = ? AND status = ?", code, domain.CertificateActive).
		Updates(map[string]interface{}{
			"status": domain.CertificateUsed,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return httpErrors.NotFound
	}
	return nil
}
