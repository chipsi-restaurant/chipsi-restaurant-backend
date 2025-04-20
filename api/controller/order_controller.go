package controller

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"net/http"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest domain.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
	}

	response, err := c.OrderUsecase.CreateOrderWithDelivery(r.Context(), &orderRequest)

	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *OrderController) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	response, err := c.OrderUsecase.GetUserOrders(r.Context())
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *OrderController) GetOrderStatuses(w http.ResponseWriter, r *http.Request) {
	orders, err := c.OrderUsecase.GetUserOrders(r.Context())
	if err != nil {
		httpErrors.JSONError(w, "failed to fetch orders", http.StatusInternalServerError)
		return
	}

	short := make([]map[string]interface{}, 0)
	for _, o := range orders {
		short = append(short, map[string]interface{}{
			"id":       o.ID,
			"status":   o.Status,
			"delivery": o.Delivery.Status,
		})
	}

	err = json.NewEncoder(w).Encode(short)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
