package domain

type OrderItem struct {
	ID         uint `gorm:"primaryKey"`
	OrderID    uint
	MenuItemID uint
	MenuItem   MenuItem
	Quantity   int     `gorm:"not null"`
	Price      float64 `gorm:"not null"`
}
