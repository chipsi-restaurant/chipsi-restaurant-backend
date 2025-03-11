package domain

import "time"

type Event struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Date      time.Time `gorm:"type:date;not null"`
	StartTime time.Time `gorm:"type:time;not null"`
	Duration  int       `gorm:"not null"` // minutes
	Guests    int
	Type      string    `gorm:"size:50"`
	Price     float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
