package usecase

import (
	"chipsiBackend/api/middleware"
	"chipsiBackend/domain"
	"context"
	"fmt"
	"strconv"
)

type orderUsecase struct {
	orderRepository           domain.OrderRepository
	menuItemRepository        domain.MenuItemRepository
	userRepository            domain.UserRepository
	giftCertificateRepository domain.GiftCertificateRepository
	bonusRepository           domain.BonusRepository
}

func NewOrderUsecase(orderRepository domain.OrderRepository,
	menuItemRepository domain.MenuItemRepository,
	userRepository domain.UserRepository,
	giftCertificateRepository domain.GiftCertificateRepository,
	bonusRepository domain.BonusRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository:           orderRepository,
		menuItemRepository:        menuItemRepository,
		userRepository:            userRepository,
		giftCertificateRepository: giftCertificateRepository,
		bonusRepository:           bonusRepository,
	}
}

func (u *orderUsecase) CreateOrderWithDelivery(ctx context.Context, request *domain.OrderRequest) (*domain.CreateOrderResponse, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("userId not found in context")
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	var orderItems []domain.OrderItem
	var totalPrice float64

	for _, itemReq := range request.OrderItems {
		menuItem, err := u.menuItemRepository.GetByID(ctx, int64(itemReq.ID))
		if err != nil {
			return nil, fmt.Errorf("menu item not found: %w", err)
		}

		price := menuItem.Price * float64(itemReq.Quantity)
		totalPrice += price

		orderItems = append(orderItems, domain.OrderItem{
			MenuItemID: menuItem.ID,
			Quantity:   int(itemReq.Quantity),
			Price:      menuItem.Price,
		})
	}

	var certificateAmount float64

	// Сертификат: если указан
	if request.Code != "" {
		cert, err := u.giftCertificateRepository.GetByCode(ctx, request.Code)
		if err != nil {
			return nil, fmt.Errorf("certificate error: %w", err)
		}
		if cert.Status == domain.CertificateUsed {
			return nil, fmt.Errorf("certificate already used")
		}
		certificateAmount = float64(cert.Amount)
	}

	delivery := &domain.Delivery{
		Address:         request.OrderAddresses.Address,
		Floor:           request.OrderAddresses.Floor,
		ApartmentNumber: request.OrderAddresses.ApartmentNumber,
		IntercomCode:    request.OrderAddresses.IntercomCode,
		Notes:           request.OrderAddresses.Notes,
	}

	finalPrice := totalPrice - float64(request.UsedBonuses) - certificateAmount
	if finalPrice < 0 {
		finalPrice = 0
	}

	order := &domain.Order{
		UserID:      uint(userID),
		TotalPrice:  totalPrice,
		UsedBonuses: request.UsedBonuses,
		FinalPrice:  finalPrice,
		OrderItems:  orderItems,
		Delivery:    *delivery,
	}

	err = u.orderRepository.WithTransaction(ctx, func(tx domain.OrderRepository) error {
		if err := tx.Create(ctx, order); err != nil {
			return err
		}

		// Списать бонусы и начислить 5%
		bonusesToAdd := int(finalPrice * 0.05)
		err := u.bonusRepository.ChangeAmount(ctx, uint(userID), request.UsedBonuses, bonusesToAdd)
		if err != nil {
			return err
		}

		// Если использовали сертификат — пометить как использованный
		if request.Code != "" {
			if err := u.giftCertificateRepository.UseCertificate(ctx, request.Code); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return domain.NewCreateOrderResponse(order), nil
}

func (u *orderUsecase) GetUserOrders(ctx context.Context) ([]*domain.OrderResponse, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("userId not found in context")
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid userId: %w", err)
	}

	orders, err := u.orderRepository.GetByUserID(ctx, uint(userID))
	if err != nil {
		return nil, err
	}

	responses := make([]*domain.OrderResponse, 0, len(orders))
	for _, order := range orders {
		responses = append(responses, domain.NewOrderResponse(order))
	}

	return responses, nil
}
