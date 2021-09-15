package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

var (
	port = flag.String("port", ":4000", "Port to listen")
)

func main() {
	// Parse flags
	flag.Parse()

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

	// Static files
	app.Static("/", "./ui/static")

	log.Printf("Starting server on %s", *port)
	log.Fatal(app.Listen(*port))
}
