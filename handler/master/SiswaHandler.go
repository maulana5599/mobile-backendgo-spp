package handler

import (
	"mobile-backendgo-spp/helpers"
	"mobile-backendgo-spp/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindSiswaByNis(ctx *fiber.Ctx) error {
	nis := ctx.Params("nis")
	nisSiswa, _ := strconv.Atoi(nis)

	result, err := models.FindSiswaByNis(nisSiswa)

	if err != nil {
		return helpers.ResponseError(http.StatusNotFound, err.Error(), ctx)
	}

	return helpers.ResponseSuccess(http.StatusOK, "Get data successfully", ctx, fiber.Map{
		"result": result,
		"name":   ctx.Get("Authorization"),
	})
}
