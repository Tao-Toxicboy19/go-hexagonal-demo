package handler

import (
	"auth/hexagonal/core/domain"
	"auth/hexagonal/core/services"
	"fmt"

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
	fmt.Println(user)
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
	return c.Status(fiber.StatusOK).JSON(user)
}
