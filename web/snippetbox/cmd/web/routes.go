package main

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router, s service) {
	route.Get("/", home(s))
	route.Get("/snippet", showSnippet(s))
	route.Post("/snippet/create", createSnippet(s))

}
