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

type OrderService interface {
	SaveOrder(order *domain.Order) (*domain.Order, error)
	ReadOrders() ([]*domain.Order, error)
}

type OrderRepository interface {
	SaveOrder(order *domain.Order) error
	ReadOrders() ([]*domain.Order, error)
}
