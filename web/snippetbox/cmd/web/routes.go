package main

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router, s service) {
	r := route.Group("", recoverPanic())

	r.Get("/", home(s))
	r.Get("/snippet", showSnippet(s))
	r.Post("/snippet/create", createSnippet(s))

}
