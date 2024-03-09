package domain

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	ShopName string
	Beers    []Beer
	Carts    []Cart
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
	Carts       []Cart
}

type Cart struct {
	ID     string `gorm:"primaryKey"`
	BeerId string
	UserId string
}
