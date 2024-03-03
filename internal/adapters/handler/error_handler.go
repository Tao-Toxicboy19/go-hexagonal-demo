package handler

import "github.com/gofiber/fiber/v2"

func HandlerError(statusCode int, err error) *fiber.Error {
	return fiber.NewError(statusCode, err.Error())
}
