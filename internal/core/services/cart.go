package services

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/ports"
)

type CartService struct {
	repo ports.CartsRepository
}

func NewCartService(repo ports.CartsRepository) *CartService {
	return &CartService{repo: repo}
}

func (a *CartService) SaveCart(cart *domain.Cart)  error {
	return a.repo.SaveCart(cart)
}

func (o *CartService) ReadCarts(id string) ([]*domain.Cart, error) {
	return o.repo.ReadCarts(id)
}

func (o *CartService) DeleteCart(id string) error {
	return o.repo.DeleteCart(id)
}