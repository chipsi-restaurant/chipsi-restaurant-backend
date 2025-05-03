package repository

import (
	"chipsiBackend/domain"
	"context"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) domain.EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) Create(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	if err := r.db.WithContext(ctx).Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) GetAll(ctx context.Context) ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.WithContext(ctx).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) GetByUserID(ctx context.Context, userID uint) ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) UpdateFields(ctx context.Context, id int64, fields map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&domain.Event{}).Where("id = ?", id).Updates(fields).Error
}

func (r *eventRepository) GetByID(ctx context.Context, id int64) (*domain.Event, error) {
	var event domain.Event
	if err := r.db.WithContext(ctx).First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
