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

func recoverPanic() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if r := recover(); r != nil {
				// Set a "Connection: close" header on the response.
				ctx.Set("Connection", "close")
				log.Printf("%s", r)
			}
		}()
		return ctx.Next()
	}
}
