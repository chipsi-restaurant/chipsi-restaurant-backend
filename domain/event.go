package domain

import (
	"context"
	"time"
)

type Event struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"userId"`
	Date      time.Time   `gorm:"type:date;not null" json:"date"`
	StartTime time.Time   `gorm:"type:time;not null" json:"startTime"`
	Duration  int         `gorm:"not null" json:"duration"`
	Guests    int         `json:"guests"`
	Type      string      `gorm:"size:50" json:"type"`
	Price     float64     `gorm:"not null" json:"price"`
	Status    EventStatus `gorm:"default:'pending'" json:"status"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"createdAt"`
}

type EventRequest struct {
	UserID    uint
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	Duration  int    `json:"duration"`
	Guests    int    `json:"guests"`
	Type      string `json:"type"`
}

type EventRepository interface {
	Create(ctx context.Context, event *Event) (*Event, error)
	GetAll(ctx context.Context) ([]Event, error)
	GetByUserID(ctx context.Context, userID uint) ([]Event, error)
	UpdateFields(ctx context.Context, id int64, fields map[string]interface{}) error
	GetByID(ctx context.Context, id int64) (*Event, error)
}

type EventUsecase interface {
	Create(ctx context.Context, request *EventRequest) (*Event, error)
	GetAll(ctx context.Context) ([]Event, error)
	GetByUserID(ctx context.Context, userID uint) ([]Event, error)
	UpdateStatus(ctx context.Context, id int64, fields map[string]interface{}) error
}
