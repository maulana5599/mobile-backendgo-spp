package handler

import (
	"mobile-backendgo-spp/helpers"
	"mobile-backendgo-spp/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindTunggakanSppSiswa(ctx *fiber.Ctx) error {

	jurusan, _ := strconv.Atoi(ctx.Query("jurusan"))
	kelas, _ := strconv.Atoi(ctx.Query("kelas"))
	tahunajaran, _ := strconv.Atoi(ctx.Query("tahunajaran"))

	result, err := models.FindTunggakanSiswa(jurusan, kelas, tahunajaran)

	if err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err.Error(), ctx)
	}

	var generateBiaya = []map[string]interface{}{
		{"bulan": "januari", "biaya": result.Januari},
		{"bulan": "februari", "biaya": result.Februari},
		{"bulan": "maret", "biaya": result.Maret},
		{"bulan": "april", "biaya": result.April},
		{"bulan": "mei", "biaya": result.Mei},
		{"bulan": "juni", "biaya": result.Juni},
		{"bulan": "juli", "biaya": result.Juli},
		{"bulan": "agustus", "biaya": result.Agustus},
		{"bulan": "september", "biaya": result.September},
		{"bulan": "oktober", "biaya": result.Oktober},
		{"bulan": "november", "biaya": result.November},
		{"bulan": "desember", "biaya": result.Desember},
	}

	resultData := fiber.Map{
		"tunggakan": generateBiaya,
		"detail":    result,
	}

	return helpers.ResponseSuccess(http.StatusOK, "Get data successfully", ctx, resultData)
}

func FindTunggakanLainSiswa(ctx *fiber.Ctx) error {
	jurusan, _ := strconv.Atoi(ctx.Query("jurusan"))
	kelas, _ := strconv.Atoi(ctx.Query("kelas"))
	tahunajaran, _ := strconv.Atoi(ctx.Query("tahunajaran"))

	result, err := models.FindTunggakanLainSiswa(jurusan, kelas, tahunajaran)

	if err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err.Error(), ctx)
	}

	return helpers.ResponseSuccess(http.StatusOK, "Get data successfully", ctx, result)
}
