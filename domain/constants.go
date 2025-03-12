package domain

type OrderStatus string
type DeliveryStatus string
type ReservationStatus string
type EventStatus string
type CertificateStatus string
type AdminRole string

const (
	OrderPending   OrderStatus = "pending"
	OrderPreparing OrderStatus = "preparing"
	OrderDelivered OrderStatus = "delivered"
	OrderCanceled  OrderStatus = "canceled"

	DeliveryPending DeliveryStatus = "pending"
	DeliveryOnWay   DeliveryStatus = "on the way"
	DeliveryDone    DeliveryStatus = "delivered"

	ReservationPending   ReservationStatus = "pending"
	ReservationConfirmed ReservationStatus = "confirmed"
	ReservationCanceled  ReservationStatus = "canceled"

	EventPending   EventStatus = "pending"
	EventConfirmed EventStatus = "confirmed"
	EventRejected  EventStatus = "rejected"

	CertificateActive  CertificateStatus = "active"
	CertificateUsed    CertificateStatus = "used"
	CertificateExpired CertificateStatus = "expired"

	AdminRoleAdmin AdminRole = "admin"
)
