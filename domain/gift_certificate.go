package domain

type GiftCertificate struct {
	ID         uint `gorm:"primaryKey"`
	SenderID   *uint
	Sender     *User `gorm:"constraint:OnDelete:SET NULL;"`
	ReceiverID uint
	Receiver   User              `gorm:"constraint:OnDelete:SET NULL;"`
	Status     CertificateStatus `gorm:"default:'active'"`
}
