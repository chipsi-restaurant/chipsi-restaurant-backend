package domain

import "time"

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	User        User
	TotalPrice  float64     `gorm:"not null"`
	UsedBonuses int         `gorm:"default:0"`
	FinalPrice  float64     `gorm:"not null"`
	Status      OrderStatus `gorm:"default:'pending'"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	OrderItems  []OrderItem `gorm:"constraint:OnDelete:CASCADE;"`
	Delivery    Delivery    `gorm:"constraint:OnDelete:CASCADE;"`
}
