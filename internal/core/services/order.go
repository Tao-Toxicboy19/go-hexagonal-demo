package services

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/ports"
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

func (o *OrderService) ReadOrder(id string) (*domain.Order, error) {
	return o.repo.ReadOrder(id)
}

func (o *OrderService) DeleteOrder(id string) error {
	return o.repo.DeleteOrder(id)
}

func (o *OrderService) UpdateOrder(id string, order *domain.Order) error {
	return o.repo.UpdateOrder(id, order)
}
