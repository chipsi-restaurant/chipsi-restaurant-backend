package repository

import (
	"chipsiBackend/domain"
	"context"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) domain.ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) Create(ctx context.Context, reservation *domain.Reservation) (*domain.Reservation, error) {
	result := r.db.WithContext(ctx).Create(reservation)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservation, nil
}

func (r *reservationRepository) GetAll(ctx context.Context) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	if err := r.db.WithContext(ctx).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) GetByUserID(ctx context.Context, userID uint) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) UpdateFields(ctx context.Context, id int64, fields map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&domain.Reservation{}).Where("id = ?", id).Updates(fields).Error
}

func (r *reservationRepository) GetByID(ctx context.Context, id int64) (*domain.Reservation, error) {
	var reservation domain.Reservation
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&reservation).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}
