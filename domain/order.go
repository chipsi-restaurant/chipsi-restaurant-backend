package domain

import (
	"context"
	"time"
)

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
type OrderItemsRequest struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type OrderAddress struct {
	Address         string `json:"address"`
	Floor           int    `json:"floor"`
	ApartmentNumber int    `json:"apartmentNumber"`
	IntercomCode    string `json:"intercomCode"`
	Notes           string `json:"notes"`
}

type OrderRequest struct {
	OrderItems     []OrderItemsRequest `json:"orderItems"`
	OrderAddresses OrderAddress        `json:"orderAddress"`
	Code           string              `json:"code"`
	UsedBonuses    int                 `json:"usedBonuses"`
}

type CreateOrderResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	TotalPrice  float64   `json:"totalPrice"`
	UsedBonuses int       `json:"usedBonuses"`
	FinalPrice  float64   `json:"finalPrice"`
	Status      string    `json:"status"`
}

type OrderResponse struct {
	ID          uint                `json:"id"`
	CreatedAt   time.Time           `json:"createdAt"`
	TotalPrice  float64             `json:"totalPrice"`
	UsedBonuses int                 `json:"usedBonuses"`
	FinalPrice  float64             `json:"finalPrice"`
	Status      string              `json:"status"`
	Items       []OrderItemResponse `json:"items"`
	Delivery    DeliveryResponse    `json:"delivery"`
}

func NewOrderResponse(order *Order) *OrderResponse {
	items := make([]OrderItemResponse, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, OrderItemResponse{
			ID:       item.MenuItem.ID,
			Name:     item.MenuItem.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
			Total:    float64(item.Quantity) * item.Price,
		})
	}

	delivery := DeliveryResponse{
		Address:         order.Delivery.Address,
		Floor:           order.Delivery.Floor,
		ApartmentNumber: order.Delivery.ApartmentNumber,
		IntercomCode:    order.Delivery.IntercomCode,
		Notes:           order.Delivery.Notes,
		Status:          string(order.Delivery.Status),
	}

	return &OrderResponse{
		ID:          order.ID,
		CreatedAt:   order.CreatedAt,
		TotalPrice:  order.TotalPrice,
		UsedBonuses: order.UsedBonuses,
		FinalPrice:  order.FinalPrice,
		Status:      string(order.Status),
		Items:       items,
		Delivery:    delivery,
	}
}

func NewCreateOrderResponse(order *Order) *CreateOrderResponse {
	return &CreateOrderResponse{
		ID:          order.ID,
		CreatedAt:   order.CreatedAt,
		TotalPrice:  order.TotalPrice,
		UsedBonuses: order.UsedBonuses,
		FinalPrice:  order.FinalPrice,
		Status:      string(order.Status),
	}
}

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error
	WithTransaction(ctx context.Context, fn func(OrderRepository) error) error
	GetByUserID(ctx context.Context, userID uint) ([]*Order, error)
}

type OrderUsecase interface {
	CreateOrderWithDelivery(ctx context.Context, request *OrderRequest) (*CreateOrderResponse, error)
	GetUserOrders(ctx context.Context) ([]*OrderResponse, error)
}
