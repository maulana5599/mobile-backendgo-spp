package handler

import (
	"mobile-backendgo-spp/helpers"
	"mobile-backendgo-spp/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FindTahunAjaranActive(ctx *fiber.Ctx) error {
	result, err := models.FindTahunAjaranActive()
	if err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err.Error(), ctx)
	}

	return helpers.ResponseSuccess(http.StatusOK, "Get data successfully", ctx, result)
}
