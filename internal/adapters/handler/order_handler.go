package handler

import (
	"auth/hexagonal/internal/core/domain"
	"auth/hexagonal/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(OrderService services.OrderService) *OrderHandler {
	return &OrderHandler{service: OrderService}
}

func (h *OrderHandler) SaveOrder(c *fiber.Ctx) error {
	var order domain.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	valid, err := ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.service.SaveOrder(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) ReadOrders(c *fiber.Ctx) error {
	orders, err := h.service.ReadOrders()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	valid, err := ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (h *OrderHandler) ReadOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	order, err := h.service.ReadOrder(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	valid, err := ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	valid, err := ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.service.DeleteOrder(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"})
}

func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	var order domain.Order
	id := c.Params("id")

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	valid, err := ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.service.UpdateOrder(id, &order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
