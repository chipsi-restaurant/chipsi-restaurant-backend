package domain

type Admin struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Role   AdminRole `gorm:"default:'admin'"`
}
