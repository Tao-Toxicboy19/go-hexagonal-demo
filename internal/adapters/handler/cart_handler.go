package handler

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	service services.CartService
}

func NewCartHandler(CartService services.CartService) *CartHandler {
	return &CartHandler{service: CartService}
}

func (h *CartHandler) SaveCart(c *fiber.Ctx) error {
	var cart domain.Cart
	if err := c.BodyParser(&cart); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := h.service.SaveCart(&cart); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusCreated).JSON(cart)
}

func (h *CartHandler) ReadCarts(c *fiber.Ctx) error {
	id := c.Params("id")
	carts, err := h.service.ReadCarts(id)
	if err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(carts)
}

func (h *CartHandler) DeleteCart(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := Validate(c); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if err := h.service.DeleteCart(id); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}
