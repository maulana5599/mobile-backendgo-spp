package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJwt(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		KeyFunc: customKeyFunc(),
	}))
}

func customKeyFunc() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// Always check the signing method
		if t.Method.Alg() != jwtware.HS256 {
			return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
		}

		// TODO custom implementation of loading signing key like from a database
		signingKey := "21RChLoDHkKQ03sWkQScxK6vtSR98pQr6hNzcegESoVhW3NRpyIoN12QyhoiHS72"

		return []byte(signingKey), nil
	}
}
