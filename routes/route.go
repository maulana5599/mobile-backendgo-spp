package routes

import (
	handlerKeuangan "mobile-backendgo-spp/handler/keuangan"
	handlerMaster "mobile-backendgo-spp/handler/master"
	handlerTransaksi "mobile-backendgo-spp/handler/transaksi"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	/** Handler Keuangan */
	KeuanganRoutes(app)

	/** Handler Siswa */
	SiswaRoutes(app)

	/** Handler Tunggakan */
	TunggakanRoutes(app)

	/** Handler Transaksi Lain */
	TransaksiLain(app)

	/** Handler Transaksi Spp */
	TransaksiSpp(app)
}

func SiswaRoutes(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/siswa/:nis", handlerMaster.FindSiswaByNis)
}

func KeuanganRoutes(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/keuangan-pembayaran-siswa/:nis", handlerKeuangan.FindPeriodeKeuanganSiswa)
}

func TunggakanRoutes(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/tunggakan-spp", handlerKeuangan.FindTunggakanSppSiswa)
	v1.Get("/tunggakan-lain", handlerKeuangan.FindTunggakanLainSiswa)
}

func TransaksiLain(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/transaksi-lain-periode/:nis/:tahunajaran", handlerTransaksi.FindTransaksiLainPeriode)
}

func TransaksiSpp(app *fiber.App) {
	v1 := app.Group("v1")
	v1.Get("/transaksi-spp-periode/:nis/:tahunajaran", handlerTransaksi.FindTransaksiDetailPeriode)
}
