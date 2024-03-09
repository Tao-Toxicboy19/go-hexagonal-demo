package repositorys

import (
	"auth/hexagonal/internal/core/domain"
	"errors"

	"github.com/google/uuid"
)

func (c *DB) SaveCart(cart *domain.Cart) error {
	cart.ID = uuid.New().String()

	if result := c.db.Create(&cart); result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *DB) ReadCarts(id string) ([]*domain.Cart, error) {
	var carts []*domain.Cart
	if err := c.db.Where("user_id = ?", id).Find(&carts).Error; err != nil {
		return nil, err
	}

	if len(carts) == 0 {
		return nil, errors.New("order not found")
	}

	return carts, nil
}

func (c *DB) DeleteCart(id string) error {
	cart := &domain.Cart{}
	req := c.db.Where("id = ?", id).First(&cart)
	if req.RowsAffected == 0 {
		return errors.New("order not found")
	}

	// ลบข้อมูล Beer จากฐานข้อมูล
	result := c.db.Delete(&cart)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
