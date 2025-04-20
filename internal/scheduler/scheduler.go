package scheduler

import (
	"chipsiBackend/domain"
	"gorm.io/gorm"
	"log/slog"
	"math/rand"
	"time"
)

func StartOrderStatusUpdater(db *gorm.DB, log *slog.Logger) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	log.Info("order status scheduler started")

	for {
		select {
		case <-ticker.C:
			if err := updateOrderStatuses(db); err != nil {
				log.Error("failed to update order statuses", "error", err)
			} else {
				log.Info("order statuses updated")
			}
		}
	}
}

func updateOrderStatuses(db *gorm.DB) error {
	type OrderWithDelivery struct {
		ID             uint
		Status         domain.OrderStatus
		DeliveryID     uint
		DeliveryStatus domain.DeliveryStatus
	}

	var orders []OrderWithDelivery

	err := db.Table("orders").
		Select("orders.id, orders.status, deliveries.id as delivery_id, deliveries.status as delivery_status").
		Joins("JOIN deliveries ON deliveries.order_id = orders.id").
		Where("orders.status != ?", domain.OrderDelivered).
		Find(&orders).Error

	if err != nil {
		return err
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, order := range orders {
		// Случайность: только 50% заказов будут обновлены
		if rnd.Intn(2) != 0 {
			continue
		}

		newOrderStatus, newDeliveryStatus := determineNextStatuses(order.Status, order.DeliveryStatus)

		if newOrderStatus != order.Status {
			if err := db.Model(&domain.Order{}).
				Where("id = ?", order.ID).
				Update("status", newOrderStatus).Error; err != nil {
				return err
			}
		}

		if newDeliveryStatus != order.DeliveryStatus {
			if err := db.Model(&domain.Delivery{}).
				Where("id = ?", order.DeliveryID).
				Update("status", newDeliveryStatus).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func determineNextStatuses(orderStatus domain.OrderStatus, deliveryStatus domain.DeliveryStatus) (domain.OrderStatus, domain.DeliveryStatus) {
	switch orderStatus {
	case domain.OrderPending:
		return domain.OrderPreparing, domain.DeliveryOnWay
	case domain.OrderPreparing:
		return domain.OrderDelivered, domain.DeliveryDone
	default:
		return orderStatus, deliveryStatus
	}
}
