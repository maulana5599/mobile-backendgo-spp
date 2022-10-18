package handler

import (
	"log"
	"mobile-backendgo-spp/helpers"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	// Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	privateKey = "21RChLoDHkKQ03sWkQScxK6vtSR98pQr6hNzcegESoVhW3NRpyIoN12QyhoiHS72"
)

func Auth(ctx *fiber.Ctx) error {

	username := ctx.FormValue("user")
	password := ctx.FormValue("password")

	if username != "maulana" || password != "maulana" {
		return helpers.ResponseError(http.StatusUnauthorized, "Authorized Failed!", ctx)
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  username,
		"admin": true,
		"role":  []string{"admin", "user"},
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(privateKey))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	result := fiber.Map{
		"token": t,
	}
	return helpers.ResponseSuccess(http.StatusOK, "Authentication Successfully", ctx, result)

}
