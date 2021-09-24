package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func logRequest() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Printf("Method: %s - IP: %s", ctx.Request().Header.Method(), ctx.IP())
		return ctx.Next()
	}
}
