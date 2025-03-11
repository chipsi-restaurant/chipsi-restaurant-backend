package domain

import "time"

type Bonus struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"unique"`
	Amount      int       `gorm:"not null"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
}
