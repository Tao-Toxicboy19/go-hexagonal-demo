package repositorys

import (
	"auth/hexagonal/core/domain"

	"github.com/google/uuid"
)

func (o *DB) SaveOrder(order *domain.Order) error {
	order.ID = uuid.New().String()
	if result := o.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *DB) ReadOrders() ([]*domain.Order, error) {
	var orders []*domain.Order
	if result := o.db.Find(&orders); result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
