package ports

import (
	"auth/hexagonal/internal/adapters/repositorys"
	"auth/hexagonal/internal/core/domain"
)

type AuthService interface {
	SignUp(user *domain.User) (*domain.User, error)
	SignIn(username, password string) (*repositorys.LoginResponse, error)
}

type AuthRepository interface {
	SignUp(user *domain.User) (*domain.User, error)
	SignIn(username, password string) (*repositorys.LoginResponse, error)
}

type BeerService interface {
	SaveBeer(order *domain.Beer) (*domain.Beer, error)
	ReadByUserId(id string) ([]*domain.Beer, error)
	ReadBeer(id string) (*domain.Beer, error)
	DeleteBeer(id string) error
	UpdateBeer(id string, order *domain.Beer) error
}

type BeerRepository interface {
	SaveBeer(order *domain.Beer) error
	ReadByUserId(id string) ([]*domain.Beer, error)
	ReadBeers() ([]*domain.Beer, error)
	ReadBeer(id string) (*domain.Beer, error)
	DeleteBeer(id string) error
	UpdateBeer(id string, order *domain.Beer) error
}

type CartsService interface {
	SaveCart(cart *domain.Cart) error
	ReadCarts(id string) ([]*domain.Cart, error)
	DeleteCart(id string) error
}

type CartsRepository interface {
	SaveCart(cart *domain.Cart) error
	ReadCarts(id string) ([]*domain.Cart, error)
	DeleteCart(id string) error
}
