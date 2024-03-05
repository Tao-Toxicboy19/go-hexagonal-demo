package handler

import "github.com/gofiber/fiber/v2"

func HandlerError(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": err.Error()})
}