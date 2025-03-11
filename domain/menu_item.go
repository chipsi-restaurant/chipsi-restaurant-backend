package domain

type MenuItem struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Description string
	Price       float64 `gorm:"not null"`
	CategoryID  uint
	ImageURL    string
}
