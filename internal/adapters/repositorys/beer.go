package repositorys

import (
	"auth/hexagonal/internal/core/domain"
	"errors"
	"os"

	// "github.com/google/uuid"
)

func (b *DB) SaveBeer(beer *domain.Beer) error {
	// beer.ID = uuid.New().String()
	// if beer.UserId == "" {
	// 	return errors.New("UserId must be empty")
	// }

	// var user domain.User
	// if result := b.db.First(&user, "id = ?", beer.UserId); result.Error != nil {
	// 	return result.Error
	// }

	// beer.ShopName = user.ShopName

	// if result := b.db.Create(&beer); result.Error != nil {
	// 	return result.Error
	// }

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
	// ดึงข้อมูล Beer จากฐานข้อมูล
	beer, err := b.ReadBeer(id)
	if err != nil {
		return err
	}

	var path = "./uploads/" + beer.Image
	// ลบไฟล์รูปภาพ
	if err := deleteImage(path); err != nil {
		return err
	}

	// ลบข้อมูล Beer จากฐานข้อมูล
	result := b.db.Delete(&beer)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *DB) UpdateBeer(id string, beer *domain.Beer) error {

	value, err := b.ReadBeer(id)
	if err != nil {
		return err
	}

	// // ลบไฟล์รูปภาพ
	var path = "./uploads/" + value.Image
	if err := deleteImage(path); err != nil {
		return err
	}

	req := b.db.Model(&domain.Beer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"BeerName":    beer.BeerName,
		"Description": beer.Description,
		"Alcohol":     beer.Alcohol,
		"Stock":       beer.Stock,
		"Image":       beer.Image,
	})

	if req.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func deleteImage(imagePath string) error {
	// ตรวจสอบว่าไฟล์มีอยู่หรือไม่
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return errors.New("image not found")
	}

	// ลบไฟล์
	if err := os.Remove(imagePath); err != nil {
		return err
	}

	return nil
}
