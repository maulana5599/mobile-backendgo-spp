package helpers

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(status int, message string, ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ResponseError(status int, message string, ctx *fiber.Ctx) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}
