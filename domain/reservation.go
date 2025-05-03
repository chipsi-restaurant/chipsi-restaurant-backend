package domain

import (
	"context"
	"time"
)

type Reservation struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	UserID    uint              `json:"userId"`
	Date      time.Time         `gorm:"type:date;not null" json:"date"`
	Time      time.Time         `gorm:"type:time;not null" json:"time"`
	Guests    int               `gorm:"not null" json:"guests"`
	Status    ReservationStatus `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time         `gorm:"autoCreateTime" json:"createdAt"`
}

type UpdateStatusRequest struct {
	Status  string `json:"status"`
	Comment string `json:"comment"`
}

type ReservationRequest struct {
	UserID uint
	Date   string `json:"date"`
	Time   string `json:"time"`
	Guests int    `json:"guests"`
}

type ReservationRepository interface {
	Create(ctx context.Context, reservation *Reservation) (*Reservation, error)
	GetAll(ctx context.Context) ([]Reservation, error)
	GetByUserID(ctx context.Context, userID uint) ([]Reservation, error)
	UpdateFields(ctx context.Context, id int64, fields map[string]interface{}) error
	GetByID(ctx context.Context, id int64) (*Reservation, error)
}

type ReservationUsecase interface {
	Create(ctx context.Context, request *ReservationRequest) (*Reservation, error)
	GetAll(ctx context.Context) ([]Reservation, error)
	GetByUserID(ctx context.Context, userID uint) ([]Reservation, error)
	UpdateStatus(ctx context.Context, id int64, fields map[string]interface{}) error
}
