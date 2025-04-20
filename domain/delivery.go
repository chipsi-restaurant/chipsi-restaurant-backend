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

type DeliveryResponse struct {
	Address         string `json:"address"`
	Floor           int    `json:"floor"`
	ApartmentNumber int    `json:"apartmentNumber"`
	IntercomCode    string `json:"intercomCode"`
	Notes           string `json:"notes"`
	Status          string `json:"status"`
}
