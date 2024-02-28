package repositorys

import (
	"auth/hexagonal/internal/core/domain"
	"errors"

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

func (o *DB) ReadOrder(id string) (*domain.Order, error) {
	order := &domain.Order{}
	req := o.db.Where("id = ?", id).First(&order)
	if req.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (o *DB) DeleteOrder(id string) error {
	order := &domain.Order{}
	req := o.db.Where("id = ?", id).First(&order)
	if req.RowsAffected == 0 {
		return errors.New("order not found")
	}

	result := o.db.Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *DB) UpdateOrder(id string, order *domain.Order) error {
	req := o.db.Model(&domain.Order{}).Where("id = ?", id).Updates(map[string]interface{}{
		"ID":        id,
		"OrderName": order.OrderName,
		"Price":     order.Price,
	})
	if req.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
