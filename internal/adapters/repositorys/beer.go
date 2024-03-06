package repositorys

import (
	"auth/hexagonal/internal/core/domain"
	"errors"

	"github.com/google/uuid"
)

func (b *DB) SaveBeer(beer *domain.Beer) error {
	beer.ID = uuid.New().String()
	if beer.UserId == "" {
		return errors.New("UserId must be empty")
	}

	var user domain.User
	if result := b.db.First(&user, "id = ?", beer.UserId); result.Error != nil {
		return result.Error
	}

	beer.ShopName = user.ShopName

	if result := b.db.Create(&beer); result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *DB) ReadBeers() ([]*domain.Beer, error) {
	var beers []*domain.Beer
	if result := b.db.Find(&beers); result.Error != nil {
		return nil, result.Error
	}

	return beers, nil
}

func (b *DB) ReadBeer(id string) (*domain.Beer, error) {
	beer := &domain.Beer{}
	req := b.db.Where("id = ?", id).First(&beer)
	if req.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return beer, nil
}

func (b *DB) DeleteBeer(id string) error {
	order := &domain.Beer{}
	req := b.db.Where("id = ?", id).First(&order)
	if req.RowsAffected == 0 {
		return errors.New("order not found")
	}

	result := b.db.Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *DB) UpdateBeer(id string, beer *domain.Beer) error {
	req := b.db.Model(&domain.Beer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"ID":          id,
		"OrderName":   beer.BeerName,
		"Description": beer.Description,
		"Price":       beer.Price,
		"Alcohol":     beer.Alcohol,
		"Stock":       beer.Stock,
	})
	if req.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
