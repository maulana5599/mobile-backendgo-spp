package handler

import (
	"github.com/gofiber/fiber/v2"
)

func FindAllKeuanganSiswa(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"status": "maualan",
	})
}
