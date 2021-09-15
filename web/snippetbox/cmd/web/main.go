package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize html tamplate engine
	engine := html.New("./ui/html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Log middleware config
	app.Use(logger.New())

	app.Get("/", home)
	app.Get("/snippet", showSnippet)
	app.Post("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	app.Listen(":4000")
}
