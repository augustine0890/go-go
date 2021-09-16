package main

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router, s service) {
	route.Get("/", home())
	route.Get("/snippet", showSnippet)
	route.Post("/snippet/create", createSnippet(s))

}
