package repositorys

import (
	"auth/hexagonal/internal/core/domain"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a *DB) SignUp(user *domain.User) (*domain.User, error) {
	req := a.db.Where("username = ? ", user.Username).First(&user)

	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}

	user.ID = uuid.New().String()
	user.Password = string(hash)

	req = a.db.Create(&user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}
	return user, nil
}

func (a *DB) SignIn(username, password string) (*LoginResponse, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// Implement the SignIn logic here
	jwtSecret := os.Getenv("JWT_SECRET")

	user, err := a.findUsername(username)
	if err != nil {
		return nil, err
	}

	err = a.verifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	accessToken, err := a.generateAccessToken(user.ID, jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.generateRefreshToken(user.ID, jwtSecret)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *DB) findUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	req := a.db.Where("username = ? ", username).First(&user)
	
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (a *DB) verifyPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password not matched")
	}
	return nil
}

func (a *DB) generateAccessToken(userId, jwtSecret string) (string, error) {
	payload := jwt.RegisteredClaims{
		Issuer:    "nayok-access",
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour).UTC()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(jwtSecret))
}

func (a *DB) generateRefreshToken(userId, jwtSecret string) (string, error) {
	payload := jwt.RegisteredClaims{
		Issuer:    "nayok-refresh",
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour).UTC()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(jwtSecret))
}
