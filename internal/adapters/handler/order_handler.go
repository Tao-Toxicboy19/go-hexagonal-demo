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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.SaveOrder(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) ReadOrders(c *fiber.Ctx) error {
	orders, err := h.service.ReadOrders()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	return c.JSON(orders)
}
