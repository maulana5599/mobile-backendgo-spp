package main

import (
	handler "mobile-backendgo-spp/handler/keuangan"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("21RChLoDHkKQ03sWkQScxK6vtSR98pQr6hNzcegESoVhW3NRpyIoN12QyhoiHS72"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/maulana", handler.FindAllKeuanganSiswa)

	app.Listen(":3000")
}
