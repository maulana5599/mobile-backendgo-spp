package middleware

import (
	jwtware "github.com/gofiber/jwt/v3"
)

func ValidateJwt() bool {
	err := jwtware.New(jwtware.Config{
		SigningKey: []byte("21RChLoDHkKQ03sWkQScxK6vtSR98pQr6hNzcegESoVhW3NRpyIoN12QyhoiHS72"),
	})

	return err != nil
}
