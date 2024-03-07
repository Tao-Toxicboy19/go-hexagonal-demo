package handler

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(AuthService services.AuthService) *AuthHandler {
	return &AuthHandler{service: AuthService}
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if _, err := h.service.SignUp(&user); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	response, err := h.service.SignIn(user.Username, user.Password)
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
	})
}

func (h *AuthHandler) DecodeToken(c *fiber.Ctx) error {

	return c.SendString("Token decoded successfully")
}

func ValidateToken(header string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if header == "" {
		return false, errors.New("token not found")
	}

	// Check if header starts with "Bearer "
	if !strings.HasPrefix(header, "Bearer ") {
		return false, errors.New("invalid token format")
	}

	// Extract Bearer token
	tokenHeader := header[7:]

	token, err := jwt.ParseWithClaims(tokenHeader, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return false, errors.New("token invalid")
	}

	if !token.Valid {
		return false, errors.New("token not valid")
	}

	payload, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || payload.ExpiresAt == nil || payload.ExpiresAt.Before(time.Now().UTC()) {
		return false, errors.New("token has expired")
	}

	// Check if token is a refresh token
	// if payload.Issuer == "nayok-refresh" {
	// 	return false, errors.New("token is a refresh token, please use access token")
	// }

	return true, nil
}

func Validate(c *fiber.Ctx) error {
	valid, err := ValidateToken(c.Get("Authorization"))
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
