package domain

type Delivery struct {
	ID              uint   `gorm:"primaryKey"`
	OrderID         uint   `gorm:"uniqueIndex"`
	Address         string `gorm:"not null"`
	Floor           int    `gorm:"not null"`
	ApartmentNumber int    `gorm:"not null"`
	IntercomCode    string
	Notes           string
	Status          DeliveryStatus `gorm:"default:'pending'"`
}
