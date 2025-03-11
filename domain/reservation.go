package domain

import "time"

type Reservation struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Date      time.Time         `gorm:"type:date;not null"`
	Time      time.Time         `gorm:"type:time;not null"`
	Guests    int               `gorm:"not null"`
	Status    ReservationStatus `gorm:"default:'pending'"`
	CreatedAt time.Time         `gorm:"autoCreateTime"`
}
