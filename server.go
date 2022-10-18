package main

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	config.DatabaseConnection()
	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("21RChLoDHkKQ03sWkQScxK6vtSR98pQr6hNzcegESoVhW3NRpyIoN12QyhoiHS72"),
	// }))

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", "admin")
		c.Locals("Authorization")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.RouteInit(app)

	app.Listen(":1323")
}
