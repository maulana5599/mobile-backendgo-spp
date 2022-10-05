package routes

import (
	handlerKeuangan "mobile-backendgo-spp/handler/keuangan"
	handlerMaster "mobile-backendgo-spp/handler/master"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	/** Handler Keuangan */
	KeuanganRoutes(app)

	/** Handler Siswa */
	SiswaRoutes(app)
}

func SiswaRoutes(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/siswa/:nis", handlerMaster.FindSiswaByNis)
}

func KeuanganRoutes(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/keuangan-pembayaran-siswa/:nis", handlerKeuangan.FindPeriodeKeuanganSiswa)
}
