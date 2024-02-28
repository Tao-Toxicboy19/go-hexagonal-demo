package services

import (
	"auth/hexagonal/internal/adapters/repositorys"
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/ports"
)

type AuthService struct {
	repo ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) SignUp(user *domain.User) (*domain.User, error) {
	return a.repo.SignUp(user)
}

func (a *AuthService) SignIn(username, password string) (*repositorys.LoginResponse, error) {
	return a.repo.SignIn(username, password)
}
