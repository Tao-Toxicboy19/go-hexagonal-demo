package services

import (
	"auth/hexagonal/core/domain"
	"auth/hexagonal/core/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (o *OrderService) SaveOrder(order *domain.Order) error {
	return o.repo.SaveOrder(order)
}

func (o *OrderService) ReadOrders() ([]*domain.Order, error) {
	return o.repo.ReadOrders()
}
