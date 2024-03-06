package domain

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	ShopName string
	Beers    []Beer
}

type Beer struct {
	ID          string `gorm:"primaryKey"`
	BeerName    string
	Description string
	Alcohol     string
	Price       float64
	Stock       float64
	Image       string
	ShopName    string
	UserId      string
}
