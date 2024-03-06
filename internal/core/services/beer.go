package services

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/ports"
)

type BeerService struct {
	repo ports.BeerRepository
}

func NewBeerService(repo ports.BeerRepository) *BeerService {
	return &BeerService{repo: repo}
}

func (o *BeerService) SaveBeer(beer *domain.Beer) error {
	return o.repo.SaveBeer(beer)
}

func (o *BeerService) ReadBeers() ([]*domain.Beer, error) {
	return o.repo.ReadBeers()
}

func (o *BeerService) ReadBeer(id string) (*domain.Beer, error) {
	return o.repo.ReadBeer(id)
}

func (o *BeerService) DeleteBeer(id string) error {
	return o.repo.DeleteBeer(id)
}

func (o *BeerService) UpdateBeer(id string, order *domain.Beer) error {
	return o.repo.UpdateBeer(id, order)
}
