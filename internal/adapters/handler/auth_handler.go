package handler

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"

	"github.com/gofiber/fiber/v2"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if _, err := h.service.SignUp(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	response, err := h.service.SignIn(user.Username, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
	})
}
