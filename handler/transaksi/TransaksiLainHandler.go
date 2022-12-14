package handler

import (
	"mobile-backendgo-spp/helpers"
	"mobile-backendgo-spp/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindTransaksiLainPeriode(ctx *fiber.Ctx) error {
	nis := ctx.Params("nis")
	tahunajaran := ctx.Params("tahunajaran")
	nisSiswa, _ := strconv.Atoi(nis)
	tahunAjaran, _ := strconv.Atoi(tahunajaran)

	result, err := models.FindAllTransaksiPeriode(nisSiswa, tahunAjaran)

	if err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err.Error(), ctx)
	}

	return helpers.ResponseSuccess(http.StatusOK, "Get data successfully", ctx, result)
}
