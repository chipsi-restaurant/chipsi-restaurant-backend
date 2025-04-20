package domain

type OrderItem struct {
	ID         uint `gorm:"primaryKey"`
	OrderID    uint
	MenuItemID uint
	MenuItem   MenuItem
	Quantity   int     `gorm:"not null"`
	Price      float64 `gorm:"not null"`
}

type OrderItemResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
