package routes

import (
	handler "mobile-backendgo-spp/handler/keuangan"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/keuangan-pembayaran-siswa/:nis", handler.FindPeriodeKeuanganSiswa)
}
